package elemental

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAtomicJob(t *testing.T) {

	Convey("Given I call a wrapped job once and it works", t, func() {

		var counter int64

		f := AtomicJob(func() error {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			return nil
		})

		err := f(context.Background())

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("Then the job should have been executed", func() {
			So(atomic.LoadInt64(&counter), ShouldEqual, 1)
		})
	})

	Convey("Given I call a wrapped job once and it fails", t, func() {

		var counter int64

		f := AtomicJob(func() error {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			return fmt.Errorf("oh noes")
		})

		err := f(context.Background())

		Convey("Then err should be nil", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "oh noes")
		})

		Convey("Then the job should have been executed", func() {
			So(atomic.LoadInt64(&counter), ShouldEqual, 1)
		})
	})

	Convey("Given I call a wrapped job but context is canceled", t, func() {

		var counter int64

		f := AtomicJob(func() error {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			return nil
		})

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := f(ctx)

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "context canceled")
		})

		Convey("Then the job should have been executed", func() {
			So(atomic.LoadInt64(&counter), ShouldEqual, 0)
		})
	})

	Convey("Given I call a wrapped job but context cancels in the middle", t, func() {

		var counter int64

		f := AtomicJob(func() error {
			time.Sleep(3 * time.Second)
			atomic.AddInt64(&counter, 1)
			return nil
		})

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			time.Sleep(300 * time.Millisecond)
			cancel()
		}()

		err := f(ctx)

		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "context canceled")
		})

		Convey("Then the job should have been executed", func() {
			So(atomic.LoadInt64(&counter), ShouldEqual, 0)
		})
	})

	Convey("Given I call a wrapped job twice and it fails", t, func() {

		var counter int64

		f := AtomicJob(func() error {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			return fmt.Errorf("boom")
		})

		var err1 atomic.Value

		ready := make(chan struct{})
		go func() {
			ready <- struct{}{}
			err1.Store(f(context.Background()))
		}()

		// wait for the first trigger to be running
		select {
		case <-ready:
		case <-time.After(1 * time.Second):
			panic("job did not trigger in time")
		}

		// invoke a second marker trigger
		err2 := f(context.Background())

		time.Sleep(300 * time.Millisecond)

		Convey("Then err1 and err2 should not be nil", func() {
			e := err1.Load().(error) // nolint:revive,unchecked-type-assertion
			So(e, ShouldNotBeNil)
			So(e.Error(), ShouldEqual, "boom")

			So(err2, ShouldNotBeNil)
			So(err2.Error(), ShouldEqual, "boom")
		})

		Convey("Then the job should have been executed", func() {
			So(atomic.LoadInt64(&counter), ShouldEqual, 1)
		})
	})

	Convey("Given I call a wrapped job a lot in parallel", t, func() {

		var counter int64
		f := AtomicJob(func() error {
			time.Sleep(time.Second)
			atomic.AddInt64(&counter, 1)
			return nil
		})

		var err1 atomic.Value

		ready := make(chan struct{})
		go func() {
			ready <- struct{}{}
			if e := f(context.Background()); e != nil {
				err1.Store(e)
			}
		}()

		// wait for the first trigger to be running
		select {
		case <-ready:
		case <-time.After(1 * time.Second):
			panic("job did not trigger in time")
		}

		// invoke 100 other random trigger
		for i := 0; i < 100; i++ {
			go f(context.Background()) // nolint
		}

		// invoke a second marker trigger
		err2 := f(context.Background())

		time.Sleep(300 * time.Millisecond)

		Convey("Then err1 and err2 should be nil", func() {
			So(err1.Load(), ShouldBeNil)
			So(err2, ShouldBeNil)
		})

		Convey("Then the job should have been executed", func() {
			So(atomic.LoadInt64(&counter), ShouldEqual, 1)
		})
	})
}
