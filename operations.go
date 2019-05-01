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
