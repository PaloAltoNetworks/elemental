package genopenapi3

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"go.aporeto.io/regolithe/spec"
)

// GeneratorFunc implements the signature defined by regolithe to convert spec to openapi3 doc
func GeneratorFunc(sets []spec.SpecificationSet, out string) error {

	outFolder := path.Join(out, "openapi3")
	if err := os.MkdirAll(outFolder, 0750); err != nil && !os.IsExist(err) {
		return fmt.Errorf("'%s': error creating directory: %w", outFolder, err)
	}

	filename := filepath.Join(outFolder, "toplevel.json")
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("'%s': error creating file: %w", filename, err)
	}

	set := sets[0]
	converter := newConverter(set)
	if err = converter.Do(file); err != nil {
		return fmt.Errorf("error generating openapi3 document from spec set '%s': %w", set.Configuration().Name, err)
	}

	return nil
}
