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
						"components": {},
						"paths": {}
					}
				`,
			},
		},
	}
	runAllTestCases(t, cases)
}
