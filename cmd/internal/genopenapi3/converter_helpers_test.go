package genopenapi3

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-test/deep"
	"go.aporeto.io/regolithe/spec"
	"gopkg.in/yaml.v2"
)

type testCase struct {
	inSpec              string
	inSkipPrivateModels bool
	outDocs             map[string]string // docname -> rawDoc excluding root keys 'openapi3' and 'info'
	supportingSpecs     []string          // other dependency specs needed for test case(s)
}

type testCaseRunner struct {
	t          *testing.T
	rootTmpDir string
}

func runAllTestCases(t *testing.T, cases map[string]testCase) {

	rootTmpDir, err := ioutil.TempDir("", t.Name()+"_*")
	if err != nil {
		t.Fatalf("error creating temporary directory for test function: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(rootTmpDir); err != nil {
			// no need to fail the test; it is just a temporary dir that
			// the OS will eventually destroy, but let's log the error
			t.Logf("error removing temporary test directory: %v", err)
		}
	})

	tcRunner := &testCaseRunner{
		t:          t,
		rootTmpDir: rootTmpDir,
	}
	for name, testCase := range cases {
		tcRunner.run(name, testCase)
	}
}

// Run will execute the given testcase in parallel with any other test cases
func (r *testCaseRunner) run(name string, tc testCase) {
	r.t.Run(name, func(t *testing.T) {
		t.Parallel()

		testDataFiles := map[string]string{
			// these files are needed by regolithe to parse the raw model from the test case
			"regolithe.ini": regolitheINI,
			"_type.mapping": typeMapping,
		}

		for i, rawSpec := range append([]string{tc.inSpec}, tc.supportingSpecs...) {
			rawSpec = replaceTrailingTabsWithDoubleSpaceForYAML(rawSpec)

			// this is needed because the spec filename has to match the rest_name of the spec model
			var spec struct {
				Model struct {
					RESTName string `yaml:"rest_name"`
				}
			}
			if err := yaml.Unmarshal([]byte(rawSpec), &spec); err != nil {
				t.Fatalf("error unmarshaling test spec data [%d] to read key 'rest_name': %v", i, err)
			}
			testDataFiles[spec.Model.RESTName+".spec"] = rawSpec
		}

		// this is to ensure that each test case executed within this runner is isolated
		specDir, err := ioutil.TempDir(r.rootTmpDir, name)
		if err != nil {
			t.Fatalf("error creating temporary directory for test case: %v", err)
		}

		for filename, content := range testDataFiles {
			filename = filepath.Join(specDir, filename)
			if err := ioutil.WriteFile(filename, []byte(content), os.ModePerm); err != nil {
				t.Fatalf("error writing temporary file '%s': %v", filename, err)
			}
		}

		spec, err := spec.LoadSpecificationSet(specDir, nil, nil, "openapi3")
		if err != nil {
			t.Fatalf("error parsing spec set from test data: %v", err)
		}

		cfg := Config{
			Public: tc.inSkipPrivateModels,
		}
		converter := newConverter(spec, cfg)

		output := map[string]*writeCloserMem{}
		writerFactory := func(docName string) (io.WriteCloser, error) {
			wr := &writeCloserMem{new(bytes.Buffer)}
			output[docName] = wr
			return wr, nil
		}
		if err := converter.Do(writerFactory); err != nil {
			t.Fatalf("error converting spec to openapi3: %v", err)
		}

		if la, le := len(output), len(tc.outDocs); la != le {
			t.Fatalf("expected %d output documents, got: %d", le, la)
		}
		for docName := range tc.outDocs {
			if _, ok := output[docName]; !ok {
				t.Fatalf("document with name '%s' does not exist in the actual output", docName)
			}
		}

		for expectedDocName, expectedRawDoc := range tc.outDocs {
			actualRawDoc := output[expectedDocName]
			actual := make(map[string]interface{})
			if err := json.Unmarshal(actualRawDoc.Bytes(), &actual); err != nil {
				t.Fatalf("invalid actual output data: malformed json content: %v", err)
			}

			expected := make(map[string]interface{})
			if err := json.Unmarshal([]byte(expectedRawDoc), &expected); err != nil {
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
		}
	})
}

func replaceTrailingTabsWithDoubleSpaceForYAML(s string) string {

	sb := new(strings.Builder)

	replaceNextTab := true
	for _, r := range s {

		if r == '\n' {
			sb.WriteRune(r)
			replaceNextTab = true
			continue
		}

		if replaceNextTab && r == '\t' {
			sb.WriteString("  ")
			continue
		}

		sb.WriteRune(r)
		replaceNextTab = false
	}

	return sb.String()
}

type fakeWriter struct {
	wrErr error
	cErr  error
}

func (fw *fakeWriter) Write([]byte) (int, error) {
	return 0, fw.wrErr
}

func (fw *fakeWriter) Close() error {
	return fw.cErr
}

type writeCloserMem struct {
	*bytes.Buffer
}

func (wr *writeCloserMem) Close() error {
	return nil
}
