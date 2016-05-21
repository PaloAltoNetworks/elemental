// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

var (
	nameIdentitiesRegistry     map[string]Identity
	categoryIdentitiesRegistry map[string]Identity
)

// RegisterIdentity registers the given identity in the registries
func RegisterIdentity(identity Identity) {

	nameIdentitiesRegistry[identity.Name] = identity
	categoryIdentitiesRegistry[identity.Category] = identity
}

// UnregisterIdentity unregisters the given identity in from the registries
func UnregisterIdentity(identity Identity) {

	delete(nameIdentitiesRegistry, identity.Name)
	delete(categoryIdentitiesRegistry, identity.Category)
}

// IdentityFromName returns the Identity registered with the given Name
func IdentityFromName(name string) Identity {

	return nameIdentitiesRegistry[name]
}

// IdentityFromCategory returns the Identity registered with the given Category
func IdentityFromCategory(category string) Identity {

	return categoryIdentitiesRegistry[category]
}

func init() {

	nameIdentitiesRegistry = make(map[string]Identity)
	categoryIdentitiesRegistry = make(map[string]Identity)
}
