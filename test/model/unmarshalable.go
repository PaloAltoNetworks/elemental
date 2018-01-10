package testmodel

import (
	"fmt"

	"github.com/aporeto-inc/elemental"
)

// UnmarshalableListIdentity represents the Identity of the object.
var UnmarshalableListIdentity = elemental.Identity{Name: "list", Category: "lists"}

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
