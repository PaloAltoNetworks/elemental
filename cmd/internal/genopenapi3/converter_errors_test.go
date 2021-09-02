package genopenapi3

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"go.aporeto.io/regolithe/spec"
)

func TestConverter_Do__error_no_externalType_mapping(t *testing.T) {

	specDir, err := os.MkdirTemp("", t.Name()+"_*")
	if err != nil {
		t.Fatalf("error creating temporary directory for test function: %v", err)
	}
	t.Cleanup(func() { os.RemoveAll(specDir) })

	badTypeMapping := replaceTrailingTabsWithDoubleSpaceForYAML(`
		'[]byte':
			openapi3:
				type: malformed-json }
	`)

	rawSpec := replaceTrailingTabsWithDoubleSpaceForYAML(`
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
	`)

	for filename, content := range map[string]string{
		"regolithe.ini": regolitheINI,
		"_type.mapping": badTypeMapping,
		"test.spec":     rawSpec,
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

	converter := newConverter(spec, false)
	if err := converter.Do(nil); !errors.Is(err, errUnmarshalingExternalType) {
		t.Fatalf("unexpected error\nwant: %v\n got: %v", errUnmarshalingExternalType, err)
	}
}
