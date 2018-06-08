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

	copy := append(UnmarshalableListsList{}, o...)
	return &copy
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
