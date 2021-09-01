package genopenapi3

import "testing"

func TestConverter_Do__modelRelations_nonRoot(t *testing.T) {
	cases := map[string]testCase{

		"get-by-ID": {
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

		"delete-by-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: None
					group: N/A
					description: useful description.
					delete:
						description: Deletes the resource with the given ID.
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
							"delete": {
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
							}
						}
					}
				}
			`,
		},

		"put-by-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: None
					group: N/A
					description: useful description.
					update:
						description: Updates the resource with the given ID.
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
							"put": {
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

		"get-put-delete-by-ID--do-not-duplicate-param-ID": {
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
					delete:
						description: Deletes the resource with the given ID.
					update:
						description: Updates the resource with the given ID.
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
							},
							"delete": {
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
	}
	runAllTestCases(t, cases)
}
