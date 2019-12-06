package elemental

import (
	"fmt"
	"reflect"
)

// ErrUnsupportedComparator is the error type that will be returned in the event that that an unsupported comparator
// is used in the filter.
type ErrUnsupportedComparator struct {
	err error
}

// Is reports whether the provided error has the same type as ErrUnsupportedComparator. This was added as part of the new
// error handling APIs added to Go 1.13
func (e ErrUnsupportedComparator) Is(err error) bool {
	return reflect.TypeOf(err) == reflect.TypeOf(e)
}

// Unwrap returns the embedded error in ErrUnsupportedComparator.
func (e ErrUnsupportedComparator) Unwrap() error {
	return e.err
}

func (e ErrUnsupportedComparator) Error() string {
	return e.err.Error()
}

// MatcherError is the error type that will be returned by elemental.MatchesFilter in the event that it returns an error
type MatcherError struct {
	err error
}

func (me *MatcherError) Error() string {
	return fmt.Sprintf("elemental: %s", me.err)
}

// Unwrap returns the the error contained in 'MatcherError'. This is a special method that aids in error handling for clients
// using Go 1.13 and beyond as they can now utilize the new 'Is' function added to the 'errors' package.
func (me *MatcherError) Unwrap() error {
	return me.err
}
