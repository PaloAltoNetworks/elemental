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
	splitOutput       bool
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
		outRootDoc:        newOpenAPI3Template(inSpecSet.Configuration()),
	}

	for _, spec := range inSpecSet.Specifications() {
		model := spec.Model()
		c.resourceToRest[model.ResourceName] = model.RestName
	}

	return c
}

func (c *converter) Do(newWriter func(name string) (io.WriteCloser, error)) error {

	for _, s := range c.inSpecSet.Specifications() {
		if err := c.processSpec(s); err != nil {
			return fmt.Errorf("unable to to process spec: %w", err)
		}
	}

	for name, doc := range c.convertedDocs() {
		dest, err := newWriter(name)
		if err != nil {
			return fmt.Errorf("'%s': unable to create write destination: %w", name, err)
		}

		enc := json.NewEncoder(dest)
		enc.SetIndent("", "  ")
		if err := enc.Encode(doc); err != nil {
			return fmt.Errorf("'%s': marshaling openapi3 document: %w", name, err)
		}
	}

	return nil
}

func (c *converter) convertedDocs() map[string]openapi3.T {

	if !c.splitOutput {
		return map[string]openapi3.T{"toplevel": c.outRootDoc}
	}

	out := make(map[string]openapi3.T)
	specConfig := c.inSpecSet.Configuration()
	for name, schema := range c.outRootDoc.Components.Schemas {
		template := newOpenAPI3Template(specConfig)
		template.Components.Schemas[name] = schema
		out[name] = template
	}

	for path, item := range c.outRootDoc.Paths {
		home := strings.Split(strings.Trim(path, "/"), "/")[0]
		home = c.resourceToRest[home]
		out[home].Paths[path] = item
	}
	return out
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

func newOpenAPI3Template(specConfig *spec.Config) openapi3.T {
	return openapi3.T{
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
	}
}
