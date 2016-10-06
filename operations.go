// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

// Operation represents a Cid operation.
type Operation string

const (
	// OperationRetrieveMany is the operation used to get multiple objects.
	OperationRetrieveMany Operation = "retrieve-many"

	// OperationRetrieve is the operation used to get a single object.
	OperationRetrieve Operation = "retrieve"

	// OperationCreate is the operation used to create a single object.
	OperationCreate Operation = "create"

	// OperationUpdate is the operation used to update a single object.
	OperationUpdate Operation = "update"

	// OperationDelete is the operation used to delete a single object.
	OperationDelete Operation = "delete"

	// OperationPatch is the operation used to patcj a single object.
	OperationPatch Operation = "patch"

	// OperationInfo is the operation used to get info for a single object.
	OperationInfo Operation = "info"
)
