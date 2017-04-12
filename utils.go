// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"reflect"
	"time"
)

// extractFieldNames returns all the field Name of the given
// object using reflection.
func extractFieldNames(obj interface{}) []string {

	val := reflect.ValueOf(obj).Elem()
	c := val.NumField()
	fields := make([]string, c)

	for i := 0; i < c; i++ {
		fields[i] = val.Type().Field(i).Name
	}

	return fields
}

// areFieldValuesEqual checks if the value of the given field name are
// equal in both given objects using reflection.
func areFieldValuesEqual(field string, o1, o2 interface{}) bool {

	field1 := reflect.ValueOf(o1).Elem().FieldByName(field)
	field2 := reflect.ValueOf(o2).Elem().FieldByName(field)

	// This is to handle time structure whatever their timezone
	if field1.Type() == reflect.ValueOf(time.Now()).Type() {
		return field1.Interface().(time.Time).Unix() == field2.Interface().(time.Time).Unix()
	}

	if isFieldValueZero(field, o1) && isFieldValueZero(field, o2) {
		return true
	}

	if field1.Kind() == reflect.Slice || field1.Kind() == reflect.Array {
		if field1.Len() != field2.Len() {
			return false
		}

		return reflect.DeepEqual(field1.Interface(), field2.Interface())
	}

	return field1.Interface() == field2.Interface()
}

// isFieldValueZero check if the value of the given field is set to its zero value.
func isFieldValueZero(field string, o interface{}) bool {

	v := reflect.ValueOf(o).Elem().FieldByName(field)

	var defaultTime time.Time
	if v.Type() == reflect.TypeOf(defaultTime) {
		return defaultTime.Equal(v.Interface().(time.Time))
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() == 0
	default:
		return v.Interface() == reflect.Zero(reflect.TypeOf(v.Interface())).Interface()
	}
}
