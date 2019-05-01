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

// An ModelManager is the interface that allows to search Identities
// and create Identifiable and Identifiables from Identities.
type ModelManager interface {

	// Identifiable returns an Identifiable with the given identity.
	Identifiable(Identity) Identifiable

	// SparseIdentifiable returns a SparseIdentifiable with the given identity.
	SparseIdentifiable(Identity) SparseIdentifiable

	// IdentifiableFromString returns an Identifiable from the given
	// string. The string can be an Identity name, category or alias.
	IdentifiableFromString(string) Identifiable

	// Identifiables returns an Identifiables with the given identity.
	Identifiables(Identity) Identifiables

	// SparseIdentifiables returns an Identifiables with the given identity.
	SparseIdentifiables(Identity) SparseIdentifiables

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
