package genopenapi3

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"go.aporeto.io/regolithe/spec"
)

func (c *converter) convertRelationsForRootSpec(relations []*spec.Relation) map[string]*openapi3.PathItem {

	paths := make(map[string]*openapi3.PathItem)

	for _, relation := range relations {

		if relation.Get == nil && relation.Create == nil {
			continue
		}

		pathItem := &openapi3.PathItem{
			Get:  c.convertRelationActionToGetAll(relation.Get, relation.RestName),
			Post: c.convertRelationActionToPost(relation.Create, relation.RestName),
		}

		uri := "/" + c.inSpecSet.Specification(relation.RestName).Model().ResourceName
		paths[uri] = pathItem
	}

	return paths
}

func (c *converter) convertRelationsForNonRootSpec(resourceName string, relations []*spec.Relation) map[string]*openapi3.PathItem {

	paths := make(map[string]*openapi3.PathItem)

	for _, relation := range relations {

		if relation.Get == nil && relation.Create == nil {
			continue
		}

		pathItem := &openapi3.PathItem{
			Get:  c.convertRelationActionToGetAll(relation.Get, relation.RestName),
			Post: c.convertRelationActionToPost(relation.Create, relation.RestName),
		}

		c.insertParamID(&pathItem.Parameters)
		relatedResourceName := c.inSpecSet.Specification(relation.RestName).Model().ResourceName
		uri := fmt.Sprintf("/%s/{%s}/%s", resourceName, paramNameID, relatedResourceName)
		paths[uri] = pathItem
	}

	return paths
}

func (c *converter) convertRelationsForNonRootModel(model *spec.Model) map[string]*openapi3.PathItem {

	if model.Get == nil && model.Update == nil && model.Delete == nil {
		return nil
	}

	pathItem := &openapi3.PathItem{
		Get:    c.convertRelationActionToGetByID(model.Get, model.RestName),
		Delete: c.convertRelationActionToDeleteByID(model.Delete, model.RestName),
		Put:    c.convertRelationActionToPutByID(model.Update, model.RestName),
	}
	c.insertParamID(&pathItem.Parameters)

	uri := fmt.Sprintf("/%s/{%s}", model.ResourceName, paramNameID)
	pathItems := map[string]*openapi3.PathItem{uri: pathItem}
	return pathItems
}

func (c *converter) convertRelationActionToGetAll(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	respBodySchema := openapi3.NewArraySchema()
	respBodySchema.Items = openapi3.NewSchemaRef("#/components/schemas/"+restName, nil)

	op := &openapi3.Operation{
		Description: relationAction.Description,
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: respBodySchema.NewRef(),
						},
					},
				},
			},
			// TODO: more responses like 422, 500, etc if needed
		},
		Parameters: c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}

func (c *converter) convertRelationActionToPost(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	schemaRef := openapi3.NewSchemaRef("#/components/schemas/"+restName, nil)

	op := &openapi3.Operation{
		Description: relationAction.Description,
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: schemaRef,
					},
				},
			},
		},
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: schemaRef,
						},
					},
				},
			},
			// TODO: more responses like 422, 500, etc if needed
		},
		Parameters: c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}

func (c *converter) convertRelationActionToGetByID(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	respBodySchemaRef := openapi3.NewSchemaRef("#/components/schemas/"+restName, nil)

	op := &openapi3.Operation{
		Description: relationAction.Description,
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: respBodySchemaRef,
						},
					},
				},
			},
			// TODO: more responses like 422, 500, etc if needed
		},
		Parameters: c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}

func (c *converter) convertRelationActionToDeleteByID(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	respBodySchemaRef := openapi3.NewSchemaRef("#/components/schemas/"+restName, nil)

	op := &openapi3.Operation{
		Description: relationAction.Description,
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: respBodySchemaRef,
						},
					},
				},
			},
			// TODO: more responses like 422, 500, etc if needed
		},
		Parameters: c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}

func (c *converter) convertRelationActionToPutByID(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	schemaRef := openapi3.NewSchemaRef("#/components/schemas/"+restName, nil)

	op := &openapi3.Operation{
		Description: relationAction.Description,
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: schemaRef,
					},
				},
			},
		},
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.Content{
						"application/json": &openapi3.MediaType{
							Schema: schemaRef,
						},
					},
				},
			},
			// TODO: more responses like 422, 500, etc if needed
		},
		Parameters: c.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}
