package genopenapi3

import (
	"testing"
)

func TestConverter_Do__model_root(t *testing.T) {
	t.Parallel()

	cases := map[string]testCase{

		"should-be-ignored": {
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
						"components": {},
						"paths": {}
					}
				`,
			},
		},
	}
	runAllTestCases(t, cases)
}
