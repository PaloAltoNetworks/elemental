package genopenapi3

import "testing"

func TestConverter_Do__splitOutput_emptyRootModel(t *testing.T) {
	t.Parallel()

	inSpec := `
		model:
			root: true
			rest_name: root
			resource_name: root
			entity_name: Root
			package: root
			group: core
			description: root object.
	`
	outDocs := `
		{
			"components": {},
			"paths": {}
		}
	`

	testCaseWrapper := map[string]testCase{
		"empty-toplevel-single-output": {
			inSplitOutput: true,
			inSpec:        inSpec,
			outDocs:       map[string]string{"toplevel": outDocs},
		},
	}
	runAllTestCases(t, testCaseWrapper)
}

