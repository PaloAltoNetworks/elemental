package main

import (
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"go.aporeto.io/regolithe/spec"
)

type openapi3Converter struct {
	inSpecSet  spec.SpecificationSet
	outRootDoc openapi3.T
}

func newOpenapi3Converter(inSpecSet spec.SpecificationSet) *openapi3Converter {

	sc := &openapi3Converter{
		inSpecSet: inSpecSet,
		outRootDoc: openapi3.T{
			Paths:      openapi3.Paths{},
			Components: openapi3.NewComponents(),
		},
	}

	sc.outRootDoc.Components.Schemas = make(openapi3.Schemas)
	return sc
}

func (sc *openapi3Converter) do() (string, error) {

	for _, s := range sc.inSpecSet.Specifications() {
		model, err := sc.convertModel(s)
		if err != nil {
			return "", fmt.Errorf("spec model '%s': %w", s.Model().RestName, err)
		}
		sc.outRootDoc.Components.Schemas[s.Model().RestName] = model
	}

	bytes, err := json.MarshalIndent(sc.outRootDoc, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshaling openapi3 document: %w", err)
	}
	return string(bytes), nil
}

func (sc *openapi3Converter) convertModel(s spec.Specification) (*openapi3.SchemaRef, error) {

	schema := openapi3.NewObjectSchema()
	schema.Properties = make(map[string]*openapi3.SchemaRef)

	for _, specAttr := range s.Attributes("") { // TODO: figure out versions
		attr, err := sc.convertAttribute(specAttr)
		if err != nil {
			return nil, fmt.Errorf("attribute '%s': %w", specAttr.Name, err)
		}
		schema.Properties[specAttr.Name] = attr
	}

	return openapi3.NewSchemaRef("", schema), nil
}

func (sc *openapi3Converter) convertAttribute(attr *spec.Attribute) (*openapi3.SchemaRef, error) {

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
		attr, err := sc.convertAttribute(&spec.Attribute{Type: spec.AttributeType(attr.SubType)})
		attrSchema.Items = attr
		return attrSchema.NewRef(), err // do not wrap error to avoid recursive wrapping

	case spec.AttributeTypeRef:
		return openapi3.NewSchemaRef("#/components/schemas/"+attr.SubType, nil), nil

	case spec.AttributeTypeRefList:
		attrSchema := openapi3.NewArraySchema()
		attr, err := sc.convertAttribute(&spec.Attribute{Type: spec.AttributeTypeRef, SubType: attr.SubType})
		attrSchema.Items = attr
		return attrSchema.NewRef(), err // do not wrap error to avoid recursive wrapping

	case spec.AttributeTypeRefMap:
		attrSchema := openapi3.NewObjectSchema()
		attr, err := sc.convertAttribute(&spec.Attribute{Type: spec.AttributeTypeRef, SubType: attr.SubType})
		attrSchema.AdditionalProperties = attr
		return attrSchema.NewRef(), err // do not wrap error to avoid recursive wrapping

	case spec.AttributeTypeExt:
		mapping, err := sc.inSpecSet.TypeMapping().Mapping("openapi3", attr.SubType)
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

func generatorOpenapi3(sets []spec.SpecificationSet, out string) error {
	_ = out // make linter happy for now
	set := sets[0]
	converter := newOpenapi3Converter(set)
	doc, err := converter.do()
	if err != nil {
		return fmt.Errorf("error generating openapi3 document from spec set '%s': %w", set.Configuration().Name, err)
	}
	// TODO: write doc to file
	fmt.Println(doc)
	return nil
}

var _ = generatorOpenapi3
