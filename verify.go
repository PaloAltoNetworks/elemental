package elemental

import (
	"fmt"
	"net/http"
	"reflect"
)

// ValidateAdvancedSpecification verifies advanced specifications attributes like ReadOnly and CreationOnly.
//
// For instance, it will check if the given Manipulable has field marked as
// readonly, that it has not changed according to the db.
func ValidateAdvancedSpecification(obj AttributeSpecifiable, pristine AttributeSpecifiable, op Operation) error {

	errors := NewErrors()

	for _, field := range extractFieldNames(obj) {

		spec := obj.SpecificationForAttribute(field)

		// If the field is not exposed, we don't enforce anything.
		if !spec.Exposed {
			continue
		}

		switch op {
		case OperationCreate:
			if spec.ReadOnly && !isFieldValueZero(field, obj) {
				errors = append(
					errors,
					NewError(
						"Read Only Error",
						fmt.Sprintf("Field %s is read only. You cannot set its value.", field),
						"specification",
						http.StatusUnprocessableEntity,
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
						http.StatusUnprocessableEntity,
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
						http.StatusUnprocessableEntity,
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

// BackportUnexposedFields copy the values of unexposed fields from src to dest.
func BackportUnexposedFields(src, dest AttributeSpecifiable) {

	for _, field := range extractFieldNames(src) {

		spec := src.SpecificationForAttribute(field)

		if !spec.Exposed {
			reflect.ValueOf(dest).Elem().FieldByName(field).Set(reflect.ValueOf(src).Elem().FieldByName(field))
		}
	}
}
