package main

import (
	"bytes"
	"fmt"
	"go/format"
	"go/scanner"
	"path"
	"strings"
	"text/template"

	"github.com/aporeto-inc/regolithe/spec"
)

var functions = template.FuncMap{
	"upper":        strings.ToUpper,
	"lower":        strings.ToLower,
	"capitalize":   strings.Title,
	"join":         strings.Join,
	"makeAttr":     attrToField,
	"escBackticks": escapeBackticks,
	"buildEnums":   buildEnums,
}

func writeModel(set spec.SpecificationSet, name string, outFolder string) error {

	tmpl, err := makeTemplate("templates/model.gotpl")
	if err != nil {
		return err
	}

	s := set.Specification(name)

	// // Build enums
	// var enums []Enum
	// for _, attr := range s.Attributes(s.LatestAttributeVersion()) {

	// 	if attr.Type == spec.AttributeTypeEnum {
	// 		enums = append(enums, buildEnum(s.Model().EntityName, attr))
	// 	}
	// }

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Set  spec.SpecificationSet
			Spec spec.Specification
		}{
			Set:  set,
			Spec: s,
		}); err != nil {
		return fmt.Errorf("Unable to generate model '%s': %s", name, err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		if errs, ok := err.(scanner.ErrorList); ok {
			lines := strings.Split(buf.String(), "\n")
			for i := 0; i < errs.Len(); i++ {
				fmt.Printf("Error in '%s' near:\n\n\t%s\n\n", name, lines[errs[i].Pos.Line-1])
			}
		}
		return fmt.Errorf("Unable to format model '%s': %s", name, err)
	}

	if err := writeFile(path.Join(outFolder, name+".go"), p); err != nil {
		return fmt.Errorf("Unable to write file for spec: %s", name)
	}

	return nil
}

func writeIdentitiesRegistry(set spec.SpecificationSet, outFolder string) error {

	tmpl, err := makeTemplate("templates/identities_registry.gotpl")
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Set spec.SpecificationSet
		}{
			Set: set,
		}); err != nil {
		return fmt.Errorf("Unable to generate identities_registry code:%s", err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("Unable to format identities_registry code:%s", err)
	}

	if err := writeFile(path.Join(outFolder, "identities_registry.go"), p); err != nil {
		return fmt.Errorf("Unable to write file for identities_registry: %s", err)
	}

	return nil
}

func writeRelationshipsRegistry(set spec.SpecificationSet, outFolder string) error {

	tmpl, err := makeTemplate("templates/relationships_registry.gotpl")
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Set spec.SpecificationSet
		}{
			Set: set,
		}); err != nil {
		return fmt.Errorf("Unable to generate relationships_registry code:%s", err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("Unable to format relationships_registry code:%s", err)
	}

	if err := writeFile(path.Join(outFolder, "relationships_registry.go"), p); err != nil {
		return fmt.Errorf("Unable to write file for relationships_registry: %s", err)
	}

	return nil
}
