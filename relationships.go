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

// A RelationshipInfo describe the various meta information of a relationship.
type RelationshipInfo struct {
	Deprecated         bool
	Parameters         []ParameterDefinition
	RequiredParameters ParametersRequirement
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

// RelationshipInfoForOperation returns the relationship info for the given identity, parent identity and operation.
func RelationshipInfoForOperation(registry RelationshipsRegistry, i Identity, pid Identity, op Operation) *RelationshipInfo {

	r, ok := registry[i]
	if !ok {
		return nil
	}

	switch op {
	case OperationCreate:
		return r.Create[pid.Name]
	case OperationDelete:
		return r.Delete[pid.Name]
	case OperationInfo:
		return r.Info[pid.Name]
	case OperationPatch:
		return r.Patch[pid.Name]
	case OperationRetrieve:
		return r.Retrieve[pid.Name]
	case OperationRetrieveMany:
		return r.RetrieveMany[pid.Name]
	case OperationUpdate:
		return r.Update[pid.Name]
	}

	return nil
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
	default:
		return false
	}

	return ok
}

// ParametersForOperation returns the parameters defined for the retrieve operation on the given identity.
func ParametersForOperation(registry RelationshipsRegistry, i Identity, pid Identity, op Operation) []ParameterDefinition {

	rel := RelationshipInfoForOperation(registry, i, pid, op)
	if rel == nil {
		return nil
	}

	return rel.Parameters
}
