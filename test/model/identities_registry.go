package testmodel

import "github.com/aporeto-inc/elemental"

func init() {

	elemental.RegisterIdentity(ListIdentity)
	elemental.RegisterIdentity(TaskIdentity)
	elemental.RegisterIdentity(RootIdentity)
	elemental.RegisterIdentity(UserIdentity)
}

// ModelVersion returns the current version of the model
func ModelVersion() float64 { return 1.0 }

// IdentifiableForIdentity returns a new instance of the Identifiable for the given identity name.
func IdentifiableForIdentity(identity string) elemental.Identifiable {

	switch identity {
	case ListIdentity.Name:
		return NewList()
	case TaskIdentity.Name:
		return NewTask()
	case RootIdentity.Name:
		return NewRoot()
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
	case TaskIdentity.Category:
		return NewTask()
	case RootIdentity.Category:
		return NewRoot()
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
		TaskIdentity,
		RootIdentity,
		UserIdentity,
	}
}

var aliasesMap = map[string]elemental.Identity{}

// IdentityFromAlias returns the Identity associated to the given alias.
func IdentityFromAlias(alias string) elemental.Identity {

	return aliasesMap[alias]
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity elemental.Identity) []string {

	switch identity {
	case ListIdentity:
		return []string{}
	case TaskIdentity:
		return []string{}
	case RootIdentity:
		return []string{}
	case UserIdentity:
		return []string{}
	}

	return nil
}
