// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import "fmt"

// An ModelManager is the interface that allows to search Identities
// and create Identifiable and Identifiables from Identities.
type ModelManager interface {

	// Identifiable returns an Identifiable with the given identity.
	Identifiable(Identity) Identifiable

	// IdentifiableFromString returns an Identifiable from the given
	// string. The string can be an Identity name, category or alias.
	IdentifiableFromString(string) Identifiable

	// Identifiables returns an Identifiables with the given identity.
	Identifiables(Identity) Identifiables

	// IdentifiablesFrom returns an Identifiables from the given
	// string. The string can be an Identity name, category or alias.
	IdentifiablesFromString(string) Identifiables

	// IdentityFromName returns the Identity from the given name.
	IdentityFromName(string) Identity

	// IdentityFromCategory returns the Identity from the given category.
	IdentityFromCategory(string) Identity

	// IdentityFromAlias returns the Identity from the given alias.
	IdentityFromAlias(string) Identity

	// IdentityFromAny returns the Identity from the given name, category or alias.
	IdentityFromAny(string) Identity

	// IndexesForIdentity returns the indexes of the given Identity.
	Indexes(Identity) [][]string

	// Relationships return the model's elemental.RelationshipsRegistry.
	Relationships() RelationshipsRegistry
}

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

	// Shadow returns the shadowed version of the identifiable.
	// Shadow() Identifiable

	Versionable
}

// DefaultOrderer is the interface of an object that has default ordering fields.
type DefaultOrderer interface {

	// Default order returns the keys that can be used for default ordering.
	DefaultOrder() []string
}

// An Identity is a structure that contains the necessary information about an Identifiable.
// The Name is usually the singular form of the Category.
//
// For instance, "cat" and "cats".
type Identity struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Private  bool   `json:"-"`
	Package  string `json:"-"`
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
	Name:     "*",
	Category: "*",
}

// EmptyIdentity represents an empty Identity.
var EmptyIdentity = Identity{
	Name:     "",
	Category: "",
}

// RootIdentity represents an root Identity.
var RootIdentity = Identity{
	Name:     "root",
	Category: "root",
}

// Identifiables is the interface of a list of Identifiable that can
// returns the Identity of the objects it contains.
type Identifiables interface {
	Identity() Identity
	List() IdentifiablesList
	Copy() Identifiables
	Append(...Identifiable) Identifiables
	Versionable
}

// A Documentable is an object that can be documented.
type Documentable interface {
	Doc() string
}

// A Versionable is an object that can be versioned.
type Versionable interface {
	Version() int
}
