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

import "go.aporeto.io/elemental"

var (
	identityNamesMap = map[string]elemental.Identity{
		"list": ListIdentity,
		"root": RootIdentity,
		"task": TaskIdentity,
		"user": UserIdentity,
	}

	identitycategoriesMap = map[string]elemental.Identity{
		"lists": ListIdentity,
		"root":  RootIdentity,
		"tasks": TaskIdentity,
		"users": UserIdentity,
	}

	aliasesMap = map[string]elemental.Identity{
		"lst": ListIdentity,
		"tsk": TaskIdentity,
		"usr": UserIdentity,
	}

	indexesMap = map[string][][]string{
		"list": nil,
		"root": nil,
		"task": nil,
		"user": nil,
	}
)

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

type modelManager struct{}

func (f modelManager) IdentityFromName(name string) elemental.Identity {

	return identityNamesMap[name]
}

func (f modelManager) IdentityFromCategory(category string) elemental.Identity {

	return identitycategoriesMap[category]
}

func (f modelManager) IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

func (f modelManager) IdentityFromAny(any string) (i elemental.Identity) {

	if i = f.IdentityFromName(any); !i.IsEmpty() {
		return i
	}

	if i = f.IdentityFromCategory(any); !i.IsEmpty() {
		return i
	}

	return f.IdentityFromAlias(any)
}

func (f modelManager) Identifiable(identity elemental.Identity) elemental.Identifiable {

	switch identity {

	case ListIdentity:
		return NewList()
	case RootIdentity:
		return NewRoot()
	case TaskIdentity:
		return NewTask()
	case UserIdentity:
		return NewUser()
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiable(identity elemental.Identity) elemental.SparseIdentifiable {

	switch identity {

	case ListIdentity:
		return NewSparseList()
	case TaskIdentity:
		return NewSparseTask()
	case UserIdentity:
		return NewSparseUser()
	default:
		return nil
	}
}

func (f modelManager) Indexes(identity elemental.Identity) [][]string {

	return indexesMap[identity.Name]
}

func (f modelManager) IdentifiableFromString(any string) elemental.Identifiable {

	return f.Identifiable(f.IdentityFromAny(any))
}

func (f modelManager) Identifiables(identity elemental.Identity) elemental.Identifiables {

	switch identity {

	case ListIdentity:
		return &ListsList{}
	case TaskIdentity:
		return &TasksList{}
	case UserIdentity:
		return &UsersList{}
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiables(identity elemental.Identity) elemental.SparseIdentifiables {

	switch identity {

	case ListIdentity:
		return &SparseListsList{}
	case TaskIdentity:
		return &SparseTasksList{}
	case UserIdentity:
		return &SparseUsersList{}
	default:
		return nil
	}
}

func (f modelManager) IdentifiablesFromString(any string) elemental.Identifiables {

	return f.Identifiables(f.IdentityFromAny(any))
}

func (f modelManager) Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

var manager = modelManager{}

// Manager returns the model elemental.ModelManager.
func Manager() elemental.ModelManager { return manager }

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		ListIdentity,
		RootIdentity,
		TaskIdentity,
		UserIdentity,
	}
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case ListIdentity:
		return []string{
			"lst",
		}
	case RootIdentity:
		return []string{}
	case TaskIdentity:
		return []string{
			"tsk",
		}
	case UserIdentity:
		return []string{
			"usr",
		}
	}

	return nil
}
