package genopenapi3

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
	"go.aporeto.io/regolithe/spec"
)

func (c *converter) convertModel(s spec.Specification) (*openapi3.SchemaRef, error) {

	schema := openapi3.NewObjectSchema()
	schema.Properties = make(map[string]*openapi3.SchemaRef)

	for _, specAttr := range s.Attributes("") { // TODO: figure out versions

		attr, err := c.convertAttribute(specAttr)
		if err != nil {
			return nil, fmt.Errorf("attribute '%s': %w", specAttr.Name, err)
		}
		schema.Properties[specAttr.Name] = attr

		if specAttr.Required {
			schema.Required = append(schema.Required, specAttr.Name)
		}
	}

	sort.Strings(schema.Required)

	return openapi3.NewSchemaRef("", schema), nil
}

func (c *converter) convertAttribute(attr *spec.Attribute) (schemaRef *openapi3.SchemaRef, err error) {

	defer func() {
		if schemaRef == nil || schemaRef.Value == nil {
			return
		}
		s := schemaRef.Value
		s.Description = attr.Description
		s.Default = attr.DefaultValue
		s.Example = attr.ExampleValue
		s.Deprecated = attr.Deprecated
		s.ReadOnly = attr.ReadOnly
	}()

	switch attr.Type {

	case spec.AttributeTypeString:
		return openapi3.NewStringSchema().NewRef(), nil

	case spec.AttributeTypeInt:
		return openapi3.NewIntegerSchema().NewRef(), nil

	case spec.AttributeTypeFloat:
		return openapi3.NewFloat64Schema().NewRef(), nil

	case spec.AttributeTypeBool:
		return openapi3.NewBoolSchema().NewRef(), nil

	case spec.AttributeTypeTime:
		return openapi3.NewDateTimeSchema().NewRef(), nil

	case spec.AttributeTypeEnum:
		enumVals := make([]interface{}, len(attr.AllowedChoices))
		for i, val := range attr.AllowedChoices {
			enumVals[i] = val
		}
		return openapi3.NewArraySchema().WithEnum(enumVals...).NewRef(), nil

	case spec.AttributeTypeObject:
		return openapi3.NewObjectSchema().NewRef(), nil

	case spec.AttributeTypeList:
		attrSchema := openapi3.NewArraySchema()
		attr, err := c.convertAttribute(&spec.Attribute{Type: spec.AttributeType(attr.SubType)})
		attrSchema.Items = attr
		return attrSchema.NewRef(), err // do not wrap error to avoid recursive wrapping

	case spec.AttributeTypeRef:
		return openapi3.NewSchemaRef("#/components/schemas/"+attr.SubType, nil), nil

	case spec.AttributeTypeRefList:
		attrSchema := openapi3.NewArraySchema()
		attr, err := c.convertAttribute(&spec.Attribute{Type: spec.AttributeTypeRef, SubType: attr.SubType})
		attrSchema.Items = attr
		return attrSchema.NewRef(), err // do not wrap error to avoid recursive wrapping

	case spec.AttributeTypeRefMap:
		attrSchema := openapi3.NewObjectSchema()
		attr, err := c.convertAttribute(&spec.Attribute{Type: spec.AttributeTypeRef, SubType: attr.SubType})
		attrSchema.AdditionalProperties = attr
		return attrSchema.NewRef(), err // do not wrap error to avoid recursive wrapping

	case spec.AttributeTypeExt:
		mapping, err := c.inSpecSet.TypeMapping().Mapping("openapi3", attr.SubType)
		if err != nil {
			return nil, fmt.Errorf("retrieving 'openapi3' type mapping for external attribute subtype '%s': %w", attr.SubType, err)
		}

		attrSchema := new(openapi3.Schema)
		if err := json.Unmarshal([]byte(mapping.Type), attrSchema); err != nil {
			return nil, fmt.Errorf("unmarshaling openapi3 external type mapping '%s': %w", attr.SubType, err)
		}

		return attrSchema.NewRef(), nil
	}

	return nil, fmt.Errorf("unhandled attribute type: '%s'", attr.Type)
}
