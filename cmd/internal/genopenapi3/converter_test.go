package genopenapi3

import (
	"os"
	"testing"
)

func TestConverter_Do__schemas_and_attributes(t *testing.T) {

	cases := map[string]testCase{

		//
		"model-with-no-attributes": {
			inSpec: `
				model:
					rest_name: void
					resource_name: voids
					entity_name: Void
					package: None
					group: N/A
					description: empty model.
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"void": {
								"type": "object"
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		"model-with-unexposed-attribute--attr-ignored": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: somefield
						description: useful description.
						type: integer
						exposed: false
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object"
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		"model-is-ignored-if-private-and-skip-flag-is-set": {
			inSkipPrivateModels: true,
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
					private: true
				attributes:
					v1:
					- name: somefield
						description: useful description.
						type: integer
						exposed: true
			`,
			outDoc: `
				{
					"components": {},
					"paths": {}
				}
			`,
		},

		//
		"model-with-primitive-attributes": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: stringField
						description: useful description for string.
						type: string
						exposed: true
					- name: intField
						description: useful description for integer.
						type: integer
						exposed: true
					- name: floatField
						description: useful description for float.
						type: float
						exposed: true
					- name: booleanField
						description: useful description for boolean.
						type: boolean
						exposed: true
					- name: timeField
						description: useful description for time.
						type: time
						exposed: true
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object",
								"properties": {
									"stringField": {
										"description": "useful description for string.",
										"type": "string"
									},
									"intField": {
										"description": "useful description for integer.",
										"type": "integer"
									},
									"floatField": {
										"description": "useful description for float.",
										"type": "number"
									},
									"booleanField": {
										"description": "useful description for boolean.",
										"type": "boolean"
									},
									"timeField": {
										"description": "useful description for time.",
										"type": "string",
										"format": "date-time"
									}
								}
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		"model-with-enum-attribute": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField
						description: useful description.
						type: enum
						allowed_choices:
							- Choice1
							- Choice2
						exposed: true
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object",
								"properties": {
									"someField": {
										"description": "useful description.",
										"enum": ["Choice1", "Choice2"]
									}
								}
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		"model-with-object-attribute": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField
						description: useful description.
						type: object
						exposed: true
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object",
								"properties": {
									"someField": {
										"description": "useful description.",
										"type": "object"
									}
								}
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		"model-with-list-of-primitive-attributes": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: stringListField
						description: useful stringListField description.
						type: list
						subtype: string
						exposed: true
					- name: integerListField
						description: useful integerListField description.
						type: list
						subtype: integer
						exposed: true
					- name: floatListField
						description: useful floatListField description.
						type: list
						subtype: float
						exposed: true
					- name: booleanListField
						description: useful booleanListField description.
						type: list
						subtype: boolean
						exposed: true
					- name: timeListField
						description: useful timeListField description.
						type: list
						subtype: time
						exposed: true
				`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object",
								"properties": {
									"stringListField": {
										"description": "useful stringListField description.",
										"type": "array",
										"items": {
											"type": "string"
										}
									},
									"integerListField": {
										"description": "useful integerListField description.",
										"type": "array",
										"items": {
											"type": "integer"
										}
									},
									"floatListField": {
										"description": "useful floatListField description.",
										"type": "array",
										"items": {
											"type": "number"
										}
									},
									"booleanListField": {
										"description": "useful booleanListField description.",
										"type": "array",
										"items": {
											"type": "boolean"
										}
									},
									"timeListField": {
										"description": "useful timeListField description.",
										"type": "array",
										"items": {
											"type": "string",
											"format": "date-time"
										}
									}
								}
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		// we assume any referenced type is already defined in 'components.schemas'
		"model-with-ref-attribute": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField
						description: this should be ignored per openapi3 specs.
						type: ref
						subtype: imaginary
						exposed: true
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object",
								"properties": {
									"someField": {
										"$ref": "#/components/schemas/imaginary"
									}
								}
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		// we assume any referenced type is already defined in 'components.schemas'
		"model-with-refList-attributes": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField1
						description: useful someField1 description.
						type: refList
						subtype: imaginary1
						exposed: true
					- name: someField2
						description: useful someField2 description.
						type: refList
						subtype: imaginary2
						exposed: true
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object",
								"properties": {
									"someField1": {
										"description": "useful someField1 description.",
										"type": "array",
										"items": {
											"$ref": "#/components/schemas/imaginary1"
										}
									},
									"someField2": {
										"description": "useful someField2 description.",
										"type": "array",
										"items": {
											"$ref": "#/components/schemas/imaginary2"
										}
									}
								}
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		// we assume any referenced type is already defined in 'components.schemas'
		"model-with-refMap-attributes": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField1
						description: useful someField1 description.
						type: refMap
						subtype: imaginary1
						exposed: true
					- name: someField2
						description: useful someField2 description.
						type: refMap
						subtype: imaginary2
						exposed: true
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object",
								"properties": {
									"someField1": {
										"description": "useful someField1 description.",
										"type": "object",
										"additionalProperties": {
											"$ref": "#/components/schemas/imaginary1"
										}
									},
									"someField2": {
										"description": "useful someField2 description.",
										"type": "object",
										"additionalProperties": {
											"$ref": "#/components/schemas/imaginary2"
										}
									}
								}
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		"model-with-externalType-attributes--[]byte-turns-into-string": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: someField
						description: useful description.
						type: external
						subtype: '[]byte'
						exposed: true
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object",
								"properties": {
									"someField": {
										"description": "useful description.",
										"type": "string"
									}
								}
							}
						}
					},
					"paths": {}
				}
			`,
		},

		//
		"model-with-required-attributes": {
			inSpec: `
				model:
					rest_name: test
					resource_name: tests
					entity_name: Test
					package: None
					group: N/A
					description: dummy.
				attributes:
					v1:
					- name: stringField
						description: useful description for string.
						type: string
						exposed: true
						required: true
						default_value: hello-world
					- name: intField
						description: useful description for integer.
						type: integer
						exposed: true
			`,
			outDoc: `
				{
					"components": {
						"schemas": {
							"test": {
								"type": "object",
								"required": ["stringField"],
								"properties": {
									"stringField": {
										"description": "useful description for string.",
										"default": "hello-world",
										"type": "string"
									},
									"intField": {
										"description": "useful description for integer.",
										"type": "integer"
									}
								}
							}
						}
					},
					"paths": {}
				}
			`,
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
