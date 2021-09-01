package genopenapi3

import "testing"

func TestConverter_Do__relations_nonroot_spec(t *testing.T) {
	cases := map[string]testCase{
		"model-get-by-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: None
					group: N/A
					description: useful description.
					get:
						description: Retrieves the resource with the given ID.
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"resource": {
								"type": "object"
							}
						}
					},
					"paths": {
						"/resources/{id}": {
							"parameters": [
								{
									"in": "path",
									"name": "id",
									"required": true,
									"schema": {
										"type": "string"
									}
								}
							],
							"get": {
								"description": "Retrieves the resource with the given ID.",
								"responses": {
									"200": {
										"description": "n/a",
										"content": {
											"application/json": {
												"schema": {
													"$ref": "#/components/schemas/resource"
												}
											}
										}
									}
								}
							}
						}
					}
				}
			`,
		},
	}
	runAllTestCases(t, cases)
}
