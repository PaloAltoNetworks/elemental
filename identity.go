// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import "fmt"

// An IdentifiablesList is a list of objects implementing the Identifiable interface.
type IdentifiablesList []Identifiable

// An Identifiable is the interface that Elemental objects must implement.
type Identifiable interface {

	// Identity returns the Identity of the of the receiver.
	Identity() Identity

	// Identifier returns the unique identifier of the of the receiver.
	Identifier() string

	// SetIdentifier sets the unique identifier of the of the receiver.
	SetIdentifier(string)
}

// An Identity is a structure that contains the necessary information about an Identifiable.
// The Name is usually the singular form of the Category.
//
// For instance, "cat" and "cats".
type Identity struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

// MakeIdentity returns a new Identity.
func MakeIdentity(name, category string) Identity {

	return Identity{
		Name:     name,
		Category: category,
	}
}

// String returns the string representation of the identity.
func (i Identity) String() string {

	return fmt.Sprintf("<Identity %s|%s>", i.Name, i.Category)
}

// IsEmpty checks if the identity is empty.
func (i Identity) IsEmpty() bool {

	return i.Name == "" && i.Category == ""
}

// IsEqual checks if the given identity is equal to the receiver.
func (i Identity) IsEqual(identity Identity) bool {

	return i.Name == identity.Name && i.Category == identity.Category
}

// AllIdentity represents all possible Identities.
var AllIdentity = Identity{
	Name:     "__all__",
	Category: "__all__",
}

// EmptyIdentity represents an empty Identity.
var EmptyIdentity = Identity{
	Name:     "",
	Category: "",
}
