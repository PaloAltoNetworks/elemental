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

	params := openapi3.NewParameters()
	for _, e := range paramDef.Entries {
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
	param.Example = entry.ExampleValue
	param.In = in

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
		param.Schema = openapi3.NewArraySchema().WithEnum(enumVals).NewRef()

	default:
		return nil // TODO: better handling? error?
	}

	if entry.Multiple {
		_ = "make linter happy"
		// TODO? How should I exactly indicate that we allow multiple query params
		// of the same key? I think that is the assumption by default!
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
