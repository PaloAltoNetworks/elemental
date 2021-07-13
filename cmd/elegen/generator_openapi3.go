package main

import (
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"go.aporeto.io/regolithe/spec"
)

const paramNameID = "id"

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

func (sc *openapi3Converter) Do() (string, error) {

	for _, s := range sc.inSpecSet.Specifications() {
		if err := sc.processSpec(s); err != nil {
			return "", fmt.Errorf("unable to to process spec: %w", err)
		}
	}

	doc, err := json.MarshalIndent(sc.outRootDoc, "", "  ")
	if err != nil {
		return "", fmt.Errorf("marshaling openapi3 document: %w", err)
	}
	return string(doc), nil
}

func (sc *openapi3Converter) processSpec(s spec.Specification) error {

	if s.Model().IsRoot {
		pathItems := sc.convertRelationsForRootSpec(s.Relations())
		for path, item := range pathItems {
			sc.outRootDoc.Paths[path] = item
		}
		// we don't care about root model's relations, so we are done for root spec
		return nil
	}

	schema, err := sc.convertModel(s)
	if err != nil {
		return fmt.Errorf("model '%s': %w", s.Model().RestName, err)
	}
	sc.outRootDoc.Components.Schemas[s.Model().RestName] = schema

	pathItems := sc.convertRelationsForNonRootModel(s.Model())
	for path, item := range pathItems {
		sc.outRootDoc.Paths[path] = item
	}

	pathItems = sc.convertRelationsForNonRootSpec(s.Model().ResourceName, s.Relations())
	for path, item := range pathItems {
		sc.outRootDoc.Paths[path] = item
	}

	return nil
}

func generatorOpenapi3(sets []spec.SpecificationSet, out string) error {
	_ = out // make linter happy for now
	set := sets[0]
	converter := newOpenapi3Converter(set)
	doc, err := converter.Do()
	if err != nil {
		return fmt.Errorf("error generating openapi3 document from spec set '%s': %w", set.Configuration().Name, err)
	}
	// TODO: write doc to file
	fmt.Println(doc)
	return nil
}

var _ = generatorOpenapi3
