package elemental

import "fmt"

const (
	// KindUnsupportedComparator represents a MatcherError kind for when an unsupported comparator has been used
	KindUnsupportedComparator MatcherErrorKind = iota + 1
)

// MatcherErrorKind represents the kind of matcher error that has occurred, client's can use this in their error handling
// to deal with specific error cases
type MatcherErrorKind int

func (k MatcherErrorKind) String() string {
	switch k {
	case KindUnsupportedComparator:
		return "KindUnsupportedComparator"
	default:
		return "UnknownMatcherErrorKind"
	}
}

// MatcherError is the error type that will be returned by elemental.MatchesFilter in the event that it returns an error
type MatcherError struct {
	kind        MatcherErrorKind
	description string
}

func (me *MatcherError) Error() string {
	return fmt.Sprintf("elemental: %s - kind: %s", me.description, me.kind)
}

// Kind returns the MatcherErrorKind for the matcher error
func (me *MatcherError) Kind() MatcherErrorKind {
	return me.kind
}
