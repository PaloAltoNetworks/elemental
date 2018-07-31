package elemental

import (
	"testing"
	"time"

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
		name    string
		fields  fields
		args    args
		wantErr bool
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
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Name:           tt.fields.Name,
				Type:           tt.fields.Type,
				AllowedChoices: tt.fields.AllowedChoices,
				DefaultValue:   tt.fields.DefaultValue,
				Required:       tt.fields.Required,
				Multiple:       tt.fields.Multiple,
			}
			if err := p.Parse(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("Parameter.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameters_Value(t *testing.T) {

	Convey("Given I call Parse with unknown type", t, func() {

		p := Parameter{
			Type: ParameterType("yo"),
		}

		Convey("Then it should panic", func() {
			So(func() { _ = p.Parse([]string{"a"}) }, ShouldPanicWith, `unknown parameter type: 'yo'`)
		})
	})

	Convey("Given I have a 2 string parameter", t, func() {

		p := Parameter{
			Type:     ParameterTypeString,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			err := p.Parse([]string{"a", "b"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(p.StringValue(), ShouldResemble, "a")
			})
			Convey("Then the all values should be accessible", func() {
				So(p.Values(), ShouldResemble, []interface{}{"a", "b"})
			})
		})
	})

	Convey("Given I have a 2 enum parameter", t, func() {

		p := Parameter{
			Type:           ParameterTypeEnum,
			AllowedChoices: []string{"A", "B"},
			Multiple:       true,
		}

		Convey("When I parse it", func() {

			err := p.Parse([]string{"A", "B"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(p.StringValue(), ShouldResemble, "A")
			})
			Convey("Then the all values should be accessible", func() {
				So(p.Values(), ShouldResemble, []interface{}{"A", "B"})
			})
		})
	})

	Convey("Given I have a 2 int parameter", t, func() {

		p := Parameter{
			Type:     ParameterTypeInt,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			err := p.Parse([]string{"1", "2"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(p.IntValue(), ShouldResemble, 1)
			})
			Convey("Then the all values should be accessible", func() {
				So(p.Values(), ShouldResemble, []interface{}{1, 2})
			})
		})
	})

	Convey("Given I have a 2 float parameter", t, func() {

		p := Parameter{
			Type:     ParameterTypeFloat,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			err := p.Parse([]string{"1.1", "2.2"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(p.FloatValue(), ShouldResemble, 1.1)
			})
			Convey("Then the all values should be accessible", func() {
				So(p.Values(), ShouldResemble, []interface{}{1.1, 2.2})
			})
		})
	})

	Convey("Given I have a 2 bool parameter", t, func() {

		p := Parameter{
			Type:     ParameterTypeBool,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			err := p.Parse([]string{"true", "false", "yes"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(p.BoolValue(), ShouldResemble, true)
			})
			Convey("Then the all values should be accessible", func() {
				So(p.Values(), ShouldResemble, []interface{}{true, false, true})
			})
		})
	})

	Convey("Given I have a 2 duration parameter", t, func() {

		p := Parameter{
			Type:     ParameterTypeDuration,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			err := p.Parse([]string{"2s", "2h"})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(p.DurationValue(), ShouldResemble, 2*time.Second)
			})
			Convey("Then the all values should be accessible", func() {
				So(p.Values(), ShouldResemble, []interface{}{2 * time.Second, 2 * time.Hour})
			})
		})
	})

	Convey("Given I have a 2 time parameter", t, func() {

		t1 := time.Now()
		t2 := time.Now().Add(-3 * time.Hour)

		p := Parameter{
			Type:     ParameterTypeTime,
			Multiple: true,
		}

		Convey("When I parse it", func() {

			err := p.Parse([]string{t1.Format(time.RFC3339), t2.Format(time.RFC3339)})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the first value should be accessible", func() {
				So(p.TimeValue().Format(time.RFC3339), ShouldResemble, t1.Format(time.RFC3339))
			})
			Convey("Then the all values should be accessible", func() {
				So(len(p.Values()), ShouldEqual, 2)
				So(p.Values()[0].(time.Time).Format(time.RFC3339), ShouldResemble, t1.Format(time.RFC3339))
				So(p.Values()[1].(time.Time).Format(time.RFC3339), ShouldResemble, t2.Format(time.RFC3339))
			})
		})
	})
}
