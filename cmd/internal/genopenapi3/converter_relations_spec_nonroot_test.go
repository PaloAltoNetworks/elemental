package genopenapi3

import "testing"

func TestConverter_Do__specRelations_nonRoot(t *testing.T) {
	t.Parallel()

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
						parameters:
							entries:
							- name: fancyParam
								description: This is a fancy parameter.
								type: string
								default_value: "this is a value"
							- name: aParam
								description: should appear at the beginning.
								type: string
								default_value: "this is a value 2"
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"openapi": "3.0.3",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							},
							{
								"name": "resource",
								"description": "This tag is for package 'resource'"
							},
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
							"title": "toplevel"
						},
						"components": {
							"schemas": {
								"minesite": {
									"description": "Represents a resource mine site.",
									"type": "object"
								},
								"resource": {
									"description": "useful description.",
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
									"operationId": "get-all-minesites-for-a-given-resource",
									"tags": ["useful/thing", "usefulPackageName"],
									"description": "Retrieve all mine sites.",
									"parameters": [
										{
											"description": "should appear at the beginning.",
											"in": "query",
											"name": "aParam",
											"schema": {
												"type": "string"
											},
											"example": "this is a value 2"
										},
										{
											"description": "This is a fancy parameter.",
											"in": "query",
											"name": "fancyParam",
											"schema": {
												"type": "string"
											},
											"example": "this is a value"
										}
									],
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
			},
			supportingSpecs: []string{`
				model:
					rest_name: minesite
					resource_name: minesites
					entity_name: MineSites
					package: usefulPackageName
					group: useful/thing
					description: Represents a resource mine site.
			`},
		},

		"relation-create": {
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
						parameters:
							entries:
							- name: fancyParam
								description: This is a fancy parameter.
								type: float


			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"openapi": "3.0.3",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							},
							{
								"name": "resource",
								"description": "This tag is for package 'resource'"
							},
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
							"title": "toplevel"
						},
						"components": {
							"schemas": {
								"minesite": {
									"description": "Represents a resource mine site.",
									"type": "object"
								},
								"resource": {
									"description": "useful description.",
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
									"operationId": "create-a-new-minesite-for-a-given-resource",
									"tags": ["useful/thing", "usefulPackageName"],
									"description": "Creates a mine site.",
									"parameters": [
										{
											"description": "This is a fancy parameter.",
											"in": "query",
											"name": "fancyParam",
											"schema": {
												"type": "number"
											}
										}
									],
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
			},
			supportingSpecs: []string{`
				model:
					rest_name: minesite
					resource_name: minesites
					entity_name: MineSites
					package: usefulPackageName
					group: useful/thing
					description: Represents a resource mine site.
			`},
		},

		"relation-without-get-or-create": {
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
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"openapi": "3.0.3",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							},
							{
								"name": "None",
								"description": "This tag is for package 'None'"
							},
							{
								"name": "resource",
								"description": "This tag is for package 'resource'"
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
							"title": "toplevel"
						},
						"components": {
							"schemas": {
								"minesite": {
									"description": "Represents a resource mine site.",
									"type": "object"
								},
								"resource": {
									"description": "useful description.",
									"type": "object"
								}
							}
						},
						"paths": {}
					}
				`,
			},
			supportingSpecs: []string{`
				model:
					rest_name: minesite
					resource_name: minesites
					entity_name: MineSites
					package: None
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
			outDocs: map[string]string{
				"toplevel": `
					{
						"openapi": "3.0.3",
						"tags": [
							{
								"name": "N/A",
								"description": "This tag is for group 'N/A'"
							},
							{
								"name": "resource",
								"description": "This tag is for package 'resource'"
							},
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
							"title": "toplevel"
						},
						"components": {
							"schemas": {
								"minesite": {
									"description": "Represents a resource mine site.",
									"type": "object"
								},
								"resource": {
									"description": "useful description.",
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
									"operationId": "create-a-new-minesite-for-a-given-resource",
									"tags": ["useful/thing", "usefulPackageName"],
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
									"operationId": "get-all-minesites-for-a-given-resource",
									"tags": ["useful/thing", "usefulPackageName"],
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
			},
			supportingSpecs: []string{`
				model:
					rest_name: minesite
					resource_name: minesites
					entity_name: MineSites
					package: usefulPackageName
					group: useful/thing
					description: Represents a resource mine site.
			`},
		},
	}
	runAllTestCases(t, cases)
}
