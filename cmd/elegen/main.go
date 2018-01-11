package main

import (
	"bytes"
	"fmt"
	"go/format"
	"strings"
	"text/template"

	"golang.org/x/sync/errgroup"

	"github.com/Sirupsen/logrus"

	"github.com/aporeto-inc/regolithe"
	"github.com/aporeto-inc/regolithe/spec"
)

// An Enum is cool.
type Enum struct {
	Type   string
	Values map[string]string
}

var functions = template.FuncMap{
	"upper":      strings.ToUpper,
	"lower":      strings.ToLower,
	"capitalize": strings.Title,
	"join":       strings.Join,
	"makeAttr":   attrToField,
}

func main() {

	if err := regolithe.NewCommand(
		"elegen",
		"elegen",
		attributeNameConverter,
		attributeTypeConverter,
		"elemental",
		generator,
	).Execute(); err != nil {
		logrus.WithError(err).Fatal("Error during generation")
	}
}

func generator(set *spec.SpecificationSet) error {

	var g errgroup.Group

	g.Go(func() error {
		if err := writeIdentitiesRegistry(set); err != nil {
			return err
		}
		return nil
	})

	for _, s := range set.Specifications() {
		g.Go(func() error {
			if err := writeModel(set, s.RestName); err != nil {
				return err
			}
			return nil
		})
	}

	return g.Wait()
}

func writeModel(set *spec.SpecificationSet, name string) error {

	tmpl, err := template.New("model.gotpl").
		Funcs(functions).
		ParseFiles("./templates/model.gotpl")
	if err != nil {
		return err
	}

	s := set.Specification(name)

	// Build enums
	var enums []Enum
	for _, attr := range s.Attributes {

		if attr.Type == spec.AttributeTypeEnum {
			attr.ConvertedType = fmt.Sprintf("%sTypeValue", s.EntityName)
			enums = append(enums, buildEnum(s.EntityName, attr))
		}

	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Set   *spec.SpecificationSet
			Spec  *spec.Specification
			Enums []Enum
		}{
			Set:   set,
			Spec:  s,
			Enums: enums,
		}); err != nil {
		return fmt.Errorf("Unable to generate data :%s", err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		// fmt.Println(buf.String())
		return fmt.Errorf("Unable to format data :%s", err)
	}

	fmt.Println(string(p))

	return nil
}

func writeIdentitiesRegistry(set *spec.SpecificationSet) error {

	tmpl, err := template.New("identities_registry.gotpl").
		Funcs(functions).
		ParseFiles("./templates/identities_registry.gotpl")
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err = tmpl.Execute(
		&buf,
		struct {
			Set *spec.SpecificationSet
		}{
			Set: set,
		}); err != nil {
		return fmt.Errorf("Unable to generate data :%s", err)
	}

	p, err := format.Source(buf.Bytes())
	if err != nil {
		// fmt.Println(buf.String())
		return fmt.Errorf("Unable to format data :%s", err)
	}

	fmt.Println(string(p))

	return nil
}

func attributeTypeConverter(typ spec.AttributeType, subtype string) (string, string) {

	switch typ {

	case spec.AttributeTypeString, spec.AttributeTypeEnum:
		return "string", ""

	case spec.AttributeTypeFloat:
		return "float64", ""

	case spec.AttributeTypeBool:
		return "bool", ""

	case spec.AttributeTypeInt:
		return "int", ""

	case spec.AttributeTypeTime:
		return "time.Time", "time"

	case spec.AttributeTypeList:
		if subtype == "" {
			return "[]interface{}", ""
		}
		return "[]" + subtype, ""

	default:
		return "interface{}", ""
	}
}

func attributeNameConverter(attrName string) string {

	return strings.Title(attrName)
}

func attrToField(attr *spec.Attribute) string {

	json := attr.Name
	bson := strings.ToLower(attr.Name)

	if !attr.Exposed {
		json = "-"
	}

	if !attr.Stored {
		bson = "-"
	}

	if attr.Identifier {
		bson = "_" + bson
	}

	return fmt.Sprintf(
		"    %s %s `json:\"%s\" bson:\"%s\"` // %s\n",
		attr.ConvertedName,
		attr.ConvertedType,
		json,
		bson,
		attr.Description,
	)
}

func buildEnum(entityName string, attr *spec.Attribute) Enum {

	values := map[string]string{}
	for _, v := range attr.AllowedChoices {
		k := fmt.Sprintf("%s%s%s", entityName, attr.ConvertedName, strings.Title(strings.ToLower(v)))
		values[k] = v
	}

	return Enum{
		Type:   attr.ConvertedType,
		Values: values,
	}
}
