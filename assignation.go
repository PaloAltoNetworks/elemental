// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import "fmt"

// InternalAssignationIdentity is the Identity of an assignation
var InternalAssignationIdentity = Identity{
	Name:     "__internal_assignation__",
	Category: "__internal_assignation__",
}

// Operation represents an partial operation.
type Operation string

const (
	// OperationSet represents an full an complete assignatuon
	OperationSet Operation = "set"

	// OperationAdditive represents a partial additive assignation.
	OperationAdditive = "additive"

	// OperationSubstractive represents a partial Substractive assignation.
	OperationSubstractive = "substractive"
)

// Assignation represents an abstract assignation between two Element objects.
type Assignation struct {
	MembersIdentity Identity  `json:"membersIdentity"`
	IDs             []string  `json:"IDs"`
	Operation       Operation `json:"operation"`
}

// NewAssignation returns a new Assignation.
func NewAssignation(operation Operation, membersIdentity Identity, members ...Identifiable) *Assignation {

	var ids []string
	for _, member := range members {
		if i := member.Identifier(); i == "" {
			panic("Cannot create an assignation object with member with Identifier")
		}
		ids = append(ids, member.Identifier())
	}

	return &Assignation{
		MembersIdentity: membersIdentity,
		IDs:             ids,
		Operation:       operation,
	}
}

// Identity returns the Identity of the of the receiver.
func (a *Assignation) Identity() Identity { return InternalAssignationIdentity }

// Identifier returns the unique identifier of the of the receiver.
func (a *Assignation) Identifier() string { return "__internal__" }

// SetIdentifier sets the unique identifier of the of the receiver.
func (a *Assignation) SetIdentifier(string) {}

func (a *Assignation) String() string {

	return fmt.Sprintf("<Assignation type:%s identity:%s ids:%v>", a.Operation, a.MembersIdentity.Name, a.IDs)
}

// Validate valides the current information stored into the structure.
func (a *Assignation) Validate() Errors {

	errors := Errors{}

	if err := ValidateStringInList("operation", string(a.Operation), []string{"full", "additive", "substractive"}); err != nil {
		errors = append(errors, err)
	}

	return errors
}
