package elemental

import (
	"context"
	"sync"
)

// AtomicJob takes a func() error and returns a func(context.Context) error.
// The returned function can be called as many time as you like, but only
// one instance of the given job can be run at the same time.
//
// The returned function will either execute job if it
// it not already running or wait for the currently running job to finish.
// In both cases, the returned error from the job will be forwareded and returned
// to every caller.
//
// You must pass a context.Context to the returned function so you can
// control how much time you are willing to wait for the job to complete.
//
// If you wish to change some external state from within the job function,
// it is your responsibility to ensure everything is thread safe.
func AtomicJob(job func() error) func(context.Context) error {

	var l sync.RWMutex
	var errorChs []chan error

	sem := make(chan struct{}, 1)

	return func(ctx context.Context) error {

		errCh := make(chan error)

		l.Lock()
		errorChs = append(errorChs, errCh)
		l.Unlock()

		select {
		case sem <- struct{}{}:

			go func() {

				err := job()

				l.Lock()
				for _, ch := range errorChs {
					select {
					case ch <- err:
					default:
					}
				}
				errorChs = nil
				l.Unlock()

				<-sem
			}()

		default:
		}

		select {
		case err := <-errCh:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
