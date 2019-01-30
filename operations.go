// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"fmt"
)

// Operation represents an operation to apply on an Identifiable
// from a Request.
type Operation string

// Here are the existing Operations.
const (
	OperationRetrieveMany Operation = "retrieve-many"
	OperationRetrieve     Operation = "retrieve"
	OperationCreate       Operation = "create"
	OperationUpdate       Operation = "update"
	OperationDelete       Operation = "delete"
	OperationPatch        Operation = "patch"
	OperationInfo         Operation = "info"

	OperationEmpty Operation = ""
)

// ParseOperation parses the given string as an Operation.
func ParseOperation(op string) (Operation, error) {

	lop := Operation(op)

	if lop == OperationRetrieveMany ||
		lop == OperationRetrieve ||
		lop == OperationCreate ||
		lop == OperationUpdate ||
		lop == OperationDelete ||
		lop == OperationPatch ||
		lop == OperationInfo {
		return lop, nil
	}

	return Operation(""), fmt.Errorf("invalid operation '%s'", op)
}
