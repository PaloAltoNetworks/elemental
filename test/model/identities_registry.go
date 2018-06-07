package testmodel

import "github.com/aporeto-inc/elemental"

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
)

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

type identifiableFactory struct{}

func (f identifiableFactory) IdentityFromName(name string) elemental.Identity {

	return identityNamesMap[name]
}

func (f identifiableFactory) IdentityFromCategory(category string) elemental.Identity {

	return identitycategoriesMap[category]
}

func (f identifiableFactory) IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

func (f identifiableFactory) IdentityFromAny(any string) (i elemental.Identity) {

	if i = f.IdentityFromName(any); !i.IsEmpty() {
		return i
	}

	if i = f.IdentityFromCategory(any); !i.IsEmpty() {
		return i
	}

	return f.IdentityFromAlias(any)
}

func (f identifiableFactory) Identifiable(identity elemental.Identity) elemental.Identifiable {

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

func (f identifiableFactory) IdentifiableFromString(any string) elemental.Identifiable {

	return f.Identifiable(f.IdentityFromAny(any))
}

func (f identifiableFactory) ContentIdentifiable(identity elemental.Identity) elemental.ContentIdentifiable {

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

func (f identifiableFactory) ContentIdentifiableFromString(any string) elemental.ContentIdentifiable {

	return f.ContentIdentifiable(f.IdentityFromAny(any))
}

func (f identifiableFactory) Relationships() elemental.RelationshipsRegistry {

	return relationshipsRegistry
}

var ifactory = identifiableFactory{}

// Factory returns the model elemental.IdentifiableFactory.
func Factory() elemental.IdentifiableFactory { return ifactory }

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
