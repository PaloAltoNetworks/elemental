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

package testmodel

import (
	"fmt"

	"go.aporeto.io/elemental"
)

// UnmarshalableListIdentity represents the Identity of the object.
var UnmarshalableListIdentity = elemental.Identity{Name: "list", Category: "lists"}

// UnmarshalableListsList represents a list of UnmarshalableLists
type UnmarshalableListsList []*UnmarshalableList

// Identity returns the identity of the objects in the list.
func (o UnmarshalableListsList) Identity() elemental.Identity {

	return UnmarshalableListIdentity
}

// Copy returns a pointer to a copy the UnmarshalableListsList.
func (o UnmarshalableListsList) Copy() elemental.Identifiables {

	out := append(UnmarshalableListsList{}, o...)
	return &out
}

// Append appends the objects to the a new copy of the UnmarshalableListsList.
func (o UnmarshalableListsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(UnmarshalableListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*UnmarshalableList))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o UnmarshalableListsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o UnmarshalableListsList) DefaultOrder() []string {

	return []string{
		"flagDefaultOrderingKey",
	}
}

// Version returns the version of the content.
func (o UnmarshalableListsList) Version() int {

	return 1
}

// An UnmarshalableList is a List that cannot be marshalled  or unmarshalled.
type UnmarshalableList struct {
	List
}

// NewUnmarshalableList returns a new UnmarshalableList.
func NewUnmarshalableList() *UnmarshalableList {
	return &UnmarshalableList{List: List{}}
}

// Identity returns the identity.
func (o *UnmarshalableList) Identity() elemental.Identity { return UnmarshalableListIdentity }

// UnmarshalJSON makes the UnmarshalableList not unmarshalable.
func (o *UnmarshalableList) UnmarshalJSON([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalJSON makes the UnmarshalableList not marshalable.
func (o *UnmarshalableList) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// UnmarshalMsgpack makes the UnmarshalableList not unmarshalable.
func (o *UnmarshalableList) UnmarshalMsgpack([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalMsgpack makes the UnmarshalableList not marshalable.
func (o *UnmarshalableList) MarshalMsgpack() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// Validate validates the data
func (o *UnmarshalableList) Validate() elemental.Errors { return nil }

// An UnmarshalableError is a List that cannot be marshalled or unmarshalled.
type UnmarshalableError struct {
	elemental.Error
}

// UnmarshalJSON makes the UnmarshalableError not unmarshalable.
func (o *UnmarshalableError) UnmarshalJSON([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalJSON makes the UnmarshalableError not marshalable.
func (o *UnmarshalableError) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// UnmarshalMsgpack makes the UnmarshalableError not unmarshalable.
func (o *UnmarshalableError) UnmarshalMsgpack([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalMsgpack makes the UnmarshalableError not marshalable.
func (o *UnmarshalableError) MarshalMsgpack() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}
