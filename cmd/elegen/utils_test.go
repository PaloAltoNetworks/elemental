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

package main

import (
	"testing"

	"go.aporeto.io/regolithe/spec"
)

func Test_attributeTypeConverter(t *testing.T) {
	type args struct {
		typ     spec.AttributeType
		subtype string
	}
	tests := []struct {
		name           string
		args           args
		expectedType   string
		expectedImport string
	}{
		{
			name: "attribute type is a string",
			args: args{
				typ:     spec.AttributeTypeString,
				subtype: "",
			},
			expectedType:   "string",
			expectedImport: "",
		},
		{
			name: "attribute type is a float",
			args: args{
				typ:     spec.AttributeTypeFloat,
				subtype: "",
			},
			expectedType:   "float64",
			expectedImport: "",
		},
		{
			name: "attribute type is a bool",
			args: args{
				typ:     spec.AttributeTypeBool,
				subtype: "",
			},
			expectedType:   "bool",
			expectedImport: "",
		},
		{
			name: "attribute type is an int",
			args: args{
				typ:     spec.AttributeTypeInt,
				subtype: "",
			},
			expectedType:   "int",
			expectedImport: "",
		},
		{
			name: "attribute type is a time",
			args: args{
				typ:     spec.AttributeTypeTime,
				subtype: "",
			},
			expectedType:   "time.Time",
			expectedImport: "time",
		},
		{
			name: "attribute type is a list",
			args: args{
				typ:     spec.AttributeTypeList,
				subtype: "",
			},
			expectedType:   "[]interface{}",
			expectedImport: "",
		},
		{
			name: "attribute type is a list of objects",
			args: args{
				typ:     spec.AttributeTypeList,
				subtype: "object",
			},
			expectedType:   "[]interface{}",
			expectedImport: "",
		},
		{
			name: "attribute type is a list of int",
			args: args{
				typ:     spec.AttributeTypeList,
				subtype: "integer",
			},
			expectedType:   "[]int",
			expectedImport: "",
		},
		{
			name: "attribute type is a list of bool",
			args: args{
				typ:     spec.AttributeTypeList,
				subtype: "boolean",
			},
			expectedType:   "[]bool",
			expectedImport: "",
		},
		{
			name: "attribute type is a list of float",
			args: args{
				typ:     spec.AttributeTypeList,
				subtype: "float",
			},
			expectedType:   "[]float64",
			expectedImport: "",
		},
		{
			name: "attribute type is a list of time",
			args: args{
				typ:     spec.AttributeTypeList,
				subtype: "time",
			},
			expectedType:   "[]time.Time",
			expectedImport: "time",
		},
		{
			name: "attribute type is a list of string",
			args: args{
				typ:     spec.AttributeTypeList,
				subtype: "string",
			},
			expectedType:   "[]string",
			expectedImport: "",
		},
		{
			name: "attribute type is a random thing",
			args: args{
				typ:     "SomeThing",
				subtype: "",
			},
			expectedType:   "interface{}",
			expectedImport: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := attributeTypeConverter(tt.args.typ, tt.args.subtype)
			if got != tt.expectedType {
				t.Errorf("attributeTypeConverter() got = %v, expectedType %v", got, tt.expectedType)
			}
			if got1 != tt.expectedImport {
				t.Errorf("attributeTypeConverter() got1 = %v, expectedType %v", got1, tt.expectedImport)
			}
		})
	}
}

func Test_attrBSONFieldName(t *testing.T) {

	testCases := map[string]struct {
		attr        *spec.Attribute
		expected    string
		shouldPanic bool
	}{
		"basic - no override": {
			attr: &spec.Attribute{
				Name:       "SomeAttribute",
				Stored:     true,
				Identifier: false,
				Extensions: nil,
			},
			expected:    "someattribute",
			shouldPanic: false,
		},
		"basic - with override": {
			attr: &spec.Attribute{
				Name:       "SomeAttribute",
				Stored:     true,
				Identifier: false,
				Extensions: map[string]interface{}{
					"bson_name": "sa",
				},
			},
			expected:    "sa",
			shouldPanic: false,
		},
		"identifier": {
			attr: &spec.Attribute{
				Name:       "SomeAttribute",
				Stored:     true,
				Identifier: true,
				Extensions: nil,
			},
			expected:    "_id",
			shouldPanic: false,
		},
		"should panic if attribute is not stored": {
			attr: &spec.Attribute{
				Name:       "SomeAttribute",
				Stored:     false,
				Identifier: true,
				Extensions: nil,
			},
			expected:    "",
			shouldPanic: true,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && !tc.shouldPanic {
					t.Errorf("did not expect a panic, but one occurred: %s", err)
				}
			}()
			if actual := attrBSONFieldName(tc.attr); actual != tc.expected {
				t.Errorf("expected: '%s'\n"+
					"actual: '%s'\n",
					tc.expected,
					actual)
			}
		})
	}
}
