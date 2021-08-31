package genopenapi3

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"go.aporeto.io/regolithe/spec"
	"gopkg.in/yaml.v2"
)

func TestConverter_Do(t *testing.T) {

	cases := map[string]struct {
		inSpec              string
		inSkipPrivateModels bool
		outDoc              string // excluding root keys 'openapi3' and 'info'
	}{
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
	}

	rootTmpDir, err := os.MkdirTemp("", t.Name()+"_*")
	if err != nil {
		t.Fatalf("error creating temporary directory for test function: %v", err)
	}
	defer os.RemoveAll(rootTmpDir)

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			c.inSpec = replaceTrailingTabsWithDoubleSpaceForYAML(c.inSpec)

			// this is to ensure that each test case is isolated
			specDir, err := os.MkdirTemp(rootTmpDir, name)
			if err != nil {
				t.Fatalf("error creating temporary directory for test case: %v", err)
			}

			// this is needed because the spec filename has to match the rest_name of the spec model
			var inSpecDeserialized struct {
				Model struct {
					RESTName string `yaml:"rest_name"`
				}
			}
			if err := yaml.Unmarshal([]byte(c.inSpec), &inSpecDeserialized); err != nil {
				t.Fatalf("error unmarshaling test spec data to read key 'rest_name': %v", err)
			}

			for filename, content := range map[string]string{
				// these files are needed by regolithe to parse the raw model from the test case
				"regolithe.ini": regolitheINI,
				"_type.mapping": typemapping,
				// this is what will be parsed by regolithe
				inSpecDeserialized.Model.RESTName + ".spec": c.inSpec,
			} {
				filename = filepath.Join(specDir, filename)
				if err := os.WriteFile(filename, []byte(content), os.ModePerm); err != nil {
					t.Fatalf("error writing temporary file '%s': %v", filename, err)
				}
			}

			spec, err := spec.LoadSpecificationSet(specDir, nil, nil, "openapi3")
			if err != nil {
				t.Fatalf("error parsing spec set from test data: %v", err)
			}

			converter := newConverter(spec, c.inSkipPrivateModels)
			output := new(bytes.Buffer)
			if err := converter.Do(output); err != nil {
				t.Fatalf("error converting spec to openapi3: %v", err)
			}

			actual := make(map[string]interface{})
			if err := json.Unmarshal(output.Bytes(), &actual); err != nil {
				t.Fatalf("invalid actual output data: malformed json content: %v", err)
			}

			expected := make(map[string]interface{})
			if err := json.Unmarshal([]byte(c.outDoc), &expected); err != nil {
				t.Fatalf("invalid expected output data in test case: malformed json content: %v", err)
			}
			// root keys 'openapi3' and 'info' must be identical for all test cases;
			// therefore, we inject them here or fail the test if they are defined
			// to make test cases more readable and to prevent repeating them in all
			// test cases, which is going to be boring
			if _, ok := expected["openapi"]; ok {
				t.Fatal("key 'openapi' must not be defined in the expected outDoc as it is set by the test")
			}
			if _, ok := expected["info"]; ok {
				t.Fatal("key 'info' must not be defined in the expected outDoc as it is set by the test")
			}
			expected["openapi"] = "3.0.3"
			expected["info"] = map[string]interface{}{
				"contact": map[string]interface{}{
					"email": "dev@aporeto.com",
					"name":  "Aporeto Inc.",
					"url":   "go.aporeto.io/api",
				},
				"license": map[string]interface{}{
					"name": "TODO",
				},
				"termsOfService": "https://localhost/TODO",
				"title":          "gaia",
				"version":        "1.0",
			}

			if diff := deep.Equal(actual, expected); diff != nil {
				t.Fatal("actual != expected output\n", strings.Join(diff, "\n"))
			}
		})
	}
}
