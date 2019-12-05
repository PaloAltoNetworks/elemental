package elemental

import (
	"errors"
	"strings"
)

var (
	// ErrUnsupportedComparator is the error type that will be returned in the event that that an unsupported comparator
	// is used in the filter.
	ErrUnsupportedComparator = errors.New("filter is using an unsupported comparator")
)

// MatcherError is the error type that will be returned by elemental.MatchesFilter in the event that it returns an error
type MatcherError struct {
	err   error
	debug string
}

func (me *MatcherError) Error() string {
	var description strings.Builder
	description.WriteString("elemental: " + me.err.Error())

	// add debug copy if it is present
	if me.debug != "" {
		description.WriteString(" - " + me.debug)
	}

	return description.String()
}

// Unwrap returns the the error contained in 'MatcherError'. This is a special method that aids in error handling for clients
// using Go 1.13 and beyond as they can now utilize the new 'Is' function added to the 'errors' package.
func (me *MatcherError) Unwrap() error {
	return me.err
}
