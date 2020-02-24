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

	// parsedIdentityFilters holds the parsed `IdentityFilters` to avoid re-parsing the configured filters on each push
	// event that is using the same config.
	parsedIdentityFilters map[string]*Filter
}

// unsupportedComparators is the list of comparators that are currently not handled by the elemental API 'MatchesFilter'
// which is utilized for providing fine-grained identity filtering for websocket clients
var unsupportedComparators []FilterComparator = []FilterComparator{
	GreaterComparator,
	GreaterOrEqualComparator,
	LesserComparator,
	LesserOrEqualComparator,
	InComparator,
	NotInComparator,
	ContainComparator,
	NotContainComparator,
}

// NewPushFilter returns a new PushFilter. NewPushFilter is now aliased to NewPushConfig. This was done for backwards
// compatibility as a result of the re-naming of PushFilter to PushConfig.
//
// Deprecated: use the constructor with the new name, NewPushConfig, instead
func NewPushFilter() *PushFilter {
	fmt.Println("DEPRECATED: elemental.NewPushFilter is deprecated, use elemental.NewPushConfig instead")

	return &PushFilter{
		Identities:            map[string][]EventType{},
		IdentityFilters:       map[string]string{},
		parsedIdentityFilters: map[string]*Filter{},
	}
}

// NewPushConfig returns a new PushConfig.
func NewPushConfig() *PushConfig {

	return &PushConfig{
		Identities:            map[string][]EventType{},
		IdentityFilters:       map[string]string{},
		parsedIdentityFilters: map[string]*Filter{},
	}
}

// SetParameter sets the values of the parameter with the given key.
func (pc *PushConfig) SetParameter(key string, values ...string) {

	if pc.Params == nil {
		pc.Params = url.Values{}
	}

	pc.Params[key] = values
}

// Parameters returns a copy of all the parameters.
func (pc *PushConfig) Parameters() url.Values {

	if pc.Params == nil {
		return nil
	}

	out := make(url.Values, len(pc.Params))
	for k, v := range pc.Params {
		out[k] = v
	}

	return out
}

// FilterIdentity adds the given identity for the given eventTypes in the PushConfig.
func (pc *PushConfig) FilterIdentity(identityName string, eventTypes ...EventType) {

	pc.Identities[identityName] = eventTypes
}

// ParseIdentityFilters parses the configured PushConfig's 'IdentityFilters' attribute to elemental filters.
// The parsed filters will then be stored in the non-exposed 'parsedIdentityFilters' attribute of PushConfig. This is useful
// for clients that wish the utilize the same filter multiple times without having to incur the overhead of parsing each time.
//
// An error is returned in following situations:
//   - when a filter is declared on an identity that is not defined in the PushConfig's 'Identities' attribute
//   - when a filter cannot be parsed into an elemental.Filter
func (pc *PushConfig) ParseIdentityFilters() error {

	if pc.parsedIdentityFilters == nil {
		pc.parsedIdentityFilters = map[string]*Filter{}
	}

	for identity, unparsedFilter := range pc.IdentityFilters {
		if _, found := pc.Identities[identity]; !found {
			// in the event an error occurs we zero out the parsed identities to avoid having a partially set of parsed identities
			pc.parsedIdentityFilters = map[string]*Filter{}
			return fmt.Errorf("elemental: cannot declare an identity filter on %q as that was not declared in 'Identities'", identity)
		}

		filter, err := NewFilterParser(unparsedFilter,
			// blacklist unsupported comparators so socket can either be closed or if the client supports error events, an
			// error can be emitted.
			OptUnsupportedComparators(unsupportedComparators),
		).Parse()
		if err != nil {
			// in the event an error occurs we zero out the parsed identities to avoid having a partially set of parsed identities
			pc.parsedIdentityFilters = map[string]*Filter{}
			return fmt.Errorf("elemental: unable to parse filter %q: %s", unparsedFilter, err)
		}

		pc.parsedIdentityFilters[identity] = filter
	}

	return nil
}

// IsFilteredOut returns true if the given Identity is not part of the PushConfig's Identity mapping
func (pc *PushConfig) IsFilteredOut(identityName string, eventType EventType) bool {

	// if the identities list is empty, we filter nothing.
	if len(pc.Identities) == 0 {
		return false
	}

	// If it contains something, but not the identity, we filter out.
	types, ok := pc.Identities[identityName]
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

// FilterForIdentity returns the associated fine-grained filter for the given identity. In the event that no fine-grained
// filter has been configured for the identity, the second return value (a boolean), will be set to false.
func (pc *PushConfig) FilterForIdentity(identityName string) (*Filter, bool) {

	if pc.parsedIdentityFilters == nil {
		return nil, false
	}

	filter, found := pc.parsedIdentityFilters[identityName]
	return filter, found
}

// Duplicate duplicates the PushConfig.
func (pc *PushConfig) Duplicate() *PushConfig {

	config := NewPushConfig()

	for id, types := range pc.Identities {
		config.FilterIdentity(id, types...)
	}

	for id, f := range pc.IdentityFilters {
		config.IdentityFilters[id] = f
	}

	for id, f := range pc.parsedIdentityFilters {
		config.parsedIdentityFilters[id] = f
	}

	for k, v := range pc.Params {
		config.SetParameter(k, v...)
	}

	return config
}

func (pc *PushConfig) String() string {

	return fmt.Sprintf("<pushconfig identities:%s identityfilters:%s>", pc.Identities, pc.IdentityFilters)
}
