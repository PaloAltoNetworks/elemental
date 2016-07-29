// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import "reflect"

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

// fieldValuesEquals check if the value of the given field name are
// equal in both given objects using reflection.
func fieldValuesEquals(field string, o1, o2 interface{}) bool {

	return reflect.ValueOf(o1).Elem().FieldByName(field).Interface() == reflect.ValueOf(o2).Elem().FieldByName(field).Interface()
}
