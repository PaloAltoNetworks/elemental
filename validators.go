// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"time"
)

const (
	maximumIntFailFormat             = `Data '%d' of attribute '%s' should be less than %d`
	maximumIntExclusiveFailFormat    = `Data '%d' of attribute '%s' should be less or equal than %d`
	minimumIntFailFormat             = `Data '%d' of attribute '%s' should be greater than %d`
	minimumIntExclusiveFailFormat    = `Data '%d' of attribute '%s' should be greater or equal than %d`
	maximumFloatFailFormat           = `Data '%g' of attribute '%s' should be less than %g`
	maximumFloatExclusiveFailFormat  = `Data '%g' of attribute '%s' should be less or equal than %g`
	minimumFloatFailFormat           = `Data '%g' of attribute '%s' should be greater than %g`
	minimumFloatExclusiveFailFormat  = `Data '%g' of attribute '%s' should be greater or equal than %g`
	maximumLengthFailFormat          = `Data '%s' of attribute '%s' should be less than %d chars long`
	maximumLengthExclusiveFailFormat = `Data '%s' of attribute '%s' should be less or equal than %d chars long`
	minimumLengthFailFormat          = `Data '%s' of attribute '%s' should be greater than %d chars long`
	minimumLengthExclusiveFailFormat = `Data '%s' of attribute '%s' should be greater or equal than %d chars long`
	patternFailFormat                = `Data '%s' of attribute '%s' should match '%s'`
	requiredFloatFailFormat          = `Attribute '%s' is required`
	requiredStringFailFormat         = `Attribute '%s' is required`
	requiredTimeFailFormat           = `Attribute '%s' is required`
	requiredIntFailFormat            = `Attribute '%s' is required`
	requiredExternalFailFormat       = `Attribute '%s' is required`
	floatInListFormat                = `Data '%g' of attribute '%s' is not in list '%g'`
	stringInListFormat               = `Data '%s' of attribute '%s' is not in list '%s'`
	valInMapFormat                   = `Data '%+v' of attribute '%s' is not in map '%+v'`
	intInListFormat                  = `Data '%d' of attribute '%s' is not in list '%d'`
)

// A Validatable is the interface for objects that can be validated.
type Validatable interface {
	Validate() error
}

// ValidateStringInList validates if the string is in the list.
func ValidateStringInList(attribute string, value string, enums []string, autogenerated bool) error {

	if autogenerated && value == "" {
		return nil
	}

	for _, v := range enums {
		if v == value {
			return nil
		}
	}

	err := NewError("Validation Error", fmt.Sprintf(stringInListFormat, value, attribute, enums), "elemental", http.StatusUnprocessableEntity)
	err.Data = map[string]string{"attribute": attribute}
	return err
}

// ValidateValueInMap validates if the string is in the list.
func ValidateValueInMap(attribute string, value interface{}, enums map[interface{}]interface{}, autogenerated bool) error {

	if autogenerated && value == nil {
		return nil
	}

	if _, ok := enums[attribute]; ok {
		return nil
	}

	err := NewError("Validation Error", fmt.Sprintf(valInMapFormat, value, attribute, enums), "elemental", http.StatusUnprocessableEntity)
	err.Data = map[string]string{"attribute": attribute}
	return err
}

// ValidateFloatInList validates if the string is in the list.
func ValidateFloatInList(attribute string, value float64, enums []float64) error {

	for _, v := range enums {
		if v == value {
			return nil
		}
	}

	err := NewError("Validation Error", fmt.Sprintf(floatInListFormat, value, attribute, enums), "elemental", http.StatusUnprocessableEntity)
	err.Data = map[string]string{"attribute": attribute}
	return err
}

// ValidateIntInList validates if the string is in the list.
func ValidateIntInList(attribute string, value int, enums []int) error {

	for _, v := range enums {
		if v == value {
			return nil
		}
	}

	err := NewError("Validation Error", fmt.Sprintf(intInListFormat, value, attribute, enums), "elemental", http.StatusUnprocessableEntity)
	err.Data = map[string]string{"attribute": attribute}
	return err
}

// ValidateRequiredInt validates is the int is set to 0.
func ValidateRequiredInt(attribute string, value int) error {

	if value == 0 {
		err := NewError("Validation Error", fmt.Sprintf(requiredIntFailFormat, attribute), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateRequiredFloat validates is the int is set to 0.
func ValidateRequiredFloat(attribute string, value float64) error {

	if value == 0.0 {
		err := NewError("Validation Error", fmt.Sprintf(requiredFloatFailFormat, attribute), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateRequiredExternal validates if the given value is null or not
func ValidateRequiredExternal(attribute string, value interface{}) error {
	var valueIsNil bool

	if value == nil {
		valueIsNil = true
	}

	if !valueIsNil {
		v := reflect.ValueOf(value)

		switch v.Kind() {
		case reflect.Slice, reflect.Map:
			valueIsNil = v.IsNil() || v.Len() == 0
		default:
			valueIsNil = v.Interface() == reflect.Zero(reflect.TypeOf(v.Interface())).Interface()
		}
	}

	if valueIsNil {
		err := NewError("Validation Error", fmt.Sprintf(requiredExternalFailFormat, attribute), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateMaximumFloat validates a float against a maximum value.
func ValidateMaximumFloat(attribute string, value float64, max float64, exclusive bool) error {

	if !exclusive && value > max {
		err := NewError("Validation Error", fmt.Sprintf(maximumFloatFailFormat, value, attribute, max), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	} else if exclusive && value >= max {
		err := NewError("Validation Error", fmt.Sprintf(maximumFloatExclusiveFailFormat, value, attribute, max), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateMinimumFloat validates a float against a maximum value.
func ValidateMinimumFloat(attribute string, value float64, min float64, exclusive bool) error {

	if !exclusive && value < min {
		err := NewError("Validation Error", fmt.Sprintf(minimumFloatFailFormat, value, attribute, min), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err

	} else if exclusive && value <= min {
		err := NewError("Validation Error", fmt.Sprintf(minimumFloatExclusiveFailFormat, value, attribute, min), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateMaximumInt validates a integer against a maximum value.
func ValidateMaximumInt(attribute string, value int, max int, exclusive bool) error {

	if !exclusive && value > max {
		err := NewError("Validation Error", fmt.Sprintf(maximumIntFailFormat, value, attribute, max), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	} else if exclusive && value >= max {
		err := NewError("Validation Error", fmt.Sprintf(maximumIntExclusiveFailFormat, value, attribute, max), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateMinimumInt validates a integer against a maximum value.
func ValidateMinimumInt(attribute string, value int, min int, exclusive bool) error {

	if !exclusive && value < min {
		err := NewError("Validation Error", fmt.Sprintf(minimumIntFailFormat, value, attribute, min), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	} else if exclusive && value <= min {
		err := NewError("Validation Error", fmt.Sprintf(minimumIntExclusiveFailFormat, value, attribute, min), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateRequiredString validates if the string is empty.
func ValidateRequiredString(attribute string, value string) error {

	if value == "" {
		err := NewError("Validation Error", fmt.Sprintf(requiredStringFailFormat, attribute), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateRequiredTime validates if the time is empty.
func ValidateRequiredTime(attribute string, value time.Time) error {

	var t time.Time
	if value.Equal(t) {
		err := NewError("Validation Error", fmt.Sprintf(requiredTimeFailFormat, attribute), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidatePattern validates a string against a regular expression.
func ValidatePattern(attribute string, value string, pattern string, required bool) error {

	if !required && value == "" {
		return nil
	}

	re := regexp.MustCompile(pattern)

	if !re.MatchString(value) {
		err := NewError("Validation Error", fmt.Sprintf(patternFailFormat, value, attribute, pattern), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateMinimumLength validates the minimum length of a string.
func ValidateMinimumLength(attribute string, value string, min int, exclusive bool) error {

	length := len([]rune(value))

	if !exclusive && length < min {
		err := NewError("Validation Error", fmt.Sprintf(minimumLengthFailFormat, value, attribute, min), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	} else if exclusive && length <= min {
		err := NewError("Validation Error", fmt.Sprintf(minimumLengthExclusiveFailFormat, value, attribute, min), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}

// ValidateMaximumLength validates the maximum length of a string.
func ValidateMaximumLength(attribute string, value string, max int, exclusive bool) error {

	length := len([]rune(value))

	if !exclusive && length > max {
		err := NewError("Validation Error", fmt.Sprintf(maximumLengthFailFormat, value, attribute, max), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	} else if exclusive && length >= max {
		err := NewError("Validation Error", fmt.Sprintf(maximumLengthExclusiveFailFormat, value, attribute, max), "elemental", http.StatusUnprocessableEntity)
		err.Data = map[string]string{"attribute": attribute}
		return err
	}

	return nil
}
