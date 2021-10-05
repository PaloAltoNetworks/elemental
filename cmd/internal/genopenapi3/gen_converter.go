package genopenapi3

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/getkin/kin-openapi/openapi3"
	"go.aporeto.io/regolithe/spec"
)

const paramNameID = "id"

type converter struct {
	skipPrivateModels bool
	inSpecSet         spec.SpecificationSet
	resourceToRest    map[string]string
	outRootDoc        openapi3.T
}

func newConverter(inSpecSet spec.SpecificationSet, skipPrivateModels bool) *converter {
	specConfig := inSpecSet.Configuration()
	c := &converter{
		skipPrivateModels: skipPrivateModels,
		inSpecSet:         inSpecSet,
		resourceToRest:    make(map[string]string),
		outRootDoc: openapi3.T{
			OpenAPI: "3.0.3",
			Info: &openapi3.Info{
				Title:          specConfig.Name,
				Version:        specConfig.Version,
				Description:    specConfig.Description,
				TermsOfService: "https://localhost/TODO", // TODO
				License: &openapi3.License{
					Name: "TODO",
				},
				Contact: &openapi3.Contact{
					Name:  specConfig.Author,
					URL:   specConfig.URL,
					Email: specConfig.Email,
				},
			},
			Paths: openapi3.Paths{},
			Components: openapi3.Components{
				Schemas: make(openapi3.Schemas),
			},
		},
	}

	for _, spec := range inSpecSet.Specifications() {
		model := spec.Model()
		c.resourceToRest[model.ResourceName] = model.RestName
	}

	return c
}

func (c *converter) Do(dest io.Writer) error {

	for _, s := range c.inSpecSet.Specifications() {
		if err := c.processSpec(s); err != nil {
			return fmt.Errorf("unable to to process spec: %w", err)
		}
	}

	enc := json.NewEncoder(dest)
	enc.SetIndent("", "  ")
	if err := enc.Encode(c.outRootDoc); err != nil {
		return fmt.Errorf("marshaling openapi3 document: %w", err)
	}

	return nil
}

func (c *converter) processSpec(s spec.Specification) error {

	model := s.Model()

	if c.skipPrivateModels && model.Private {
		return nil
	}

	if model.IsRoot {
		pathItems := c.convertRelationsForRootSpec(s.Relations())
		for path, item := range pathItems {
			c.outRootDoc.Paths[path] = item
		}
		// we don't care about root model's relations for now, so we are done for root spec
		return nil
	}

	schema, err := c.convertModel(s)
	if err != nil {
		return fmt.Errorf("model '%s': %w", model.RestName, err)
	}
	c.outRootDoc.Components.Schemas[model.RestName] = schema

	pathItems := c.convertRelationsForNonRootModel(model)
	for path, item := range pathItems {
		c.outRootDoc.Paths[path] = item
	}

	pathItems = c.convertRelationsForNonRootSpec(model.ResourceName, s.Relations())
	for path, item := range pathItems {
		c.outRootDoc.Paths[path] = item
	}

	return nil
}
