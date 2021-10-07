package genopenapi3

import (
	"testing"
)

func TestConverter_Do__specRelations_root(t *testing.T) {
	t.Parallel()

	cases := map[string]testCase{

		"relation-create": {
			inSpec: `
				model:
					root: true
					rest_name: root
					resource_name: root
					entity_name: Root
					package: root
					group: core
					description: root object.

				relations:
				- rest_name: resource
					create:
						description: Creates some resource.
						parameters:
							entries:
							- name: fancyParam
								description: This is a fancy parameter.
								type: integer
			`,
			outDocs: map[string]string{
				"toplevel": `
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
							"title": "toplevel"
						},
						"components": {
							"schemas": {
								"resource": {
									"description": "Represents a resource.",
									"type": "object"
								}
							}
						},
						"paths": {
							"/resources": {
								"post": {
									"operationId": "create-a-new-resource",
									"tags": ["useful/thing", "usefulPackageName"],
									"parameters": [
										{
											"description": "This is a fancy parameter.",
											"in": "query",
											"name": "fancyParam",
											"schema": {
												"type": "integer"
											}
										}
									],
									"description": "Creates some resource.",
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
			supportingSpecs: []string{`
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Recource
					package: usefulPackageName
					group: useful/thing
					description: Represents a resource.
			`},
		},

		"relation-get": {
			inSpec: `
				model:
					root: true
					rest_name: root
					resource_name: root
					entity_name: Root
					package: root
					group: core
					description: root object.

				relations:
				- rest_name: resource
					get:
						description: Retrieve all resources.
						parameters:
						  entries:
						  - name: fancyParam
						    description: This is a fancy parameter.
						    type: boolean
			`,
			outDocs: map[string]string{
				"toplevel": `
					{
						"openapi": "3.0.3",
						"tags": [
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
								"resource": {
									"description": "Represents a resource.",
									"type": "object"
								}
							}
						},
						"paths": {
							"/resources": {
								"get": {
									"operationId": "get-all-resources",
									"tags": ["useful/thing", "usefulPackageName"],
									"description": "Retrieve all resources.",
									"parameters": [
									  {
									    "description": "This is a fancy parameter.",
									    "in": "query",
									    "name": "fancyParam",
									    "schema": {
									      "type": "boolean"
									    }
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
					}
				`,
			},
			supportingSpecs: []string{`
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Recource
					package: usefulPackageName
					group: useful/thing
					description: Represents a resource.
			`},
		},

		"relation-without-get-or-create": {
			inSpec: `
				model:
					root: true
					rest_name: root
					resource_name: root
					entity_name: Root
					package: root
					group: core
					description: root object.

				relations:
				- rest_name: resource
			`,
			outDocs: map[string]string{
				"toplevel": `
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
							"title": "toplevel"
						},
						"components": {
							"schemas": {
								"resource": {
									"description": "Represents a resource.",
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
					rest_name: resource
					resource_name: resources
					entity_name: Recource
					package: usefulPackageName
					group: useful/thing
					description: Represents a resource.
			`},
		},
	}
	runAllTestCases(t, cases)
}
