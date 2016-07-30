package elemental

import (
	"fmt"
)

// ValidateAdvancedSpecification verifies advanced specifications attributes like ReadOnly and CreationOnly.
//
// For instance, it will check if the given Manipulable has field marked as
// readonly, that it has not changed according to the db.
func ValidateAdvancedSpecification(obj AttributeSpecifiable, pristine AttributeSpecifiable, op Operation) Errors {

	errors := NewErrors()

	for _, field := range extractFieldNames(obj) {

		spec := obj.SpecificationForAttribute(field)

		switch op {
		case OperationCreate:
			if spec.ReadOnly && !isFieldValueZero(field, obj) {
				errors = append(
					errors,
					NewError(
						"Read Only Error",
						fmt.Sprintf("Field %s is read only. You cannot set its value.", field),
						"specification",
						3001,
					),
				)
			}

		case OperationUpdate:
			if spec.ReadOnly && !areFieldValuesEqual(field, obj, pristine) {
				errors = append(
					errors,
					NewError(
						"Read Only Error",
						fmt.Sprintf("Field %s is read only. You cannot modify its value.", field),
						"specification",
						3001,
					),
				)
			}

			if spec.CreationOnly && !areFieldValuesEqual(field, obj, pristine) {
				errors = append(
					errors,
					NewError(
						"Creation Only Error",
						fmt.Sprintf("Field %s can only be set during creation. You cannot modify its value.", field),
						"specification",
						3001,
					),
				)
			}
		}
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
