// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"reflect"
	"time"
)

var reflectedTimeType = reflect.ValueOf(time.Time{}).Type()

// ExtractFieldNames returns all the field Name of the given
// object using reflection.
func ExtractFieldNames(obj interface{}) []string {

	val := reflect.Indirect(reflect.ValueOf(obj))
	c := val.NumField()
	fields := make([]string, c)

	for i := 0; i < c; i++ {
		fields[i] = val.Type().Field(i).Name
	}

	return fields
}

// AreFieldValuesEqual checks if the value of the given field name are
// equal in both given objects using reflection.
func AreFieldValuesEqual(field string, o1, o2 interface{}) bool {

	field1 := reflect.Indirect(reflect.ValueOf(o1)).FieldByName(field)
	field2 := reflect.Indirect(reflect.ValueOf(o2)).FieldByName(field)

	if IsFieldValueZero(field, o1) && IsFieldValueZero(field, o2) {
		return true
	}

	// This is to handle time structure whatever their timezone
	if field1.Type() == reflectedTimeType {
		return field1.Interface().(time.Time).Unix() == field2.Interface().(time.Time).Unix()
	}

	if field1.Kind() == reflect.Slice || field1.Kind() == reflect.Array {

		if field1.Len() != field2.Len() {
			return false
		}

		// Same stuff we need to check all time element.
		if field1.Type().Elem() == reflectedTimeType {
			for i := 0; i < field1.Len(); i++ {
				if field1.Index(i).Interface().(time.Time).Unix() != field2.Index(i).Interface().(time.Time).Unix() {
					return false
				}
			}
			return true
		}

		return reflect.DeepEqual(field1.Interface(), field2.Interface())
	}

	if field1.Kind() == reflect.Map {

		if field1.Len() != field2.Len() {
			return false
		}

		return reflect.DeepEqual(field1.Interface(), field2.Interface())
	}

	return field1.Interface() == field2.Interface()
}

// IsFieldValueZero check if the value of the given field is set to its zero value.
func IsFieldValueZero(field string, o interface{}) bool {

	v := reflect.Indirect(reflect.ValueOf(o)).FieldByName(field)

	return IsZero(v.Interface())
}

// IsZero returns true is the given value is zero.
func IsZero(o interface{}) bool {

	v := reflect.Indirect(reflect.ValueOf(o))

	if v.Type() == reflectedTimeType {
		return time.Time{}.Equal(v.Interface().(time.Time))
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() == 0
	default:
		return v.Interface() == reflect.Zero(reflect.TypeOf(v.Interface())).Interface()
	}
}

// AreFieldsValueEqualValue checks field with the given name of the given
// object is equal to the given value.
func AreFieldsValueEqualValue(f string, obj interface{}, value interface{}) bool {

	field := reflect.Indirect(reflect.ValueOf(obj)).FieldByName(f)

	if value == nil {
		return IsFieldValueZero(f, obj)
	}

	v2 := reflect.ValueOf(value)

	// This is to handle time structure whatever their timezone
	if field.Type() == reflect.ValueOf(time.Now()).Type() {
		return field.Interface().(time.Time).Unix() == v2.Interface().(time.Time).Unix()
	}

	if field.Kind() == reflect.Slice || field.Kind() == reflect.Array {
		if field.Len() != v2.Len() {
			return false
		}

		return reflect.DeepEqual(field.Interface(), v2.Interface())
	}

	if field.Kind() == reflect.Map {
		return reflect.DeepEqual(field.Interface(), v2.Interface())
	}

	return field.Interface() == v2.Interface()
}
