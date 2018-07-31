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
		DefaultValue   interface{}
		Required       bool
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
			"missing required",
			fields{
				Name:     "param",
				Required: true,
			},
			args{
				[]string{},
			},
			nil,
			true,
		},
		{
			"missing required with nil",
			fields{
				Name:     "param",
				Required: true,
			},
			args{
				nil,
			},
			nil,
			true,
		},
		{
			"missing required with empty",
			fields{
				Name:     "param",
				Required: true,
				Type:     ParameterTypeBool,
			},
			args{
				[]string{""},
			},
			nil,
			true,
		},
		{
			"valid string",
			fields{
				Name:     "param",
				Required: true,
				Type:     ParameterTypeString,
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
				Name:     "param",
				Required: true,
				Type:     ParameterTypeInt,
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
				Name:     "param",
				Required: true,
				Type:     ParameterTypeInt,
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
				Required: true,
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
				Name:     "param",
				Required: true,
				Type:     ParameterTypeBool,
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
				Name:     "param",
				Required: true,
				Type:     ParameterTypeFloat,
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
				Name:     "param",
				Required: true,
				Type:     ParameterTypeFloat,
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
				Name:     "param",
				Required: true,
				Type:     ParameterTypeFloat,
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
				Required:       true,
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
				Required:       true,
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
				Name:     "param",
				Required: true,
				Type:     ParameterTypeDuration,
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
				Name:     "param",
				Required: true,
				Type:     ParameterTypeDuration,
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
				Required: true,
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
				Name:     "param",
				Required: true,
				Type:     ParameterTypeTime,
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
				Required:       tt.fields.Required,
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

			pp, err := p.Parse([]string{"2s", "2h"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(pp.DurationValue(), ShouldResemble, 2*time.Second)
			})
			Convey("Then the all values should be accessible", func() {
				So(pp.Values(), ShouldResemble, []interface{}{2 * time.Second, 2 * time.Hour})
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

	Convey("Given I have a single parameter requirement", t, func() {

		req := NewParametersRequirement([][][]string{
			[][]string{
				[]string{"a", "b"},
				[]string{"c", "d"},
			},
		})

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

		req := NewParametersRequirement([][][]string{
			[][]string{
				[]string{"a", "b"},
				[]string{"c", "d"},
			},
			[][]string{
				[]string{"1", "2"},
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
}
