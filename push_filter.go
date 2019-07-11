// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elemental

import (
	"fmt"
	"net/url"
	"sync"
)

// A PushFilter represents an abstract filter for filtering out push notifications.
type PushFilter struct {
	Identities map[string][]EventType `msgpack:"identities" json:"identities"`
	Params     url.Values             `msgpack:"parameters" json:"parameters"`

	sync.RWMutex
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

	f.Lock()
	defer f.Unlock()

	f.Identities[identityName] = eventTypes
}

// IsFilteredOut returns true if the given Identity is not part of the PushFilter.
func (f *PushFilter) IsFilteredOut(identityName string, eventType EventType) bool {

	// if the identities list is empty, we filter nothing.
	f.RLock()
	if len(f.Identities) == 0 {
		f.RUnlock()
		return false
	}

	// If it contains something, but not the identity, we filter out.
	types, ok := f.Identities[identityName]
	f.RUnlock()

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

	f.RLock()
	for id, types := range f.Identities {
		nf.FilterIdentity(id, types...)
	}
	f.RUnlock()

	for k, v := range f.Params {
		nf.SetParameter(k, v...)
	}

	return nf
}

func (f *PushFilter) String() string {

	f.RLock()
	defer f.RUnlock()

	return fmt.Sprintf("<pushfilter identities:%s>", f.Identities)
}
