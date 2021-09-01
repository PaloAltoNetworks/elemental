package genopenapi3

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
	"go.aporeto.io/regolithe/spec"
)

func (c *converter) convertParamDefAsQueryParams(paramDef *spec.ParameterDefinition) openapi3.Parameters {

	if paramDef == nil {
		return nil
	}

	seen := make(map[string]struct{})

	params := openapi3.NewParameters()
	for _, e := range paramDef.Entries {

		if _, ok := seen[e.Name]; ok {
			continue
		}
		seen[e.Name] = struct{}{}

		p := c.convertParam(e, openapi3.ParameterInQuery)
		params = append(params, p)
	}

	sort.Slice(params, func(i, j int) bool {
		return params[i].Value.Name < params[j].Value.Name
	})

	return params
}

func (*converter) convertParam(entry *spec.Parameter, in string) *openapi3.ParameterRef {

	param := openapi3.NewQueryParameter(entry.Name)
	param.Description = entry.Description
	param.In = in
	param.Example = entry.ExampleValue
	if param.Example == nil {
		param.Example = entry.DefaultValue
	}

	switch entry.Type {
	case spec.ParameterTypeInt:
		param.Schema = openapi3.NewIntegerSchema().NewRef()

	case spec.ParameterTypeBool:
		param.Schema = openapi3.NewBoolSchema().NewRef()

	case spec.ParameterTypeString:
		param.Schema = openapi3.NewStringSchema().NewRef()

	case spec.ParameterTypeFloat:
		param.Schema = openapi3.NewFloat64Schema().NewRef()

	case spec.ParameterTypeTime:
		param.Schema = openapi3.NewDateTimeSchema().NewRef()

	case spec.ParameterTypeDuration:
		param.Schema = openapi3.NewStringSchema().NewRef() // TODO: this needs to be verified

	case spec.ParameterTypeEnum:
		enumVals := make([]interface{}, len(entry.AllowedChoices))
		for i, val := range entry.AllowedChoices {
			enumVals[i] = val
		}
		param.Schema = openapi3.NewSchema().WithEnum(enumVals...).NewRef()

	default:
		return nil // TODO: better handling? error?
	}

	ref := &openapi3.ParameterRef{
		Value: param,
	}

	return ref
}

func (*converter) insertParamID(params *openapi3.Parameters) {
	paramID := openapi3.NewPathParameter(paramNameID)
	paramID.Schema = openapi3.NewStringSchema().NewRef()
	*params = append(*params, &openapi3.ParameterRef{Value: paramID})
}
