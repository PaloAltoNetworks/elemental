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
