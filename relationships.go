package elemental

// A RelationshipsRegistry maintains the relationship for Identities.
type RelationshipsRegistry map[Identity]*Relationship

// A Relationship describes the hierarchical relationship of the models.
type Relationship struct {
	Type string

	AllowsRetrieve     map[string]bool
	AllowsRetrieveMany map[string]bool
	AllowsInfo         map[string]bool
	AllowsCreate       map[string]bool
	AllowsUpdate       map[string]bool
	AllowsDelete       map[string]bool
	AllowsPatch        map[string]bool

	// Children RelationshipsRegistry
}

// // AddChild adds a new child Relationship.
// func (r *Relationship) AddChild(identity Identity, relationship *Relationship) {
//
// 	if r.Children == nil {
// 		r.Children = RelationshipsRegistry{}
// 	}
//
// 	r.Children[identity] = relationship
// }

// IsRetrieveAllowed returns true if retrieving the given identity is allowed.
func IsRetrieveAllowed(registry RelationshipsRegistry, i Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsRetrieve[RootIdentity.Name]
}

// IsUpdateAllowed returns true if updating the given identity is allowed.
func IsUpdateAllowed(registry RelationshipsRegistry, i Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsUpdate[RootIdentity.Name]
}

// IsDeleteAllowed returns true if deleting the given identity is allowed.
func IsDeleteAllowed(registry RelationshipsRegistry, i Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsDelete[RootIdentity.Name]
}

// IsRetrieveManyAllowed returns true if retrieving many children with the given identity under the parentIdentity is allowed.
func IsRetrieveManyAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsRetrieveMany[pid.Name]
}

// IsInfoAllowed returns true if retrieving info on children with the given identity under the parentIdentity is allowed.
func IsInfoAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsInfo[pid.Name]
}

// IsPatchAllowed returns true if patching children with the given identity under the parentIdentity is allowed.
func IsPatchAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsPatch[pid.Name]
}

// IsCreateAllowed returns true if creating the given identity under the parentIdentity is allowed.
func IsCreateAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsCreate[pid.Name]
}
