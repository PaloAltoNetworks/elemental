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
	DefaultValue   string
	Multiple       bool
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

	var vs []interface{} // nolint: prealloc
	for _, v := range values {
		out, err := parse(p.Type, p.Name, v, p.AllowedChoices)
		if err != nil {
			return nil, err
		}
		vs = append(vs, out)
	}

	var dv interface{}
	if p.DefaultValue != "" && len(vs) == 0 {
		out, err := parse(p.Type, p.Name, p.DefaultValue, p.AllowedChoices)
		if err != nil {
			return nil, err
		}
		dv = out
	}

	return &Parameter{
		ptype:        p.Type,
		values:       vs,
		defaultValue: dv,
	}, nil
}

// Parameters represents a set of Parameters.
type Parameters map[string]Parameter

// Get returns the Parameter with the given name
func (p Parameters) Get(name string) Parameter {
	return p[name]
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

				if _, ok := p[k]; ok && len(p[k].values) > 0 {
					innerMatch++
				}
			}

			if innerMatch == len(ands) {
				outerMatch++
			}
		}
	}

	if outerMatch == len(r.match) {
		return nil
	}

	return NewError("Bad Request", fmt.Sprintf("Missing Required parameters: `%s`", r.String()), "elemental", http.StatusBadRequest)
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

func (r ParametersRequirement) String() string {

	var out string
	for i, lvl1 := range r.match {
		if len(r.match) > 1 {
			out += "("
		}
		for j, lvl2 := range lvl1 {
			if len(lvl1) > 1 {
				out += "("
			}
			out += strings.Join(lvl2, " and ")
			if len(lvl1) > 1 {
				out += ")"
			}
			if j+1 != len(lvl1) {
				out += " or "
			}
		}
		if len(r.match) > 1 {
			out += ")"
		}
		if i+1 != len(r.match) {
			out += " and "
		}
	}

	return out
}

// A Parameter represent one parameter that can be sent with a query.
type Parameter struct {
	ptype        ParameterType
	values       []interface{}
	defaultValue interface{}
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

	if (len(p.values) == 0 && p.defaultValue == nil) || (p.ptype != ParameterTypeString && p.ptype != ParameterTypeEnum) {
		return ""
	}

	if len(p.values) == 0 {
		return p.defaultValue.(string)
	}

	return p.values[0].(string)
}

// StringValues returns all the values as a []string.
func (p Parameter) StringValues() []string {

	if len(p.values) == 0 || (p.ptype != ParameterTypeString && p.ptype != ParameterTypeEnum) {
		return nil
	}

	out := make([]string, len(p.values))
	for i := range p.values {
		out[i] = p.values[i].(string)
	}

	return out
}

// IntValue returns the value as a int.
func (p Parameter) IntValue() int {

	if (len(p.values) == 0 && p.defaultValue == nil) || p.ptype != ParameterTypeInt {
		return 0
	}

	if len(p.values) == 0 {
		return p.defaultValue.(int)
	}

	return p.values[0].(int)
}

// IntValues returns all the values as a []int.
func (p Parameter) IntValues() []int {

	if len(p.values) == 0 || p.ptype != ParameterTypeInt {
		return nil
	}

	out := make([]int, len(p.values))
	for i := range p.values {
		out[i] = p.values[i].(int)
	}

	return out
}

// FloatValue returns the value as a float.
func (p Parameter) FloatValue() float64 {

	if (len(p.values) == 0 && p.defaultValue == nil) || p.ptype != ParameterTypeFloat {
		return 0.0
	}

	if len(p.values) == 0 {
		return p.defaultValue.(float64)
	}

	return p.values[0].(float64)
}

// FloatValues returns all the values as a []float64.
func (p Parameter) FloatValues() []float64 {

	if len(p.values) == 0 || p.ptype != ParameterTypeFloat {
		return nil
	}

	out := make([]float64, len(p.values))
	for i := range p.values {
		out[i] = p.values[i].(float64)
	}

	return out
}

// BoolValue returns the value as a bool.
func (p Parameter) BoolValue() bool {

	if (len(p.values) == 0 && p.defaultValue == nil) || p.ptype != ParameterTypeBool {
		return false
	}

	if len(p.values) == 0 {
		return p.defaultValue.(bool)
	}

	return p.values[0].(bool)
}

// BoolValues returns all the values as a []bool.
func (p Parameter) BoolValues() []bool {

	if len(p.values) == 0 || p.ptype != ParameterTypeBool {
		return nil
	}

	out := make([]bool, len(p.values))
	for i := range p.values {
		out[i] = p.values[i].(bool)
	}

	return out
}

// DurationValue returns the value as a time.Duration.
func (p Parameter) DurationValue() time.Duration {

	if (len(p.values) == 0 && p.defaultValue == nil) || p.ptype != ParameterTypeDuration {
		return 0
	}

	if len(p.values) == 0 {
		return p.defaultValue.(time.Duration)
	}

	return p.values[0].(time.Duration)
}

// DurationValues returns all the values as a []time.Duration.
func (p Parameter) DurationValues() []time.Duration {

	if len(p.values) == 0 || p.ptype != ParameterTypeDuration {
		return nil
	}

	out := make([]time.Duration, len(p.values))
	for i := range p.values {
		out[i] = p.values[i].(time.Duration)
	}

	return out
}

// TimeValue returns the value as a time.Time.
func (p Parameter) TimeValue() time.Time {

	if (len(p.values) == 0 && p.defaultValue == nil) || p.ptype != ParameterTypeTime {
		return time.Time{}
	}

	if len(p.values) == 0 {
		return p.defaultValue.(time.Time)
	}

	return p.values[0].(time.Time)
}

// TimeValues returns all the values as a []time.Time.
func (p Parameter) TimeValues() []time.Time {

	if len(p.values) == 0 || p.ptype != ParameterTypeTime {
		return nil
	}

	out := make([]time.Time, len(p.values))
	for i := range p.values {
		out[i] = p.values[i].(time.Time)
	}

	return out
}

// Values returns all the parsed values
func (p Parameter) Values() []interface{} {

	return p.values
}

func parse(ptype ParameterType, pname string, v string, allowedChoices []string) (out interface{}, err error) {

	switch ptype {

	case ParameterTypeString:
		return v, nil

	case ParameterTypeInt:
		parsed, err := strconv.Atoi(v)
		if err != nil {
			return nil, NewError(
				invalidParamMsg,
				fmt.Sprintf("Parameter '%s' must be an integer, got '%s'", pname, v),
				"elemental",
				http.StatusBadRequest,
			)
		}
		return parsed, nil

	case ParameterTypeFloat:

		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, NewError(
				invalidParamMsg,
				fmt.Sprintf("Parameter '%s' must be a float, got '%s'", pname, v),
				"elemental",
				http.StatusBadRequest,
			)
		}
		return parsed, nil

	case ParameterTypeBool:

		switch strings.ToLower(v) {
		case "true", "yes", "1":
			return true, nil
		case "false", "no", "0", "":
			return false, nil
		default:
			return nil, NewError(
				invalidParamMsg,
				fmt.Sprintf("Parameter '%s' must be a boolean, got '%s'", pname, v),
				"elemental",
				http.StatusBadRequest,
			)
		}

	case ParameterTypeEnum:

		var matched bool
		for _, allowed := range allowedChoices {
			if v == allowed {
				matched = true
				break
			}
		}

		if !matched {
			return nil, NewError(
				invalidParamMsg,
				fmt.Sprintf("Parameter '%s' must be one of '%s', got '%s'", pname, strings.Join(allowedChoices, ", "), v),
				"elemental",
				http.StatusBadRequest,
			)
		}

		return v, nil

	case ParameterTypeDuration:
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, NewError(
				invalidParamMsg,
				fmt.Sprintf("Parameter '%s' must be a valid duration, got '%s'", pname, v),
				"elemental",
				http.StatusBadRequest,
			)
		}

		return d, nil

	case ParameterTypeTime:
		t, err := dateparse.ParseAny(v)
		if err != nil {
			return nil, NewError(
				invalidParamMsg,
				fmt.Sprintf("Parameter '%s' must be a valid date, got '%s'", pname, v),
				"elemental",
				http.StatusBadRequest,
			)
		}

		return t, nil

	default:
		panic(fmt.Sprintf("unknown parameter type: '%s'", ptype))
	}
}
