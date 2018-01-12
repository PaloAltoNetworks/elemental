package testmodel

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(ListIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(TaskIdentity)
	elemental.RegisterIdentity(UserIdentity)
}

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {

	case ListIdentity.Name:
		return NewList()
	case RootIdentity.Name:
		return NewRoot()
	case TaskIdentity.Name:
		return NewTask()
	case UserIdentity.Name:
		return NewUser()
	default:
		return nil
	}
}

// IdentifiableForCategory returns a new instance of the Identifiable for the given category name.
func IdentifiableForCategory(category string) elemental.Identifiable {

	switch category {

	case ListIdentity.Category:
		return NewList()
	case RootIdentity.Category:
		return NewRoot()
	case TaskIdentity.Category:
		return NewTask()
	case UserIdentity.Category:
		return NewUser()
	default:
		return nil
	}
}

// ContentIdentifiableForIdentity returns a new instance of a ContentIdentifiable for the given identity name.
func ContentIdentifiableForIdentity(identity string) elemental.ContentIdentifiable {

	switch identity {

	case ListIdentity.Name:
		return &ListsList{}
	case RootIdentity.Name:
		return &RootsList{}
	case TaskIdentity.Name:
		return &TasksList{}
	case UserIdentity.Name:
		return &UsersList{}
	default:
		return nil
	}
}

// ContentIdentifiableForCategory returns a new instance of a ContentIdentifiable for the given category name.
func ContentIdentifiableForCategory(category string) elemental.ContentIdentifiable {

	switch category {

	case ListIdentity.Category:
		return &ListsList{}
	case RootIdentity.Category:
		return &RootsList{}
	case TaskIdentity.Category:
		return &TasksList{}
	case UserIdentity.Category:
		return &UsersList{}
	default:
		return nil
	}
}

// AllIdentities returns all existing identities.
func AllIdentities() []elemental.Identity {

	return []elemental.Identity{
		ListIdentity,
		RootIdentity,
		TaskIdentity,
		UserIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{
	"lst": ListIdentity,
	"tsk": TaskIdentity,
	"usr": UserIdentity,
}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
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
