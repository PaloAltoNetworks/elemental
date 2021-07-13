package main

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"go.aporeto.io/regolithe/spec"
)

func (sc *openapi3Converter) convertRelationsForRootSpec(relations []*spec.Relation) map[string]*openapi3.PathItem {

	paths := make(map[string]*openapi3.PathItem)

	for _, relation := range relations {

		if relation.Get == nil && relation.Create == nil {
			continue
		}

		pathItem := &openapi3.PathItem{
			Get:  sc.convertRelationActionToGetAll(relation.Get, relation.RestName),
			Post: sc.convertRelationActionToPost(relation.Create, relation.RestName),
		}

		relatedResourceName := sc.inSpecSet.Specification(relation.RestName).Model().ResourceName
		uri := fmt.Sprintf("/%s", relatedResourceName)
		paths[uri] = pathItem
	}

	return paths
}

func (sc *openapi3Converter) convertRelationsForNonRootSpec(resourceName string, relations []*spec.Relation) map[string]*openapi3.PathItem {

	paths := make(map[string]*openapi3.PathItem)

	for _, relation := range relations {

		if relation.Get == nil && relation.Create == nil {
			continue
		}

		pathItem := &openapi3.PathItem{
			Get:  sc.convertRelationActionToGetAll(relation.Get, relation.RestName),
			Post: sc.convertRelationActionToPost(relation.Create, relation.RestName),
		}

		sc.insertParamID(&pathItem.Parameters)
		relatedResourceName := sc.inSpecSet.Specification(relation.RestName).Model().ResourceName
		uri := fmt.Sprintf("/%s/{%s}/%s", resourceName, paramNameID, relatedResourceName)
		paths[uri] = pathItem
	}

	return paths
}

func (sc *openapi3Converter) convertRelationsForNonRootModel(model *spec.Model) map[string]*openapi3.PathItem {

	if model.Get == nil && model.Update == nil && model.Delete == nil {
		return nil
	}

	pathItem := &openapi3.PathItem{
		Get:    sc.convertRelationActionToGetByID(model.Get, model.RestName),
		Delete: sc.convertRelationActionToDeleteByID(model.Delete, model.RestName),
		Put:    sc.convertRelationActionToPutByID(model.Update, model.RestName),
	}
	sc.insertParamID(&pathItem.Parameters)

	uri := fmt.Sprintf("/%s/{%s}", model.ResourceName, paramNameID)
	pathItems := map[string]*openapi3.PathItem{uri: pathItem}
	return pathItems
}

func (sc *openapi3Converter) convertRelationActionToGetAll(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

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
		},
		Parameters: sc.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}

func (sc *openapi3Converter) convertRelationActionToPost(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

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
		},
		Parameters: sc.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}

func (sc *openapi3Converter) convertRelationActionToGetByID(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

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
		},
		Parameters: sc.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}

func (sc *openapi3Converter) convertRelationActionToDeleteByID(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

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
		},
		Parameters: sc.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}

func (sc *openapi3Converter) convertRelationActionToPutByID(relationAction *spec.RelationAction, restName string) *openapi3.Operation {

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
		},
		Parameters: sc.convertParamDefAsQueryParams(relationAction.ParameterDefinition),
	}

	return op
}
