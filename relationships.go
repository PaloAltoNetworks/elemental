package elemental

// A RelationshipsRegistry maintains the relationship for Identities.
type RelationshipsRegistry map[Identity]*Relationship

// A Relationship describes the hierachical relationship of the models.
type Relationship struct {
	Type string

	AllowsRetrieve     bool
	AllowsRetrieveMany bool
	AllowsInfo         bool
	AllowsCreate       bool
	AllowsUpdate       bool
	AllowsDelete       bool
	AllowsPatch        bool

	Children RelationshipsRegistry
}

// AddChild adds a new child Relationship.
func (r *Relationship) AddChild(identity Identity, relationship *Relationship) {

	if r.Children == nil {
		r.Children = RelationshipsRegistry{}
	}

	r.Children[identity] = relationship
}

// IsRetrieveAllowed returns true if retrieving the given identity is allowed.
func IsRetrieveAllowed(registry RelationshipsRegistry, i Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsRetrieve
}

// IsUpdateAllowed returns true if updating the given identity is allowed.
func IsUpdateAllowed(registry RelationshipsRegistry, i Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsUpdate
}

// IsDeleteAllowed returns true if deleting the given identity is allowed.
func IsDeleteAllowed(registry RelationshipsRegistry, i Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.AllowsDelete
}

// IsRetrieveManyAllowed returns true if retrieving many children with the given identity under the parentIdentity is allowed.
func IsRetrieveManyAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[pid]
	if !ok {
		return false
	}

	cr, ok := r.Children[i]
	if !ok {
		return false
	}

	return cr.AllowsRetrieveMany
}

// IsInfoAllowed returns true if retrieving info on children with the given identity under the parentIdentity is allowed.
func IsInfoAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[pid]
	if !ok {
		return false
	}

	cr, ok := r.Children[i]
	if !ok {
		return false
	}

	return cr.AllowsInfo
}

// IsPatchAllowed returns true if patching children with the given identity under the parentIdentity is allowed.
func IsPatchAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[pid]
	if !ok {
		return false
	}

	cr, ok := r.Children[i]
	if !ok {
		return false
	}

	return cr.AllowsPatch
}

// IsCreateAllowed returns true if creating the given identity under the parentIdentity is allowed.
func IsCreateAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[pid]
	if !ok {
		return false
	}

	cr, ok := r.Children[i]
	if !ok {
		return false
	}

	return cr.AllowsCreate
}
