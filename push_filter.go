package elemental

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// A PushFilter represents an abstract filter for filtering out push notifications.
type PushFilter struct {
	RequestID  string      `json:"rid"`
	EventTypes []EventType `json:"events"`
	Identities []Identity  `json:"identities"`
}

// NewPushFilter returns a new PushFilter.
func NewPushFilter() *PushFilter {

	return &PushFilter{
		RequestID:  uuid.NewV4().String(),
		EventTypes: []EventType{},
		Identities: []Identity{},
	}
}

// IsEventTypeFiltered returns true if the given EventType is not part of the PushFilter.
func (f *PushFilter) IsEventTypeFiltered(t EventType) bool {

	if len(f.EventTypes) == 0 {
		return false
	}

	for _, o := range f.EventTypes {
		if o == t {
			return false
		}
	}

	return true
}

// IsIdentityFiltered returns true if the given Identity is not part of the PushFilter.
func (f *PushFilter) IsIdentityFiltered(idName string) bool {

	if len(f.Identities) == 0 {
		return false
	}

	for _, i := range f.Identities {
		if i.Name == idName {
			return false
		}
	}

	return true
}

// Duplicate duplicates the PushFilter.
func (f *PushFilter) Duplicate() *PushFilter {

	nf := NewPushFilter()

	for _, i := range f.Identities {
		nf.Identities = append(nf.Identities, i)
	}

	for _, t := range f.EventTypes {
		nf.EventTypes = append(nf.EventTypes, t)
	}

	return nf
}

func (f *PushFilter) String() string {

	return fmt.Sprintf("<pushfilter id:%s event-types:%s identities:%s>",
		f.RequestID,
		f.EventTypes,
		f.Identities,
	)
}
