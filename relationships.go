package elemental

import (
	"fmt"
	"net/http"
	"net/url"
)

// A RelationshipInfo describe the various meta information of a relationship.
type RelationshipInfo struct {
	Deprecated bool
	Parameters []Parameter

	parameterMap map[string]*Parameter
}

// Build builds various internal representation of the info contained in the structure.
func (i *RelationshipInfo) Build() *RelationshipInfo {

	i.parameterMap = map[string]*Parameter{}
	for _, p := range i.Parameters {
		i.parameterMap[p.Name] = &p
	}

	return i
}

// GetParameter returns the parameter with the given name
func (i *RelationshipInfo) GetParameter(name string) *Parameter {
	return i.parameterMap[name]
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

// ValidateParameters validates the given url.Values for the given relationship information.
func ValidateParameters(registry RelationshipsRegistry, i Identity, pid Identity, op Operation, values url.Values) (err error) {

	r, ok := registry[i]
	if !ok {
		return nil
	}

	var ri *RelationshipInfo

	switch op {
	case OperationCreate:
		ri = r.Create[pid.Name]
	case OperationDelete:
		ri = r.Delete[pid.Name]
	case OperationInfo:
		ri = r.Info[pid.Name]
	case OperationPatch:
		ri = r.Patch[pid.Name]
	case OperationRetrieve:
		ri = r.Retrieve[pid.Name]
	case OperationRetrieveMany:
		ri = r.RetrieveMany[pid.Name]
	case OperationUpdate:
		ri = r.Update[pid.Name]
	}

	for k, v := range values {

		qp := ri.GetParameter(k)

		// If the parameter is not specified, this is an error.
		if qp != nil {
			return NewError("Invalid Parameter", fmt.Sprintf("Parameter '%s' is invalid", k), "elemental", http.StatusBadRequest)
		}

		// Otherwise we validate the parameter.
		if err = qp.Parse(v); err != nil {
			return err
		}
	}

	return nil
}

// IsRetrieveAllowed returns true if retrieving the given identity is allowed.
func IsRetrieveAllowed(registry RelationshipsRegistry, i Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.Retrieve[RootIdentity.Name] != nil
}

// IsUpdateAllowed returns true if updating the given identity is allowed.
func IsUpdateAllowed(registry RelationshipsRegistry, i Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.Update[RootIdentity.Name] != nil
}

// IsDeleteAllowed returns true if deleting the given identity is allowed.
func IsDeleteAllowed(registry RelationshipsRegistry, i Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.Delete[RootIdentity.Name] != nil
}

// IsRetrieveManyAllowed returns true if retrieving many children with the given identity under the parentIdentity is allowed.
func IsRetrieveManyAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.RetrieveMany[pid.Name] != nil
}

// IsInfoAllowed returns true if retrieving info on children with the given identity under the parentIdentity is allowed.
func IsInfoAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.Info[pid.Name] != nil
}

// IsPatchAllowed returns true if patching children with the given identity under the parentIdentity is allowed.
func IsPatchAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.Patch[pid.Name] != nil
}

// IsCreateAllowed returns true if creating the given identity under the parentIdentity is allowed.
func IsCreateAllowed(registry RelationshipsRegistry, i Identity, pid Identity) bool {

	r, ok := registry[i]
	if !ok {
		return false
	}

	return r.Create[pid.Name] != nil
}

// ParametersForOperation returns the parameters defined for the retrieve operation on the given identity.
func ParametersForOperation(registry RelationshipsRegistry, i Identity, pid Identity, op Operation) []Parameter {

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
