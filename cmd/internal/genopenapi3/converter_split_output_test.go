package genopenapi3

import "testing"

func TestConverter_Do__splitOutput_emptyRootModel(t *testing.T) {
	t.Parallel()

	inSpec := `
		model:
			root: true
			rest_name: root
			resource_name: root
			entity_name: Root
			package: root
			group: core
			description: root object.
	`
	outDocs := `
		{
			"components": {},
			"paths": {}
		}
	`

	testCaseWrapper := map[string]testCase{
		"empty-toplevel-single-output": {
			inSplitOutput: true,
			inSpec:        inSpec,
			outDocs:       map[string]string{"toplevel": outDocs},
		},
	}
	runAllTestCases(t, testCaseWrapper)
}

func TestConverter_Do__split_output_complex(t *testing.T) {
	t.Parallel()

	inSpec := `
		model:
			root: true
			rest_name: root
			resource_name: root
			entity_name: Root
			package: root
			group: core
			description: root object.

		relations:
		- rest_name: minesite
			get:
				description: Retrieves all minesites.
			create:
				description: Creates a new minesite.
	`

	supportingSpecs := []string{
		`
		model:
			rest_name: minesite
			resource_name: minesites
			entity_name: MineSites
			package: usefulPackageName
			group: useful/thing
			description: Represents a resource mine site.
			get:
				description: Retrieves a mine site by ID.
			update:
				description: Updates a mine site by ID.
			delete:
				description: Delete a minesite by ID.

		relations:
		- rest_name: resource
			get:
				description: Retrieves a list of resources for a given mine site.
			create:
				description: assign a new resource for a given mine site.
		`,
		`
		model:
			rest_name: resource
			resource_name: resources
			entity_name: Resources
			package: naturalResources
			group: oil/gas
			description: Represents a natural resource.
		attributes:
			v1:
			- name: supervisor
				description: The supervisor of this natural resource.
				exposed: true
				type: ref
				subtype: employee
		`,
		`
		model:
			rest_name: employee
			resource_name: employees
			entity_name: Employees
			package: people
			group: employee/affairs
			description: Represents a full-time employee.
		`,
	}

	outDocs := map[string]string{
		"minesite": `
			{
				"components": {
					"schemas": {
						"minesite": {
							"type": "object"
						}
					}
				},
				"paths": {
					"/minesites": {
						"get": {
							"description": "Retrieves all minesites.",
							"operationId": "get-all-minesites",
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"items": {
													"$ref": "#/components/schemas/minesite"
												},
												"type": "array"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing",
								"usefulPackageName"
							]
						},
						"post": {
							"description": "Creates a new minesite.",
							"operationId": "create-a-new-minesite",
							"requestBody": {
								"content": {
									"application/json": {
										"schema": {
											"$ref": "#/components/schemas/minesite"
										}
									}
								}
							},
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "#/components/schemas/minesite"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing",
								"usefulPackageName"
							]
						}
					},
					"/minesites/{id}": {
						"delete": {
							"description": "Delete a minesite by ID.",
							"operationId": "delete-minesite-by-ID",
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "#/components/schemas/minesite"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing",
								"usefulPackageName"
							]
						},
						"get": {
							"description": "Retrieves a mine site by ID.",
							"operationId": "get-minesite-by-ID",
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "#/components/schemas/minesite"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing",
								"usefulPackageName"
							]
						},
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
						"put": {
							"description": "Updates a mine site by ID.",
							"operationId": "update-minesite-by-ID",
							"requestBody": {
								"content": {
									"application/json": {
										"schema": {
											"$ref": "#/components/schemas/minesite"
										}
									}
								}
							},
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "#/components/schemas/minesite"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"useful/thing",
								"usefulPackageName"
							]
						}
					},
					"/minesites/{id}/resources": {
						"get": {
							"description": "Retrieves a list of resources for a given mine site.",
							"operationId": "get-all-resources-for-a-given-minesite",
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"items": {
													"$ref": "./resource#/components/schemas/resource"
												},
												"type": "array"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"oil/gas",
								"naturalResources"
							]
						},
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
						"post": {
							"description": "assign a new resource for a given mine site.",
							"operationId": "create-a-new-resource-for-a-given-minesite",
							"requestBody": {
								"content": {
									"application/json": {
										"schema": {
											"$ref": "./resource#/components/schemas/resource"
										}
									}
								}
							},
							"responses": {
								"200": {
									"content": {
										"application/json": {
											"schema": {
												"$ref": "./resource#/components/schemas/resource"
											}
										}
									},
									"description": "n/a"
								}
							},
							"tags": [
								"oil/gas",
								"naturalResources"
							]
						}
					}
				}
			}
		`,

		"resource": `
			{
				"components": {
					"schemas": {
						"resource": {
							"properties": {
								"supervisor": {
									"$ref": "./employee#/components/schemas/employee"
								}
							},
							"type": "object"
						}
					}
				},
				"paths": {}
			}
		`,

		"employee": `
			{
				"components": {
					"schemas": {
						"employee": {
							"type": "object"
						}
					}
				},
				"paths": {}
			}
		`,
	}

	testCaseWrapper := map[string]testCase{
		"multiple-models-and-relations": {
			inSplitOutput:   true,
			inSpec:          inSpec,
			supportingSpecs: supportingSpecs,
			outDocs:         outDocs,
		},
	}
	runAllTestCases(t, testCaseWrapper)
}
