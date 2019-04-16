package elemental

import (
	"fmt"
	"net/url"
)

// A PushFilter represents an abstract filter for filtering out push notifications.
type PushFilter struct {
	Identities map[string][]EventType `msgpack:"identities" json:"identities"`
	Params     url.Values             `msgpack:"parameters" json:"parameters"`
}

// NewPushFilter returns a new PushFilter.
func NewPushFilter() *PushFilter {

	return &PushFilter{
		Identities: map[string][]EventType{},
	}
}

// SetParameter sets the values of the parameter with the given key.
func (f *PushFilter) SetParameter(key string, values ...string) {

	if f.Params == nil {
		f.Params = url.Values{}
	}

	f.Params[key] = values
}

// Parameters returns a copy of all the parameters.
func (f *PushFilter) Parameters() url.Values {

	if f.Params == nil {
		return nil
	}

	out := url.Values{}
	for k, v := range f.Params {
		out[k] = v
	}

	return out
}

// FilterIdentity adds the given identity for the given eventTypes in the PushFilter.
func (f *PushFilter) FilterIdentity(identityName string, eventTypes ...EventType) {

	f.Identities[identityName] = eventTypes
}

// IsFilteredOut returns true if the given Identity is not part of the PushFilter.
func (f *PushFilter) IsFilteredOut(identityName string, eventType EventType) bool {

	// if the identities list is empty, we filter nothing.
	if len(f.Identities) == 0 {
		return false
	}

	// If it contains something, but not the identity, we filter out.
	types, ok := f.Identities[identityName]
	if !ok {
		return true
	}

	// If there is no event types defined we don't filter
	if len(types) == 0 {
		return false
	}

	// If if there are some event types defined, we don't filter out
	// if the current event type is in the list.
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

	for k, v := range f.Params {
		nf.SetParameter(k, v...)
	}

	return nf
}

func (f *PushFilter) String() string {

	return fmt.Sprintf("<pushfilter identities:%s>", f.Identities)
}
