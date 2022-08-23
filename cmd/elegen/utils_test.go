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
	"reflect"
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

func Test_modelCommentFlags(t *testing.T) {
	type args struct {
		exts map[string]interface{}
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 []string
	}{
		{
			"nil",
			func(t *testing.T) args {
				return args{
					nil,
				}
			},
			nil,
		},
		{
			"no key",
			func(t *testing.T) args {
				return args{
					map[string]interface{}{"nope": nil},
				}
			},
			nil,
		},
		{
			"nil key",
			func(t *testing.T) args {
				return args{
					map[string]interface{}{"commentFlags": nil},
				}
			},
			nil,
		},
		{
			"empty key",
			func(t *testing.T) args {
				return args{
					map[string]interface{}{"commentFlags": []interface{}{}},
				}
			},
			nil,
		},
		{
			"key with //",
			func(t *testing.T) args {
				return args{
					map[string]interface{}{"commentFlags": []interface{}{"// hello world"}},
				}
			},
			[]string{"hello world"},
		},
		{
			"key without //",
			func(t *testing.T) args {
				return args{
					map[string]interface{}{"commentFlags": []interface{}{" hello world"}},
				}
			},
			[]string{"hello world"},
		},
		{
			"key with weird spacing",
			func(t *testing.T) args {
				return args{
					map[string]interface{}{"commentFlags": []interface{}{" //	 hello world"}},
				}
			},
			[]string{"hello world"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := modelCommentFlags(tArgs.exts)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("modelCommentFlags got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func Test_sortAttributes(t *testing.T) {
	type args struct {
		attrs []*spec.Attribute
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
		want []*spec.Attribute
	}{
		{
			"nil",
			func(t *testing.T) args {
				return args{
					nil,
				}
			},
			[]*spec.Attribute{},
		},
		{
			"one item",
			func(t *testing.T) args {
				return args{
					[]*spec.Attribute{
						{Name: "stuff"},
					},
				}
			},
			[]*spec.Attribute{
				{Name: "stuff"},
			},
		},
		{
			"two items",
			func(t *testing.T) args {
				return args{
					[]*spec.Attribute{
						{Name: "stuff"},
						{Name: "other"},
					},
				}
			},
			[]*spec.Attribute{
				{Name: "other"},
				{Name: "stuff"},
			},
		},
		{
			"three items",
			func(t *testing.T) args {
				return args{
					[]*spec.Attribute{
						{Name: "otherthings"},
						{Name: "stuff"},
						{Name: "moreitems"},
					},
				}
			},
			[]*spec.Attribute{
				{Name: "moreitems"},
				{Name: "otherthings"},
				{Name: "stuff"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortAttributes(tt.args(t).attrs)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortAttributes got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func Test_sortIndexes(t *testing.T) {
	type args struct {
		indexes [][]string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args
		want [][]string
	}{
		{
			"nil",
			func(t *testing.T) args {
				return args{
					nil,
				}
			},
			[][]string{},
		},
		{
			"one item",
			func(t *testing.T) args {
				return args{
					[][]string{
						{"stuff"},
					},
				}
			},
			[][]string{
				{"stuff"},
			},
		},
		{
			"two items",
			func(t *testing.T) args {
				return args{
					[][]string{
						{"stuff"},
						{"other"},
					},
				}
			},
			[][]string{
				{"other"},
				{"stuff"},
			},
		},
		{
			"three items",
			func(t *testing.T) args {
				return args{
					[][]string{
						{"stuff"},
						{"moreitems"},
						{"other"},
					},
				}
			},
			[][]string{
				{"moreitems"},
				{"other"},
				{"stuff"},
			},
		},
		{
			"four items, two long",
			func(t *testing.T) args {
				return args{
					[][]string{
						{"stuff", "two"},
						{"stuff", "one"},
						{"moreitems"},
						{"other"},
					},
				}
			},
			[][]string{
				{"moreitems"},
				{"other"},
				{"stuff", "one"},
				{"stuff", "two"},
			},
		},
		{
			"three items, diff item amounts",
			func(t *testing.T) args {
				return args{
					[][]string{
						{"stuff", "two", "other"},
						{"stuff", "two"},
						{"stuff", "three", "things", "items"},
					},
				}
			},
			[][]string{
				{"stuff", "three", "things", "items"},
				{"stuff", "two"},
				{"stuff", "two", "other"},
			},
		},
		{
			"two items, properly sorted",
			func(t *testing.T) args {
				return args{
					[][]string{
						{"stuff", "two"},
						{"stuff", "two", "other"},
					},
				}
			},
			[][]string{
				{"stuff", "two"},
				{"stuff", "two", "other"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortIndexes(tt.args(t).indexes)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortIndexes got: %v, want: %v", got, tt.want)
			}
		})
	}
}
