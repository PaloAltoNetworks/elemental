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

import "fmt"

// An IdentifiablesList is a list of objects implementing the Identifiable interface.
type IdentifiablesList []Identifiable

// Identifiables is the interface of a list of Identifiable that can
// returns the Identity of the objects it contains.
type Identifiables interface {
	Identity() Identity
	List() IdentifiablesList
	Copy() Identifiables
	Append(...Identifiable) Identifiables
	Versionable
}

// A PlainIdentifiables is the interface of an object that can return a sparse
// version of itself.
type PlainIdentifiables interface {

	// ToSparse returns a sparsed version of the object.
	ToSparse(...string) Identifiables

	Identifiables
}

// A SparseIdentifiables is the interface of an object that can return a full
// version of itself.
type SparseIdentifiables interface {

	// ToPlain returns the full version of the object.
	ToPlain() IdentifiablesList

	Identifiables
}

// An Identifiable is the interface that Elemental objects must implement.
type Identifiable interface {

	// Identity returns the Identity of the of the receiver.
	Identity() Identity

	// Identifier returns the unique identifier of the of the receiver.
	Identifier() string

	// SetIdentifier sets the unique identifier of the of the receiver.
	SetIdentifier(string)

	Versionable
}

// A PlainIdentifiable is the interface of an object that can return a sparse
// version of itself.
type PlainIdentifiable interface {

	// ToSparse returns a sparsed version of the object.
	ToSparse(...string) SparseIdentifiable

	Identifiable
}

// A SparseIdentifiable is the interface of an object that can return a full
// version of itself.
type SparseIdentifiable interface {

	// ToPlain returns the full version of the object.
	ToPlain() PlainIdentifiable

	Identifiable
}

// DefaultOrderer is the interface of an object that has default ordering fields.
type DefaultOrderer interface {

	// Default order returns the keys that can be used for default ordering.
	DefaultOrder() []string
}

// AttributeEncryptable is the interface of on object that
// has encryptable
type AttributeEncryptable interface {
	EncryptAttributes(encrypter AttributeEncrypter) error
	DecryptAttributes(encrypter AttributeEncrypter) error
}

// An Identity is a structure that contains the necessary information about an Identifiable.
// The Name is usually the singular form of the Category.
//
// For instance, "cat" and "cats".
type Identity struct {
	Name     string `msgpack:"name" json:"name"`
	Category string `msgpack:"category" json:"category"`
	Private  bool   `msgpack:"-" json:"-"`
	Package  string `msgpack:"-" json:"-"`
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

// A Documentable is an object that can be documented.
type Documentable interface {
	Doc() string
}

// A Versionable is an object that can be versioned.
type Versionable interface {
	Version() int
}

// A Patchable the interface of an object that can be patched.
type Patchable interface {

	// Patch patches the receiver using the given SparseIdentifiable.
	Patch(SparseIdentifiable)
}
