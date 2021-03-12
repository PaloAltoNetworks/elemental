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
	"testing"
	"time"

	"github.com/araddon/dateparse"
	. "github.com/smartystreets/goconvey/convey"
)

func TestParameter_Validate(t *testing.T) {
	type fields struct {
		Name           string
		Type           ParameterType
		AllowedChoices []string
		DefaultValue   string
		Multiple       bool
	}
	type args struct {
		values []string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantParam *Parameter
		wantErr   bool
	}{
		{
			"multiple while forbidden",
			fields{
				Name:     "param",
				Multiple: false,
			},
			args{
				[]string{"a", "b"},
			},
			nil,
			true,
		},
		{
			"valid string",
			fields{
				Name: "param",
				Type: ParameterTypeString,
			},
			args{
				[]string{"string"},
			},
			&Parameter{
				ptype:  ParameterTypeString,
				values: []interface{}{"string"},
			},
			false,
		},
		{
			"valid int",
			fields{
				Name: "param",
				Type: ParameterTypeInt,
			},
			args{
				[]string{"1"},
			},
			&Parameter{
				ptype:  ParameterTypeInt,
				values: []interface{}{1},
			},
			false,
		},
		{
			"invalid int",
			fields{
				Name: "param",
				Type: ParameterTypeInt,
			},
			args{
				[]string{"not1"},
			},
			nil,
			true,
		},
		{
			"valid bools",
			fields{
				Name:     "param",
				Type:     ParameterTypeBool,
				Multiple: true,
			},
			args{
				[]string{"TRUE", "FALSE", "YES", "NO", "1", "0"},
			},
			&Parameter{
				ptype:  ParameterTypeBool,
				values: []interface{}{true, false, true, false, true, false},
			},
			false,
		},
		{
			"invalid bool",
			fields{
				Name: "param",
				Type: ParameterTypeBool,
			},
			args{
				[]string{"NOTTRUE"},
			},
			nil,
			true,
		},
		{
			"valid float",
			fields{
				Name: "param",
				Type: ParameterTypeFloat,
			},
			args{
				[]string{"1.004"},
			},
			&Parameter{
				ptype:  ParameterTypeFloat,
				values: []interface{}{1.004},
			},
			false,
		},
		{
			"valid float",
			fields{
				Name: "param",
				Type: ParameterTypeFloat,
			},
			args{
				[]string{"1"},
			},
			&Parameter{
				ptype:  ParameterTypeFloat,
				values: []interface{}{1.0},
			},
			false,
		},
		{
			"invalid float",
			fields{
				Name: "param",
				Type: ParameterTypeFloat,
			},
			args{
				[]string{"not1.0"},
			},
			nil,
			true,
		},
		{
			"valid enum",
			fields{
				Name:           "param",
				Type:           ParameterTypeEnum,
				AllowedChoices: []string{"A", "B"},
			},
			args{
				[]string{"A"},
			},
			&Parameter{
				ptype:  ParameterTypeEnum,
				values: []interface{}{"A"},
			},
			false,
		},
		{
			"invalid enum",
			fields{
				Name:           "param",
				Type:           ParameterTypeEnum,
				AllowedChoices: []string{"A", "B"},
			},
			args{
				[]string{"C"},
			},
			nil,
			true,
		},
		{
			"valid duration",
			fields{
				Name: "param",
				Type: ParameterTypeDuration,
			},
			args{
				[]string{"3s"},
			},
			&Parameter{
				ptype:  ParameterTypeDuration,
				values: []interface{}{3 * time.Second},
			},
			false,
		},
		{
			"invalid duration",
			fields{
				Name: "param",
				Type: ParameterTypeDuration,
			},
			args{
				[]string{"3apples"},
			},
			nil,
			true,
		},
		{
			"valid date",
			fields{
				Name:     "param",
				Type:     ParameterTypeTime,
				Multiple: true,
			},
			args{
				[]string{"oct 7, 1970", "04/08/2014 22:05", "1384216367189", "1384216367111222333", "03/19/2012 10:11:59.3186369"},
			},
			&Parameter{
				ptype: ParameterTypeTime,
				values: []interface{}{
					dateparse.MustParse("oct 7, 1970"),
					dateparse.MustParse("04/08/2014 22:05"),
					dateparse.MustParse("1384216367189"),
					dateparse.MustParse("1384216367111222333"),
					dateparse.MustParse("03/19/2012 10:11:59.3186369"),
				},
			},
			false,
		},
		{
			"invalid date",
			fields{
				Name: "param",
				Type: ParameterTypeTime,
			},
			args{
				[]string{"not date"},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ParameterDefinition{
				Name:           tt.fields.Name,
				Type:           tt.fields.Type,
				AllowedChoices: tt.fields.AllowedChoices,
				DefaultValue:   tt.fields.DefaultValue,
				Multiple:       tt.fields.Multiple,
			}
			if p, err := p.Parse(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("Parameter.Parse() error = %v, wantErr %v", err, tt.wantErr)
			} else if !reflect.DeepEqual(p, tt.wantParam) {
				t.Errorf("Parameter.Parse() param = %v, wantParam %v", p, tt.wantParam)
			}
		})
	}
}

func TestParameters_Value(t *testing.T) {

	Convey("Given I call Parse with unknown type", t, func() {

		p := &ParameterDefinition{
			Type: ParameterType("yo"),
		}

		Convey("Then it should panic", func() {
			So(func() { _, _ = p.Parse([]string{"a"}) }, ShouldPanicWith, `unknown parameter type: 'yo'`)
		})
	})

	Convey("Given I have a 2 string parameter", t, func() {

		p := &ParameterDefinition{
			Type:     ParameterTypeString,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			pp, err := p.Parse([]string{"a", "b"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(pp.StringValue(), ShouldResemble, "a")
			})

			Convey("Then the multiple values should be accessible", func() {
				So(pp.StringValues(), ShouldResemble, []string{"a", "b"})
			})

			Convey("Then the all values should be accessible", func() {
				So(pp.Values(), ShouldResemble, []interface{}{"a", "b"})
			})
		})
	})

	Convey("Given I have a 2 enum parameter", t, func() {

		p := &ParameterDefinition{
			Type:           ParameterTypeEnum,
			AllowedChoices: []string{"A", "B"},
			Multiple:       true,
		}

		Convey("When I parse it", func() {

			pp, err := p.Parse([]string{"A", "B"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(pp.StringValue(), ShouldResemble, "A")
			})

			Convey("Then the multiple values should be accessible", func() {
				So(pp.StringValues(), ShouldResemble, []string{"A", "B"})
			})

			Convey("Then the all values should be accessible", func() {
				So(pp.Values(), ShouldResemble, []interface{}{"A", "B"})
			})
		})
	})

	Convey("Given I have a 2 int parameter", t, func() {

		p := &ParameterDefinition{
			Type:     ParameterTypeInt,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			pp, err := p.Parse([]string{"1", "2"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(pp.IntValue(), ShouldResemble, 1)
			})

			Convey("Then the multiple values should be accessible", func() {
				So(pp.IntValues(), ShouldResemble, []int{1, 2})
			})

			Convey("Then the all values should be accessible", func() {
				So(pp.Values(), ShouldResemble, []interface{}{1, 2})
			})
		})
	})

	Convey("Given I have a 2 float parameter", t, func() {

		p := &ParameterDefinition{
			Type:     ParameterTypeFloat,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			pp, err := p.Parse([]string{"1.1", "2.2"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(pp.FloatValue(), ShouldResemble, 1.1)
			})

			Convey("Then the multiple values should be accessible", func() {
				So(pp.FloatValues(), ShouldResemble, []float64{1.1, 2.2})
			})

			Convey("Then the all values should be accessible", func() {
				So(pp.Values(), ShouldResemble, []interface{}{1.1, 2.2})
			})
		})
	})

	Convey("Given I have a 2 bool parameter", t, func() {

		p := &ParameterDefinition{
			Type:     ParameterTypeBool,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			pp, err := p.Parse([]string{"true", "false", "yes"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(pp.BoolValue(), ShouldResemble, true)
			})

			Convey("Then the multiple values should be accessible", func() {
				So(pp.BoolValues(), ShouldResemble, []bool{true, false, true})
			})

			Convey("Then the all values should be accessible", func() {
				So(pp.Values(), ShouldResemble, []interface{}{true, false, true})
			})
		})
	})

	Convey("Given I have a 2 duration parameter", t, func() {

		p := &ParameterDefinition{
			Type:     ParameterTypeDuration,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			pp, err := p.Parse([]string{"-2s", "2h"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(pp.DurationValue(), ShouldResemble, -2*time.Second)
			})

			Convey("Then the multiple values should be accessible", func() {
				So(pp.DurationValues(), ShouldResemble, []time.Duration{-2 * time.Second, 2 * time.Hour})
			})

			Convey("Then the all values should be accessible", func() {
				So(pp.Values(), ShouldResemble, []interface{}{-2 * time.Second, 2 * time.Hour})
			})
		})
	})

	Convey("Given I have a 2 time parameter", t, func() {

		t1 := time.Now()
		t2 := time.Now().Add(-3 * time.Hour)

		p := &ParameterDefinition{
			Type:     ParameterTypeTime,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			pp, err := p.Parse([]string{t1.Format(time.RFC3339), t2.Format(time.RFC3339)})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(pp.TimeValue().Format(time.RFC3339), ShouldResemble, t1.Format(time.RFC3339))
			})

			Convey("Then the multiple values should be accessible", func() {
				So(len(pp.TimeValues()), ShouldEqual, 2)
				So(pp.TimeValues()[0].Format(time.RFC3339), ShouldResemble, t1.Format(time.RFC3339))
				So(pp.TimeValues()[1].Format(time.RFC3339), ShouldResemble, t2.Format(time.RFC3339))
			})

			Convey("Then the all values should be accessible", func() {
				So(len(pp.Values()), ShouldEqual, 2)
				So(pp.Values()[0].(time.Time).Format(time.RFC3339), ShouldResemble, t1.Format(time.RFC3339))
				So(pp.Values()[1].(time.Time).Format(time.RFC3339), ShouldResemble, t2.Format(time.RFC3339))
			})
		})
	})
}

func TestParameters_NewParameter(t *testing.T) {

	Convey("Given I call NewParameter", t, func() {

		p := NewParameter(ParameterTypeString, "a", "b")

		Convey("Then the parameter should be correct", func() {
			So(p.StringValue(), ShouldEqual, "a")
			So(p.Values(), ShouldResemble, []interface{}{"a", "b"})
		})
	})
}

func TestParameters_Requirements(t *testing.T) {

	Convey("Given I have an empty parameter requirement", t, func() {

		req := NewParametersRequirement([][][]string{})

		Convey("When I call Validate params a and b and 1", func() {

			params := Parameters{
				"a": NewParameter(ParameterTypeString, "a"),
				"b": NewParameter(ParameterTypeString, "b"),
				"1": NewParameter(ParameterTypeString, "1"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate no params", func() {

			params := Parameters{}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a simple parameter requirement", t, func() {

		req := NewParametersRequirement([][][]string{
			{
				{"a"},
			},
		})

		Convey("When I call Validate params a", func() {

			params := Parameters{
				"a": NewParameter(ParameterTypeString, "a"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate no params", func() {

			params := Parameters{}

			err := params.Validate(req)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a single parameter requirement", t, func() {

		req := NewParametersRequirement(
			[][][]string{
				{
					{"a", "b"},
					{"c", "d"},
				},
			},
		)

		Convey("When I call Validate params a and b and 1", func() {

			params := Parameters{
				"a": NewParameter(ParameterTypeString, "a"),
				"b": NewParameter(ParameterTypeString, "b"),
				"1": NewParameter(ParameterTypeString, "1"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate params c and d and 1 ", func() {

			params := Parameters{
				"c": NewParameter(ParameterTypeString, "c"),
				"d": NewParameter(ParameterTypeString, "d"),
				"1": NewParameter(ParameterTypeString, "1"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate params a and 1 ", func() {

			params := Parameters{
				"a": NewParameter(ParameterTypeString, "a"),
				"1": NewParameter(ParameterTypeString, "1"),
			}

			err := params.Validate(req)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I call Validate params b and c ", func() {

			params := Parameters{
				"b": NewParameter(ParameterTypeString, "a"),
				"c": NewParameter(ParameterTypeString, "c"),
			}

			err := params.Validate(req)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a multiple parameter requirement", t, func() {

		req := NewParametersRequirement(
			[][][]string{
				{
					{"a", "b"},
					{"c", "d"},
				},
				{
					{"1", "2"},
				},
			})

		Convey("When I call Validate params a and b and 1 and 2", func() {

			params := Parameters{
				"a": NewParameter(ParameterTypeString, "a"),
				"b": NewParameter(ParameterTypeString, "b"),
				"1": NewParameter(ParameterTypeString, "1"),
				"2": NewParameter(ParameterTypeString, "2"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate params c and d and 1 and 2", func() {

			params := Parameters{
				"c": NewParameter(ParameterTypeString, "c"),
				"d": NewParameter(ParameterTypeString, "d"),
				"1": NewParameter(ParameterTypeString, "1"),
				"2": NewParameter(ParameterTypeString, "2"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate params a and d and 1 and 2", func() {

			params := Parameters{
				"a": NewParameter(ParameterTypeString, "a"),
				"d": NewParameter(ParameterTypeString, "d"),
				"1": NewParameter(ParameterTypeString, "1"),
				"2": NewParameter(ParameterTypeString, "2"),
			}

			err := params.Validate(req)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I call Validate params a and b and 1", func() {

			params := Parameters{
				"a": NewParameter(ParameterTypeString, "a"),
				"b": NewParameter(ParameterTypeString, "b"),
				"1": NewParameter(ParameterTypeString, "1"),
			}

			err := params.Validate(req)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When I call Validate params a and 1 and 2", func() {

			params := Parameters{
				"a": NewParameter(ParameterTypeString, "a"),
				"1": NewParameter(ParameterTypeString, "1"),
				"2": NewParameter(ParameterTypeString, "2"),
			}

			err := params.Validate(req)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a complex multiple parameter requirement", t, func() {

		req := NewParametersRequirement(
			[][][]string{
				{
					{"er"},
					{"sr"},
					{"sr", "er"},
					{"sr", "ea"},
					{"sa", "er"},
					{"sa", "ea"},
				},
			})

		Convey("When I call Validate params er", func() {

			params := Parameters{
				"er": NewParameter(ParameterTypeString, "a"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate params sr", func() {

			params := Parameters{
				"sr": NewParameter(ParameterTypeString, "a"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate params sr and er", func() {

			params := Parameters{
				"sr": NewParameter(ParameterTypeString, "a"),
				"er": NewParameter(ParameterTypeString, "a"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate params sr and ea", func() {

			params := Parameters{
				"sr": NewParameter(ParameterTypeString, "a"),
				"ea": NewParameter(ParameterTypeString, "a"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate params sa and er", func() {

			params := Parameters{
				"sa": NewParameter(ParameterTypeString, "a"),
				"er": NewParameter(ParameterTypeString, "a"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When I call Validate params sa and ea", func() {

			params := Parameters{
				"sa": NewParameter(ParameterTypeString, "a"),
				"ea": NewParameter(ParameterTypeString, "a"),
			}

			err := params.Validate(req)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestParameter_DefaultValue(t *testing.T) {

	Convey("Given I have a string parameter with default value", t, func() {

		dv := "hello"
		p := NewParameter(ParameterTypeString)
		p.defaultValue = dv

		Convey("When I call the accessor", func() {

			v := p.StringValue()

			Convey("Then it should return the default value", func() {
				So(v, ShouldEqual, dv)
			})
		})

		Convey("When I call the wrong accessor", func() {

			v := p.IntValue()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldEqual, 0)
			})
		})

		Convey("When I call the wrong multiple accessor", func() {

			v := p.IntValues()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a int parameter with default value", t, func() {

		dv := 42
		p := NewParameter(ParameterTypeInt)
		p.defaultValue = dv

		Convey("When I call the accessor", func() {

			v := p.IntValue()

			Convey("Then it should return the default value", func() {
				So(v, ShouldEqual, dv)
			})
		})

		Convey("When I call the wrong accessor", func() {

			v := p.FloatValue()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldEqual, 0.0)
			})
		})

		Convey("When I call the wrong multiple accessor", func() {

			v := p.FloatValues()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a float parameter with default value", t, func() {

		dv := 42.42
		p := NewParameter(ParameterTypeFloat)
		p.defaultValue = dv

		Convey("When I call the accessor", func() {

			v := p.FloatValue()

			Convey("Then it should return the default value", func() {
				So(v, ShouldEqual, dv)
			})
		})

		Convey("When I call the wrong accessor", func() {

			v := p.BoolValue()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldEqual, false)
			})
		})

		Convey("When I call the wrong multiple accessor", func() {

			v := p.BoolValues()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a bool parameter with default value", t, func() {

		dv := true
		p := NewParameter(ParameterTypeBool)
		p.defaultValue = dv

		Convey("When I call the accessor", func() {

			v := p.BoolValue()

			Convey("Then it should return the default value", func() {
				So(v, ShouldEqual, dv)
			})
		})

		Convey("When I call the wrong accessor", func() {

			v := p.DurationValue()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldEqual, 0*time.Second)
			})
		})

		Convey("When I call the wrong multiple accessor", func() {

			v := p.DurationValues()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a duration parameter with default value", t, func() {

		dv := 3 * time.Second
		p := NewParameter(ParameterTypeDuration)
		p.defaultValue = dv

		Convey("When I call the accessor", func() {

			v := p.DurationValue()

			Convey("Then it should return the default value", func() {
				So(v, ShouldEqual, dv)
			})
		})

		Convey("When I call the wrong accessor", func() {

			v := p.TimeValue()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldResemble, time.Time{})
			})
		})

		Convey("When I call the wrong multiple accessor", func() {

			v := p.TimeValues()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a time parameter with default value", t, func() {

		dv := time.Now()
		p := NewParameter(ParameterTypeTime)
		p.defaultValue = dv

		Convey("When I call the accessor", func() {

			v := p.TimeValue()

			Convey("Then it should return the default value", func() {
				So(v, ShouldEqual, dv)
			})
		})

		Convey("When I call the wrong accessor", func() {

			v := p.StringValue()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldEqual, "")
			})
		})

		Convey("When I call the wrong multiple accessor", func() {

			v := p.StringValues()

			Convey("Then it should return the zero value", func() {
				So(v, ShouldBeNil)
			})
		})
	})
}

func TestParameterRequirement_String(t *testing.T) {

	Convey("Given I have some requirements", t, func() {

		req := NewParametersRequirement(
			[][][]string{
				{
					{"a", "b"},
					{"c", "d"},
				},
				{
					{"1", "2"},
				},
			},
		)

		Convey("When I call String", func() {

			out := req.String()

			Convey("Then it should be correct", func() {
				So(out, ShouldEqual, `((a and b) or (c and d)) and (1 and 2)`)
			})
		})
	})

	Convey("Given I have some simple requirements", t, func() {

		req := NewParametersRequirement(
			[][][]string{
				{
					{"a"},
				},
			},
		)

		Convey("When I call String", func() {

			out := req.String()

			Convey("Then it should be correct", func() {
				So(out, ShouldEqual, `a`)
			})
		})
	})
}

func TestParameters_Get(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		p    Parameters
		args args
		want Parameter
	}{
		{
			"get non existig param",
			Parameters{},
			args{
				"key",
			},
			Parameter{},
		},
		{
			"get existig param",
			Parameters{
				"key": NewParameter(ParameterTypeString, "a"),
			},
			args{
				"key",
			},
			NewParameter(ParameterTypeString, "a"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Get(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameters.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {

	type args struct {
		ptype          ParameterType
		pname          string
		v              string
		allowedChoices []string
	}
	tests := []struct {
		name    string
		args    args
		wantOut interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := parse(tt.args.ptype, tt.args.pname, tt.args.v, tt.args.allowedChoices)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("parse() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestParameterDefinition_Parse(t *testing.T) {
	type fields struct {
		Name           string
		Type           ParameterType
		AllowedChoices []string
		DefaultValue   string
		Multiple       bool
	}
	type args struct {
		values []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Parameter
		wantErr bool
	}{
		{
			"test multiple values while not allowed",
			fields{
				Name:     "p",
				Type:     ParameterTypeString,
				Multiple: false,
			},
			args{
				[]string{"a", "b"},
			},
			nil,
			true,
		},

		// strings
		{
			"test set type string",
			fields{
				Name:         "p",
				Type:         ParameterTypeString,
				DefaultValue: "default",
			},
			args{
				[]string{"a"},
			},
			&Parameter{
				defaultValue: nil,
				values:       []interface{}{"a"},
				ptype:        ParameterTypeString,
			},
			false,
		},
		{
			"test not set type string",
			fields{
				Name:         "p",
				Type:         ParameterTypeString,
				DefaultValue: "default",
			},
			args{
				[]string{},
			},
			&Parameter{
				defaultValue: "default",
				values:       nil,
				ptype:        ParameterTypeString,
			},
			false,
		},

		// int
		{
			"test set type int",
			fields{
				Name:         "p",
				Type:         ParameterTypeInt,
				DefaultValue: "42",
			},
			args{
				[]string{"43"},
			},
			&Parameter{
				defaultValue: nil,
				values:       []interface{}{43},
				ptype:        ParameterTypeInt,
			},
			false,
		},
		{
			"test not set type int",
			fields{
				Name:         "p",
				Type:         ParameterTypeInt,
				DefaultValue: "42",
			},
			args{
				[]string{},
			},
			&Parameter{
				defaultValue: 42,
				values:       nil,
				ptype:        ParameterTypeInt,
			},
			false,
		},
		{
			"test bad set type int",
			fields{
				Name:         "p",
				Type:         ParameterTypeInt,
				DefaultValue: "42",
			},
			args{
				[]string{"not int"},
			},
			nil,
			true,
		},
		{
			"test bad set default type int",
			fields{
				Name:         "p",
				Type:         ParameterTypeInt,
				DefaultValue: "not int",
			},
			args{},
			nil,
			true,
		},

		// float
		{
			"test set type float",
			fields{
				Name:         "p",
				Type:         ParameterTypeFloat,
				DefaultValue: "42.42",
			},
			args{
				[]string{"43.43"},
			},
			&Parameter{
				defaultValue: nil,
				values:       []interface{}{43.43},
				ptype:        ParameterTypeFloat,
			},
			false,
		},
		{
			"test not set type float",
			fields{
				Name:         "p",
				Type:         ParameterTypeFloat,
				DefaultValue: "42.43",
			},
			args{
				[]string{},
			},
			&Parameter{
				defaultValue: 42.43,
				values:       nil,
				ptype:        ParameterTypeFloat,
			},
			false,
		},
		{
			"test bad set type float",
			fields{
				Name:         "p",
				Type:         ParameterTypeFloat,
				DefaultValue: "42",
			},
			args{
				[]string{"not float"},
			},
			nil,
			true,
		},
		{
			"test bad set default type float",
			fields{
				Name:         "p",
				Type:         ParameterTypeFloat,
				DefaultValue: "not float",
			},
			args{},
			nil,
			true,
		},

		// bool
		{
			"test set type bool true",
			fields{
				Name:         "p",
				Type:         ParameterTypeBool,
				DefaultValue: "false",
			},
			args{
				[]string{"true"},
			},
			&Parameter{
				defaultValue: nil,
				values:       []interface{}{true},
				ptype:        ParameterTypeBool,
			},
			false,
		},
		{
			"test set type bool false",
			fields{
				Name:         "p",
				Type:         ParameterTypeBool,
				DefaultValue: "true",
			},
			args{
				[]string{"false"},
			},
			&Parameter{
				defaultValue: nil,
				values:       []interface{}{false},
				ptype:        ParameterTypeBool,
			},
			false,
		},
		{
			"test not set type bool",
			fields{
				Name:         "p",
				Type:         ParameterTypeBool,
				DefaultValue: "true",
			},
			args{
				[]string{},
			},
			&Parameter{
				defaultValue: true,
				values:       nil,
				ptype:        ParameterTypeBool,
			},
			false,
		},
		{
			"test bad set type bool",
			fields{
				Name:         "p",
				Type:         ParameterTypeBool,
				DefaultValue: "true",
			},
			args{
				[]string{"not bool"},
			},
			nil,
			true,
		},
		{
			"test bad set default type bool",
			fields{
				Name:         "p",
				Type:         ParameterTypeBool,
				DefaultValue: "not bool",
			},
			args{},
			nil,
			true,
		},

		// enum
		{
			"test set type enum",
			fields{
				Name:           "p",
				Type:           ParameterTypeEnum,
				DefaultValue:   "hello",
				AllowedChoices: []string{"hello", "bye"},
			},
			args{
				[]string{"bye"},
			},
			&Parameter{
				defaultValue: nil,
				values:       []interface{}{"bye"},
				ptype:        ParameterTypeEnum,
			},
			false,
		},
		{
			"test not set type enum",
			fields{
				Name:           "p",
				Type:           ParameterTypeEnum,
				DefaultValue:   "bye",
				AllowedChoices: []string{"hello", "bye"},
			},
			args{
				[]string{},
			},
			&Parameter{
				defaultValue: "bye",
				values:       nil,
				ptype:        ParameterTypeEnum,
			},
			false,
		},
		{
			"test bad set type enum",
			fields{
				Name:           "p",
				Type:           ParameterTypeEnum,
				DefaultValue:   "bye",
				AllowedChoices: []string{"hello", "bye"},
			},
			args{
				[]string{"not good"},
			},
			nil,
			true,
		},
		{
			"test bad set default type enum",
			fields{
				Name:           "p",
				Type:           ParameterTypeEnum,
				DefaultValue:   "not good",
				AllowedChoices: []string{"hello", "bye"},
			},
			args{},
			nil,
			true,
		},

		// duration
		{
			"test set type duration",
			fields{
				Name:         "p",
				Type:         ParameterTypeDuration,
				DefaultValue: "42s",
			},
			args{
				[]string{"43s"},
			},
			&Parameter{
				defaultValue: nil,
				values:       []interface{}{43 * time.Second},
				ptype:        ParameterTypeDuration,
			},
			false,
		},
		{
			"test not set type duration",
			fields{
				Name:         "p",
				Type:         ParameterTypeDuration,
				DefaultValue: "42s",
			},
			args{
				[]string{},
			},
			&Parameter{
				defaultValue: 42 * time.Second,
				values:       nil,
				ptype:        ParameterTypeDuration,
			},
			false,
		},
		{
			"test bad set type duration",
			fields{
				Name:         "p",
				Type:         ParameterTypeDuration,
				DefaultValue: "42s",
			},
			args{
				[]string{"not duration"},
			},
			nil,
			true,
		},
		{
			"test bad set default type duration",
			fields{
				Name:         "p",
				Type:         ParameterTypeFloat,
				DefaultValue: "not duration",
			},
			args{},
			nil,
			true,
		},

		// time
		{
			"test set type time",
			fields{
				Name:         "p",
				Type:         ParameterTypeTime,
				DefaultValue: "oct 7, 1970",
			},
			args{
				[]string{"May 8, 2009 5:57:51 PM"},
			},
			&Parameter{
				defaultValue: nil,
				values:       []interface{}{dateparse.MustParse("May 8, 2009 5:57:51 PM")},
				ptype:        ParameterTypeTime,
			},
			false,
		},
		{
			"test not set type time",
			fields{
				Name:         "p",
				Type:         ParameterTypeTime,
				DefaultValue: "oct 7, 1970",
			},
			args{
				[]string{},
			},
			&Parameter{
				defaultValue: dateparse.MustParse("oct 7, 1970"),
				values:       nil,
				ptype:        ParameterTypeTime,
			},
			false,
		},
		{
			"test bad set type time",
			fields{
				Name:         "p",
				Type:         ParameterTypeTime,
				DefaultValue: "oct 7, 1970",
			},
			args{
				[]string{"not time"},
			},
			nil,
			true,
		},
		{
			"test bad set default type time",
			fields{
				Name:         "p",
				Type:         ParameterTypeTime,
				DefaultValue: "not time",
			},
			args{},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ParameterDefinition{
				Name:           tt.fields.Name,
				Type:           tt.fields.Type,
				AllowedChoices: tt.fields.AllowedChoices,
				DefaultValue:   tt.fields.DefaultValue,
				Multiple:       tt.fields.Multiple,
			}
			got, err := p.Parse(tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParameterDefinition.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParameterDefinition.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
