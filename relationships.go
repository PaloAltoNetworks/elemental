package elemental

// A RelationshipInfo describe the various meta information of a relationship.
type RelationshipInfo struct {
	Deprecated bool
	Parameters []ParameterDefinition
}

// A RelationshipsRegistry maintains the relationship for Identities.
type RelationshipsRegistry map[Identity]*Relationship

// A Relationship describes the hierarchical relationship of the models.
type Relationship struct {
	Type string

	Retrieve     map[string]*RelationshipInfo
	RetrieveMany map[string]*RelationshipInfo
	Info         map[string]*RelationshipInfo
	Create       map[string]*RelationshipInfo
	Update       map[string]*RelationshipInfo
	Delete       map[string]*RelationshipInfo
	Patch        map[string]*RelationshipInfo
}

// IsOperationAllowed returns true if given operatation on the given identity with the given parent is allowed.
func IsOperationAllowed(registry RelationshipsRegistry, i Identity, pid Identity, op Operation) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	switch op {
	case OperationCreate:
		_, ok = r.Create[pid.Name]
	case OperationDelete:
		_, ok = r.Delete[pid.Name]
	case OperationInfo:
		_, ok = r.Info[pid.Name]
	case OperationPatch:
		_, ok = r.Patch[pid.Name]
	case OperationRetrieve:
		_, ok = r.Retrieve[pid.Name]
	case OperationRetrieveMany:
		_, ok = r.RetrieveMany[pid.Name]
	case OperationUpdate:
		_, ok = r.Update[pid.Name]
	}

	return ok
}

// ParametersForOperation returns the parameters defined for the retrieve operation on the given identity.
func ParametersForOperation(registry RelationshipsRegistry, i Identity, pid Identity, op Operation) []ParameterDefinition {

	var ok bool

	r, ok := registry[i]
	if !ok {
		return nil
	}

	var rel *RelationshipInfo

	switch op {
	case OperationCreate:
		rel, ok = r.Create[pid.Name]
	case OperationDelete:
		rel, ok = r.Delete[pid.Name]
	case OperationInfo:
		rel, ok = r.Info[pid.Name]
	case OperationPatch:
		rel, ok = r.Patch[pid.Name]
	case OperationRetrieve:
		rel, ok = r.Retrieve[pid.Name]
	case OperationRetrieveMany:
		rel, ok = r.RetrieveMany[pid.Name]
	case OperationUpdate:
		rel, ok = r.Update[pid.Name]
	}

	if !ok {
		return nil
	}

	return rel.Parameters
}
