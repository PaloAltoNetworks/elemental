// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

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
)
