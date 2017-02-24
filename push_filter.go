package elemental

import "fmt"

// A PushFilter represents an abstract filter for filtering out push notifications.
type PushFilter struct {
	Identities map[string][]EventType `json:"identities"`
}

// NewPushFilter returns a new PushFilter.
func NewPushFilter() *PushFilter {

	return &PushFilter{
		Identities: map[string][]EventType{},
	}
}

// FilterIdentity adds the given identity for the given eventTypes in the PushFilter.
func (f *PushFilter) FilterIdentity(identityName string, eventTypes ...EventType) {

	f.Identities[identityName] = eventTypes
}

// IsFilteredOut returns true if the given Identity is not part of the PushFilter.
func (f *PushFilter) IsFilteredOut(identityName string, eventType EventType) bool {

	if len(f.Identities) == 0 {
		return false
	}

	types, ok := f.Identities[identityName]
	if !ok {
		return true
	}

	if types == nil || len(types) == 0 {
		return false
	}

	for _, t := range types {
		if t == eventType {
			return false
		}
	}

	return true
}

// Duplicate duplicates the PushFilter.
func (f *PushFilter) Duplicate() *PushFilter {

	nf := NewPushFilter()

	for id, types := range f.Identities {
		nf.FilterIdentity(id, types...)
	}

	return nf
}

func (f *PushFilter) String() string {

	return fmt.Sprintf("<pushfilter identities:%s>", f.Identities)
}
