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
)

// A PushFilter represents an abstract filter for filtering out push notifications. This is now aliased to PushConfig as a
// result of re-naming the type.
//
// Deprecated: use the new name PushConfig instead
type PushFilter = PushConfig

// NewPushFilter returns a new PushFilter. NewPushFilter is now aliased to NewPushConfig. This was done for backwards
// compatibility as a result of the re-naming of PushFilter to PushConfig.
//
// Deprecated: use the new name NewPushConfig instead
var NewPushFilter = NewPushConfig

// A PushConfig represents an abstract filter for filtering out push notifications.
//
// The 'IdentityFilters' field is a mapping between a filtered identity and the string representation of an elemental.Filter.
// A client will supply this attribute if they want fine-grained filtering on the set of identities that they are filtering on.
// If this attribute has been supplied, the identities passed to 'IdentityFilters' must be a subset of the identities passed to
// 'Identities'; passing in identities that are not provided in the 'Identities' field will be ignored.
type PushConfig struct {
	Identities      map[string][]EventType `msgpack:"identities" json:"identities"`
	IdentityFilters map[string]string      `msgpack:"filters"    json:"filters"`
	Params          url.Values             `msgpack:"parameters" json:"parameters"`
}

// NewPushConfig returns a new PushConfig.
func NewPushConfig() *PushConfig {

	return &PushConfig{
		Identities: map[string][]EventType{},
	}
}

// SetParameter sets the values of the parameter with the given key.
func (f *PushConfig) SetParameter(key string, values ...string) {

	if f.Params == nil {
		f.Params = url.Values{}
	}

	f.Params[key] = values
}

// Parameters returns a copy of all the parameters.
func (f *PushConfig) Parameters() url.Values {

	if f.Params == nil {
		return nil
	}

	out := url.Values{}
	for k, v := range f.Params {
		out[k] = v
	}

	return out
}

// FilterIdentity adds the given identity for the given eventTypes in the PushConfig.
func (f *PushConfig) FilterIdentity(identityName string, eventTypes ...EventType) {

	f.Identities[identityName] = eventTypes
}

// IsFilteredOut returns true if the given Identity is not part of the PushConfig.
func (f *PushConfig) IsFilteredOut(identityName string, eventType EventType) bool {

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

// Duplicate duplicates the PushConfig.
func (f *PushConfig) Duplicate() *PushConfig {

	nf := NewPushConfig()

	for id, types := range f.Identities {
		nf.FilterIdentity(id, types...)
	}

	for k, v := range f.Params {
		nf.SetParameter(k, v...)
	}

	return nf
}

func (f *PushConfig) String() string {

	return fmt.Sprintf("<pushfilter identities:%s>", f.Identities)
}
