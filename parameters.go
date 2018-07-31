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

// A Parameter represent one parameter that can be sent with a query
type Parameter struct {
	Name           string
	Type           ParameterType
	AllowedChoices []string
	DefaultValue   interface{}
	Required       bool
	Multiple       bool

	value []interface{}
}

// StringValue returns the value as a string.
func (p *Parameter) StringValue() string {

	if len(p.value) == 0 || (p.Type != ParameterTypeString && p.Type != ParameterTypeEnum) {
		return ""
	}

	return p.value[0].(string)
}

// IntValue returns the value as a int.
func (p *Parameter) IntValue() int {

	if len(p.value) == 0 || p.Type != ParameterTypeInt {
		return 0
	}

	return p.value[0].(int)
}

// FloatValue returns the value as a float.
func (p *Parameter) FloatValue() float64 {

	if len(p.value) == 0 || p.Type != ParameterTypeFloat {
		return 0.0
	}

	return p.value[0].(float64)
}

// BoolValue returns the value as a bool.
func (p *Parameter) BoolValue() bool {

	if len(p.value) == 0 || p.Type != ParameterTypeBool {
		return false
	}

	return p.value[0].(bool)
}

// DurationValue returns the value as a time.Duration.
func (p *Parameter) DurationValue() time.Duration {

	if len(p.value) == 0 || p.Type != ParameterTypeDuration {
		return 0
	}

	return p.value[0].(time.Duration)
}

// TimeValue returns the value as a time.Time.
func (p *Parameter) TimeValue() time.Time {

	if len(p.value) == 0 || p.Type != ParameterTypeTime {
		return time.Time{}
	}

	return p.value[0].(time.Time)
}

// Values returns all the parsed values
func (p *Parameter) Values() []interface{} {

	return p.value
}

// Parse parses the given value against the parameter definition
func (p *Parameter) Parse(values []string) (err error) {

	if !p.Multiple && len(values) > 1 {
		return NewError(
			invalidParamMsg,
			fmt.Sprintf("Parameter '%s' must be send only once", p.Name),
			"elemental",
			http.StatusBadRequest,
		)
	}

	if p.Required {

		if len(values) == 0 {
			return NewError(
				invalidParamMsg,
				fmt.Sprintf("Parameter '%s' is required", p.Name),
				"elemental",
				http.StatusBadRequest,
			)
		}

		for _, v := range values {
			if v == "" {
				return NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' is required", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}
		}
	}

	for _, v := range values {
		switch p.Type {

		case ParameterTypeString:
			p.value = append(p.value, v)

		case ParameterTypeInt:
			parsed, err := strconv.Atoi(v)
			if err != nil {
				return NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must be an integer", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}
			p.value = append(p.value, parsed)

		case ParameterTypeFloat:

			parsed, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must be an float", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}
			p.value = append(p.value, parsed)

		case ParameterTypeBool:

			switch strings.ToLower(v) {
			case "true", "yes", "1", "":
				p.value = append(p.value, true)
			case "false", "no", "0":
				p.value = append(p.value, false)
			default:
				return NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must be an boolean", p.Name),
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
				return NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must one of '%s'", p.Name, strings.Join(p.AllowedChoices, ", ")),
					"elemental",
					http.StatusBadRequest,
				)
			}

			p.value = append(p.value, v)

		case ParameterTypeDuration:
			d, err := time.ParseDuration(v)
			if err != nil {
				return NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must a valid duration", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}

			p.value = append(p.value, d)

		case ParameterTypeTime:
			t, err := dateparse.ParseAny(v)
			if err != nil {
				return NewError(
					invalidParamMsg,
					fmt.Sprintf("Parameter '%s' must a valid date", p.Name),
					"elemental",
					http.StatusBadRequest,
				)
			}

			p.value = append(p.value, t)

		default:
			panic(fmt.Sprintf("unknown parameter type: '%s'", p.Type))
		}
	}

	return nil
}
