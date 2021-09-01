package genopenapi3

import (
	"testing"
)

func TestConverter_Do__model_root(t *testing.T) {
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
	}
	runAllTestCases(t, cases)
}
