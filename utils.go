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
	"reflect"
	"time"
)

// RemoveZeroValues reset all pointer fields that are pointing to a zero value to nil
func RemoveZeroValues(obj interface{}) {

	vo := reflect.ValueOf(obj)
	vv := reflect.Indirect(vo)

	for _, field := range extractFieldNames(obj) {
		v := vv.FieldByName(field)

		if v.Kind() != reflect.Ptr {
			continue
		}

		if v.IsNil() {
			continue
		}

		uv := reflect.Indirect(v)

		switch uv.Kind() {
		case reflect.Map, reflect.Slice, reflect.Array:
			if uv.Len() == 0 {
				v.Set(reflect.Zero(v.Type()))
				break
			}

			fallthrough

		default:
			if uv.IsZero() {
				v.Set(reflect.Zero(v.Type()))
			}
		}
	}
}

// extractFieldNames returns all the field Name of the given
// object using reflection.
func extractFieldNames(obj interface{}) []string {

	val := reflect.Indirect(reflect.ValueOf(obj))
	c := val.NumField()
	fields := make([]string, c)

	for i := 0; i < c; i++ {
		fields[i] = val.Type().Field(i).Name
	}

	return fields
}

var reflectedTimeType = reflect.ValueOf(time.Time{}).Type()

// areFieldValuesEqual checks if the value of the given field name are
// equal in both given objects using reflection.
func areFieldValuesEqual(field string, o1, o2 interface{}) bool {

	field1 := reflect.Indirect(reflect.ValueOf(o1)).FieldByName(field)
	field2 := reflect.Indirect(reflect.ValueOf(o2)).FieldByName(field)

	if isFieldValueZero(field, o1) && isFieldValueZero(field, o2) {
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

// isFieldValueZero check if the value of the given field is set to its zero value.
func isFieldValueZero(field string, o interface{}) bool {

	return IsZero(reflect.Indirect(reflect.ValueOf(o)).FieldByName(field).Interface())
}

// IsZero returns true if the given value is set to its Zero value.
func IsZero(o interface{}) bool {

	if o == nil {
		return true
	}

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

func areFieldsValueEqualValue(f string, obj interface{}, value interface{}) bool {

	field := reflect.Indirect(reflect.ValueOf(obj)).FieldByName(f)

	if value == nil {
		return isFieldValueZero(f, obj)
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
