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
			"openapi": "3.0.3",
			"info": {
				"contact": {
					"email": "dev@aporeto.com",
					"name":  "Aporeto Inc.",
					"url":   "go.aporeto.io/api"
				},
				"license": {
					"name": "TODO"
				},
				"termsOfService": "https://localhost/TODO",
				"version": "1.0",
				"title": "toplevel"
			},
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
				"openapi": "3.0.3",
				"tags":[
					{
						"name": "useful/thing",
						"description": "This tag is for group 'useful/thing'"
					},
					{
						"name": "usefulPackageName",
						"description": "This tag is for package 'usefulPackageName'"
					}
				],
				"info": {
					"contact": {
						"email": "dev@aporeto.com",
						"name":  "Aporeto Inc.",
						"url":   "go.aporeto.io/api"
					},
					"license": {
						"name": "TODO"
					},
					"termsOfService": "https://localhost/TODO",
					"version": "1.0",
					"title": "minesite"
				},
				"components": {
					"schemas": {
						"minesite": {
							"description": "Represents a resource mine site.",
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
				"openapi": "3.0.3",
				"tags":[
					{
						"name": "oil/gas",
						"description": "This tag is for group 'oil/gas'"
					},
					{
						"name": "naturalResources",
						"description": "This tag is for package 'naturalResources'"
					}
				],
				"info": {
					"contact": {
						"email": "dev@aporeto.com",
						"name":  "Aporeto Inc.",
						"url":   "go.aporeto.io/api"
					},
					"license": {
						"name": "TODO"
					},
					"termsOfService": "https://localhost/TODO",
					"version": "1.0",
					"title": "resource"
				},
				"components": {
					"schemas": {
						"resource": {
							"description": "Represents a natural resource.",
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
				"openapi": "3.0.3",
				"tags":[
					{
						"name": "employee/affairs",
						"description": "This tag is for group 'employee/affairs'"
					},
					{
						"name": "people",
						"description": "This tag is for package 'people'"
					}
				],
				"info": {
					"contact": {
						"email": "dev@aporeto.com",
						"name":  "Aporeto Inc.",
						"url":   "go.aporeto.io/api"
					},
					"license": {
						"name": "TODO"
					},
					"termsOfService": "https://localhost/TODO",
					"version": "1.0",
					"title": "employee"
				},
				"components": {
					"schemas": {
						"employee": {
							"description": "Represents a full-time employee.",
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

func TestConverter_Do__split_output_withPrivateModel(t *testing.T) {
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
		- rest_name: hidden
			get:
				description: Retrieves all hidden secrets.
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
		`,
		`
		model:
			rest_name: hidden
			resource_name: hiddens
			entity_name: Hiddens
			package: secret
			group: secret/affairs
			description: Represents a private model.
			private: true
		`,
	}

	outDocs := map[string]string{
		"minesite": `
			{
				"openapi": "3.0.3",
				"tags":[
					{
						"name": "useful/thing",
						"description": "This tag is for group 'useful/thing'"
					},
					{
						"name": "usefulPackageName",
						"description": "This tag is for package 'usefulPackageName'"
					}
				],
				"info": {
					"contact": {
						"email": "dev@aporeto.com",
						"name":  "Aporeto Inc.",
						"url":   "go.aporeto.io/api"
					},
					"license": {
						"name": "TODO"
					},
					"termsOfService": "https://localhost/TODO",
					"version": "1.0",
					"title": "minesite"
				},
				"components": {
					"schemas": {
						"minesite": {
							"description": "Represents a resource mine site.",
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
						}
					}
				}
			}
		`,
	}

	testCaseWrapper := map[string]testCase{
		"root-relation-has-private-model": {
			inSplitOutput:       true,
			inSkipPrivateModels: true,
			inSpec:              inSpec,
			supportingSpecs:     supportingSpecs,
			outDocs:             outDocs,
		},
	}
	runAllTestCases(t, testCaseWrapper)
}
