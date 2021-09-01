package genopenapi3

import (
	"os"
	"testing"
)

func TestConverter_Do__relations_root(t *testing.T) {

	cases := map[string]testCase{

		//
		"root-model-should-be-ignored": {
			inSpec: `
				model:
					root: true
					rest_name: root
					resource_name: root
					entity_name: Root
					package: root
					group: core
					description: root object.
			`,
			outDoc: `
				{
					"components": {},
					"paths": {}
				}
			`,
		},

		//
		"root-model-relation-create": {
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
						"/resources": {
							"post": {
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
			supportingSpecs: []string{`
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Recource
					package: none
					group: N/A
					description: Represents a resource.
			`},
		},

		//
		"root-model-relation-get": {
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
						"/resources": {
							"get": {
								"description": "Retrieve all resources.",
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
			supportingSpecs: []string{`
				model:
					rest_name: resource
					resource_name: resources
					entity_name: Recource
					package: none
					group: N/A
					description: Represents a resource.
			`},
		},
	}

	rootTmpDir, err := os.MkdirTemp("", t.Name()+"_*")
	if err != nil {
		t.Fatalf("error creating temporary directory for test function: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(rootTmpDir) })

	tcRunner := &testCaseRunner{
		t:          t,
		rootTmpDir: rootTmpDir,
	}
	for name, testCase := range cases {
		tcRunner.Run(name, testCase)
	}

}
