package genopenapi3

import "testing"

func TestConverter_Do__specRelations_nonRoot(t *testing.T) {
	cases := map[string]testCase{
		"relation-get": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: resource
					group: N/A
					description: useful description.

				relations:
				- rest_name: minesite
					get:
						description: Retrieve all mine sites.
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"minesite": {
								"type": "object"
							},
							"resource": {
								"type": "object"
							}
						}
					},
					"paths": {
						"/resources/{id}/minesites": {
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
								"description": "Retrieve all mine sites.",
								"responses": {
									"200": {
										"description": "n/a",
										"content": {
											"application/json": {
												"schema": {
													"type": "array",
													"items": {
														"$ref": "#/components/schemas/minesite"
													}
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
			supportingSpecs: []string{`
				model:
					rest_name: minesite
					resource_name: minesites
					entity_name: MineSites
					package: none
					group: N/A
					description: Represents a resource mine site.
			`},
		},

		"relation-post": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: resource
					group: N/A
					description: useful description.

				relations:
				- rest_name: minesite
					create:
						description: Creates a mine site.
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"minesite": {
								"type": "object"
							},
							"resource": {
								"type": "object"
							}
						}
					},
					"paths": {
						"/resources/{id}/minesites": {
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
								"description": "Creates a mine site.",
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
										"description": "n/a",
										"content": {
											"application/json": {
												"schema": {
													"$ref": "#/components/schemas/minesite"
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
			supportingSpecs: []string{`
				model:
					rest_name: minesite
					resource_name: minesites
					entity_name: MineSites
					package: none
					group: N/A
					description: Represents a resource mine site.
			`},
		},

		"relation-get-post--do-not-duplicate-param-ID": {
			inSpec: `
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Resource
					package: resource
					group: N/A
					description: useful description.

				relations:
				- rest_name: minesite
					get:
						description: Retrieve all mine sites.
					create:
						description: Creates a mine site.
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"minesite": {
								"type": "object"
							},
							"resource": {
								"type": "object"
							}
						}
					},
					"paths": {
						"/resources/{id}/minesites": {
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
								"description": "Creates a mine site.",
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
										"description": "n/a",
										"content": {
											"application/json": {
												"schema": {
													"$ref": "#/components/schemas/minesite"
												}
											}
										}
									}
								}
							},
							"get": {
								"description": "Retrieve all mine sites.",
								"responses": {
									"200": {
										"description": "n/a",
										"content": {
											"application/json": {
												"schema": {
													"type": "array",
													"items": {
														"$ref": "#/components/schemas/minesite"
													}
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
			supportingSpecs: []string{`
				model:
					rest_name: minesite
					resource_name: minesites
					entity_name: MineSites
					package: none
					group: N/A
					description: Represents a resource mine site.
			`},
		},
	}
	runAllTestCases(t, cases)
}
