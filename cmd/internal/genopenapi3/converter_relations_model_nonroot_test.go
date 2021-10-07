package genopenapi3

import "testing"

func TestConverter_Do__modelRelations_nonRoot(t *testing.T) {
	t.Parallel()

	cases := map[string]testCase{

		"get-by-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: useful description.
					get:
						description: Retrieves the resource with the given ID.
						parameters:
							entries:
							- name: duplicateParam
								description: This is a fancy parameter that should appear only once.
								type: time
							- name: duplicateParam
								description: This is a fancy parameter that should appear only once.
								type: time
			`,
			outDocs: map[string]string{
				"toplevel": `
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
						"components": {
							"schemas": {
								"resource": {
									"description": "useful description.",
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
									"operationId": "get-resource-by-ID",
									"tags": ["useful/thing", "usefulPackageName"],
									"description": "Retrieves the resource with the given ID.",
									"parameters": [
										{
											"description": "This is a fancy parameter that should appear only once.",
											"in": "query",
											"name": "duplicateParam",
											"schema": {
												"type": "string",
												"format": "date-time"
											}
										}
									],
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
		},

		"delete-by-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: useful description.
					delete:
						description: Deletes the resource with the given ID.
						parameters:
							entries:
							- name: fancyParam
								description: This is a fancy parameter.
								type: duration
			`,
			outDocs: map[string]string{
				"toplevel": `
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
						"components": {
							"schemas": {
								"resource": {
									"description": "useful description.",
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
								"delete": {
									"operationId": "delete-resource-by-ID",
									"tags": ["useful/thing", "usefulPackageName"],
									"description": "Deletes the resource with the given ID.",
									"parameters": [
										{
											"description": "This is a fancy parameter.",
											"in": "query",
											"name": "fancyParam",
											"schema": {
												"type": "string"
											}
										}
									],
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
		},

		"put-by-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: useful description.
					update:
						description: Updates the resource with the given ID.
						parameters:
							entries:
							- name: fancyParam
								description: This is a fancy parameter.
								type: enum
								allowed_choices:
								- Choice1
								- Choice2
			`,
			outDocs: map[string]string{
				"toplevel": `
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
						"components": {
							"schemas": {
								"resource": {
									"description": "useful description.",
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
								"put": {
									"operationId": "update-resource-by-ID",
									"tags": ["useful/thing", "usefulPackageName"],
									"description": "Updates the resource with the given ID.",
									"parameters": [
										{
											"description": "This is a fancy parameter.",
											"in": "query",
											"name": "fancyParam",
											"schema": {
												"enum": ["Choice1", "Choice2"]
											}
										}
									],
									"requestBody": {
										"content": {
											"application/json": {
												"schema": {
													"$ref": "#/components/schemas/resource"
												}
											}
										}
									},
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
		},

		"get-put-delete-by-ID--do-not-duplicate-param-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: usefulPackageName
					group: useful/thing
					description: useful description.
					get:
						description: Retrieves the resource with the given ID.
					delete:
						description: Deletes the resource with the given ID.
					update:
						description: Updates the resource with the given ID.
			`,
			outDocs: map[string]string{
				"toplevel": `
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
						"components": {
							"schemas": {
								"resource": {
									"description": "useful description.",
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
									"operationId": "get-resource-by-ID",
									"tags": ["useful/thing", "usefulPackageName"],
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
								},
								"delete": {
									"operationId": "delete-resource-by-ID",
									"tags": ["useful/thing", "usefulPackageName"],
									"description": "Deletes the resource with the given ID.",
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
								},
								"put": {
									"operationId": "update-resource-by-ID",
									"tags": ["useful/thing", "usefulPackageName"],
									"description": "Updates the resource with the given ID.",
									"requestBody": {
										"content": {
											"application/json": {
												"schema": {
													"$ref": "#/components/schemas/resource"
												}
											}
										}
									},
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
		},
	}
	runAllTestCases(t, cases)
}
