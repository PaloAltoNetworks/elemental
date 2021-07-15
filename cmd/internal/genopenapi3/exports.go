package genopenapi3

import (
	"fmt"
	"os"
	"path/filepath"

	"go.aporeto.io/regolithe/spec"
)

// GeneratorFunc implements the signature defined by regolithe to convert spec to openapi3 doc
func GeneratorFunc(sets []spec.SpecificationSet, out string) error {

	file, err := os.Create(filepath.Join(out, "openapi3.json"))
	if err != nil {
		return fmt.Errorf("error creating 'openapi3.json' file: %w", err)
	}

	set := sets[0]
	converter := newConverter(set)
	if err = converter.Do(file); err != nil {
		return fmt.Errorf("error generating openapi3 document from spec set '%s': %w", set.Configuration().Name, err)
	}

	return nil
}
