package genopenapi3

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"go.aporeto.io/regolithe/spec"
)

var noDesc = "n/a"

type operationConfig struct {
	id     string
	schema string
	tags   []string
}

func (c *converter) convertRelationsForRootSpec(relations []*spec.Relation) map[string]*openapi3.PathItem {

	paths := make(map[string]*openapi3.PathItem)

	for _, relation := range relations {

		if relation.Get == nil && relation.Create == nil {
			continue
		}

		model := relation.Specification().Model()
		tags := []string{model.Group, model.Package}

		pathItem := &openapi3.PathItem{
			Get: c.convertRelationActionToGetAll(
				relation.Get,
				operationConfig{
					id:     "get-all-" + model.ResourceName,
					schema: relation.RestName,
					tags:   tags,
				},
			),
			Post: c.convertRelationActionToPost(
				relation.Create,
				operationConfig{
					id:     "create-a-new-" + relation.RestName,
					schema: relation.RestName,
					tags:   tags,
				},
			),
		}

		uri := "/" + model.ResourceName
		paths[uri] = pathItem
	}

	return paths
}

func (c *converter) convertRelationsForNonRootSpec(resourceName string, relations []*spec.Relation) map[string]*openapi3.PathItem {

	paths := make(map[string]*openapi3.PathItem)
	parentRestName := c.resourceToRest[resourceName]

	for _, relation := range relations {

		if relation.Get == nil && relation.Create == nil {
			continue
		}

		childRestName := relation.RestName
		childModel := c.inSpecSet.Specification(childRestName).Model()
		childResourceName := childModel.ResourceName
		tags := []string{childModel.Group, childModel.Package}

		pathItem := &openapi3.PathItem{
			Get: c.convertRelationActionToGetAll(
				relation.Get,
				operationConfig{
					id:     "get-all-" + childResourceName + "-for-a-given-" + parentRestName,
					tags:   tags,
					schema: childRestName,
				},
			),
			Post: c.convertRelationActionToPost(
				relation.Create,
				operationConfig{
					id:     "create-a-new-" + childRestName + "-for-a-given-" + parentRestName,
					tags:   tags,
					schema: childRestName,
				},
			),
		}

		c.insertParamID(&pathItem.Parameters)

		uri := fmt.Sprintf("/%s/{%s}/%s", resourceName, paramNameID, childResourceName)
		paths[uri] = pathItem
	}

	return paths
}

func (c *converter) convertRelationsForNonRootModel(model *spec.Model) map[string]*openapi3.PathItem {

	if model.Get == nil && model.Update == nil && model.Delete == nil {
		return nil
	}

	tags := []string{model.Group, model.Package}

	pathItem := &openapi3.PathItem{
		Get: c.convertRelationActionToGetByID(
			model.Get,
			operationConfig{
				id:     fmt.Sprintf("get-%s-by-ID", model.RestName),
				tags:   tags,
				schema: model.RestName,
			},
		),
		Delete: c.convertRelationActionToDeleteByID(
			model.Delete,
			operationConfig{
				id:     fmt.Sprintf("delete-%s-by-ID", model.RestName),
				tags:   tags,
				schema: model.RestName,
			},
		),
		Put: c.convertRelationActionToPutByID(
			model.Update,
			operationConfig{
				id:     fmt.Sprintf("update-%s-by-ID", model.RestName),
				tags:   tags,
				schema: model.RestName,
			},
		),
	}
	c.insertParamID(&pathItem.Parameters)

	uri := fmt.Sprintf("/%s/{%s}", model.ResourceName, paramNameID)
	pathItems := map[string]*openapi3.PathItem{uri: pathItem}
	return pathItems
}

func (c *converter) convertRelationActionToGetAll(relationAction *spec.RelationAction, cfg operationConfig) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	respBodySchema := openapi3.NewArraySchema()
	respBodySchema.Items = openapi3.NewSchemaRef("#/components/schemas/"+cfg.schema, nil)

	op := &openapi3.Operation{
		OperationID: cfg.id,
		Tags:        cfg.tags,
		Description: relationAction.Description,
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: &noDesc,
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

func (c *converter) convertRelationActionToPost(relationAction *spec.RelationAction, cfg operationConfig) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	schemaRef := openapi3.NewSchemaRef("#/components/schemas/"+cfg.schema, nil)

	op := &openapi3.Operation{
		OperationID: cfg.id,
		Tags:        cfg.tags,
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
					Description: &noDesc,
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

func (c *converter) convertRelationActionToGetByID(relationAction *spec.RelationAction, cfg operationConfig) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	respBodySchemaRef := openapi3.NewSchemaRef("#/components/schemas/"+cfg.schema, nil)

	op := &openapi3.Operation{
		OperationID: cfg.id,
		Tags:        cfg.tags,
		Description: relationAction.Description,
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: &noDesc,
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

func (c *converter) convertRelationActionToDeleteByID(relationAction *spec.RelationAction, cfg operationConfig) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	respBodySchemaRef := openapi3.NewSchemaRef("#/components/schemas/"+cfg.schema, nil)

	op := &openapi3.Operation{
		OperationID: cfg.id,
		Tags:        cfg.tags,
		Description: relationAction.Description,
		Responses: openapi3.Responses{
			"200": &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Description: &noDesc,
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

func (c *converter) convertRelationActionToPutByID(relationAction *spec.RelationAction, cfg operationConfig) *openapi3.Operation {

	if relationAction == nil {
		return nil
	}

	schemaRef := openapi3.NewSchemaRef("#/components/schemas/"+cfg.schema, nil)

	op := &openapi3.Operation{
		OperationID: cfg.id,
		Tags:        cfg.tags,
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
					Description: &noDesc,
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
