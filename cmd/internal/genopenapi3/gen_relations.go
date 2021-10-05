package genopenapi3

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"go.aporeto.io/regolithe/spec"
)

var noDesc = "n/a"

func (c *converter) convertRelationsForRootSpec(relations []*spec.Relation) map[string]*openapi3.PathItem {

	paths := make(map[string]*openapi3.PathItem)

	for _, relation := range relations {

		if relation.Get == nil && relation.Create == nil {
			continue
		}

		pathItem := &openapi3.PathItem{
			Get:  c.extractOperationGetAll("", relation),
			Post: c.extractOperationPost("", relation),
		}

		model := relation.Specification().Model()
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

		pathItem := &openapi3.PathItem{
			Get:  c.extractOperationGetAll(parentRestName, relation),
			Post: c.extractOperationPost(parentRestName, relation),
		}

		c.insertParamID(&pathItem.Parameters)

		childModel := c.inSpecSet.Specification(relation.RestName).Model()
		uri := fmt.Sprintf("/%s/{%s}/%s", resourceName, paramNameID, childModel.ResourceName)
		paths[uri] = pathItem
	}

	return paths
}

func (c *converter) convertRelationsForNonRootModel(model *spec.Model) map[string]*openapi3.PathItem {

	if model.Get == nil && model.Update == nil && model.Delete == nil {
		return nil
	}

	pathItem := &openapi3.PathItem{
		Get:    c.extractOperationGetByID(model),
		Delete: c.extractOperationDeleteByID(model),
		Put:    c.extractOperationPutByID(model),
	}
	c.insertParamID(&pathItem.Parameters)

	uri := fmt.Sprintf("/%s/{%s}", model.ResourceName, paramNameID)
	pathItems := map[string]*openapi3.PathItem{uri: pathItem}
	return pathItems
}

func (c *converter) extractOperationGetAll(parentRestName string, relation *spec.Relation) *openapi3.Operation {

	if relation == nil || relation.Get == nil {
		return nil
	}
	relationAction := relation.Get

	model := relation.Specification().Model()

	respBodySchema := openapi3.NewArraySchema()
	respBodySchema.Items = openapi3.NewSchemaRef("#/components/schemas/"+model.RestName, nil)

	op := &openapi3.Operation{
		OperationID: "get-all-" + model.ResourceName,
		Tags:        []string{model.Group, model.Package},
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

	if parentRestName != "" {
		op.OperationID = "get-all-" + model.ResourceName + "-for-a-given-" + parentRestName
	}
	return op
}

func (c *converter) extractOperationPost(parentRestName string, relation *spec.Relation) *openapi3.Operation {

	if relation == nil || relation.Create == nil {
		return nil
	}
	relationAction := relation.Create

	model := relation.Specification().Model()

	schemaRef := openapi3.NewSchemaRef("#/components/schemas/"+relation.RestName, nil)

	op := &openapi3.Operation{
		OperationID: "create-a-new-" + relation.RestName,
		Tags:        []string{model.Group, model.Package},
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

	if parentRestName != "" {
		op.OperationID = "create-a-new-" + model.RestName + "-for-a-given-" + parentRestName
	}
	return op
}

func (c *converter) extractOperationGetByID(model *spec.Model) *openapi3.Operation {

	if model == nil || model.Get == nil {
		return nil
	}
	relationAction := model.Get

	respBodySchemaRef := openapi3.NewSchemaRef("#/components/schemas/"+model.RestName, nil)

	op := &openapi3.Operation{
		OperationID: fmt.Sprintf("get-%s-by-ID", model.RestName),
		Tags:        []string{model.Group, model.Package},
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

func (c *converter) extractOperationDeleteByID(model *spec.Model) *openapi3.Operation {

	if model == nil || model.Delete == nil {
		return nil
	}
	relationAction := model.Delete

	respBodySchemaRef := openapi3.NewSchemaRef("#/components/schemas/"+model.RestName, nil)

	op := &openapi3.Operation{
		OperationID: fmt.Sprintf("delete-%s-by-ID", model.RestName),
		Tags:        []string{model.Group, model.Package},
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

func (c *converter) extractOperationPutByID(model *spec.Model) *openapi3.Operation {

	if model == nil || model.Update == nil {
		return nil
	}
	relationAction := model.Update

	schemaRef := openapi3.NewSchemaRef("#/components/schemas/"+model.RestName, nil)

	op := &openapi3.Operation{
		OperationID: fmt.Sprintf("update-%s-by-ID", model.RestName),
		Tags:        []string{model.Group, model.Package},
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
