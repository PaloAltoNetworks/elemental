package elemental

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

const invalidParamMsg = "Invalid Parameter"

// ParameterType represents the various type for a parameter.
type ParameterType string

// Various values for ParameterType.
const (
	ParameterTypeString   ParameterType = "string"
	ParameterTypeInt      ParameterType = "integer"
	ParameterTypeFloat    ParameterType = "float"
	ParameterTypeBool     ParameterType = "boolean"
	ParameterTypeEnum     ParameterType = "enum"
	ParameterTypeTime     ParameterType = "time"
	ParameterTypeDuration ParameterType = "duration"
)

// A ParameterDefinition represent a parameter definition that can
// be transformed into a Parameter.
type ParameterDefinition struct {
	Name           string
	Type           ParameterType
	AllowedChoices []string
	DefaultValue   interface{}
	Required       bool
	Multiple       bool

	values []interface{}
}

// Parse parses the given value against the parameter definition
func (p *ParameterDefinition) Parse(values []string) (*Parameter, error) {

	if !p.Multiple && len(values) > 1 {
		return nil, NewError(
			invalidParamMsg,
			fmt.Sprintf("Parameter '%s' must be sent only once", p.Name),
			"elemental",
			http.StatusBadRequest,
		)
	}

	if p.Required {

		if len(values) == 0 {
			return nil, NewError(
				invalidParamMsg,
				fmt.Sprintf("Parameter '%s' is required", p.Name),
				"elemental",
				http.StatusBadRequest,
			)
		}

		for _, v := range values {
			if v == "" {
				return nil, NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' is required", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}
		}
	}

	var vs []interface{}

	for _, v := range values {
		switch p.Type {

		case ParameterTypeString:
			vs = append(vs, v)

		case ParameterTypeInt:
			parsed, err := strconv.Atoi(v)
			if err != nil {
				return nil, NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must be an integer", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}
			vs = append(vs, parsed)

		case ParameterTypeFloat:

			parsed, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return nil, NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must be a float", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}
			vs = append(vs, parsed)

		case ParameterTypeBool:

			switch strings.ToLower(v) {
			case "true", "yes", "1", "":
				vs = append(vs, true)
			case "false", "no", "0":
				vs = append(vs, false)
			default:
				return nil, NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must be a boolean", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}

		case ParameterTypeEnum:

			var matched bool
			for _, allowed := range p.AllowedChoices {
				if v == allowed {
					matched = true
					break
				}
			}

			if !matched {
				return nil, NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must be one of '%s'", p.Name, strings.Join(p.AllowedChoices, ", ")),
					"elemental",
					http.StatusBadRequest,
				)
			}

			vs = append(vs, v)

		case ParameterTypeDuration:
			d, err := time.ParseDuration(v)
			if err != nil {
				return nil, NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must be a valid duration", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}

			vs = append(vs, d)

		case ParameterTypeTime:
			t, err := dateparse.ParseAny(v)
			if err != nil {
				return nil, NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must be a valid date", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}

			vs = append(vs, t)

		default:
			panic(fmt.Sprintf("unknown parameter type: '%s'", p.Type))
		}
	}

	return &Parameter{
		ptype:  p.Type,
		values: vs,
	}, nil
}

// Parameters represents a set of Parameters.
type Parameters map[string]Parameter

// Get returns the Parameter with the given name
func (p Parameters) Get(name string) (Parameter, bool) {
	param, ok := p[name]
	return param, ok
}

// Validate validates if the Parameters matches the given requirement.
func (p Parameters) Validate(r ParametersRequirement) error {

	if len(r.match) == 0 {
		return nil
	}

	var innerMatch int
	var outerMatch int

	for _, clauses := range r.match {

		for _, ands := range clauses {

			innerMatch = 0

			for _, k := range ands {

				if _, ok := p[k]; ok {
					innerMatch++
				}

				if innerMatch == len(ands) {
					outerMatch++
				}
			}
		}
	}

	if outerMatch == len(r.match) {
		return nil
	}

	return NewError("Bad Request", "Some required parameters are missing", "elemental", http.StatusBadRequest)
}

// A ParametersRequirement represents a list of ands of list of ors
// that must be passed together.
type ParametersRequirement struct {
	match [][][]string
}

// NewParametersRequirement returns a new ParametersRequirement.
func NewParametersRequirement(match [][][]string) ParametersRequirement {
	return ParametersRequirement{
		match: match,
	}
}

// A Parameter represent one parameter that can be sent with a query.
type Parameter struct {
	ptype  ParameterType
	values []interface{}
}

// NewParameter returns a new Parameter.
func NewParameter(ptype ParameterType, values ...interface{}) Parameter {

	return Parameter{
		ptype:  ptype,
		values: values,
	}
}

// StringValue returns the value as a string.
func (p Parameter) StringValue() string {

	if len(p.values) == 0 || (p.ptype != ParameterTypeString && p.ptype != ParameterTypeEnum) {
		return ""
	}

	return p.values[0].(string)
}

// IntValue returns the value as a int.
func (p Parameter) IntValue() int {

	if len(p.values) == 0 || p.ptype != ParameterTypeInt {
		return 0
	}

	return p.values[0].(int)
}

// FloatValue returns the value as a float.
func (p Parameter) FloatValue() float64 {

	if len(p.values) == 0 || p.ptype != ParameterTypeFloat {
		return 0.0
	}

	return p.values[0].(float64)
}

// BoolValue returns the value as a bool.
func (p Parameter) BoolValue() bool {

	if len(p.values) == 0 || p.ptype != ParameterTypeBool {
		return false
	}

	return p.values[0].(bool)
}

// DurationValue returns the value as a time.Duration.
func (p Parameter) DurationValue() time.Duration {

	if len(p.values) == 0 || p.ptype != ParameterTypeDuration {
		return 0
	}

	return p.values[0].(time.Duration)
}

// TimeValue returns the value as a time.Time.
func (p Parameter) TimeValue() time.Time {

	if len(p.values) == 0 || p.ptype != ParameterTypeTime {
		return time.Time{}
	}

	return p.values[0].(time.Time)
}

// Values returns all the parsed values
func (p Parameter) Values() []interface{} {

	return p.values
}
