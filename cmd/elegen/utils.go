package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"go.aporeto.io/elemental/cmd/elegen/static"
	"go.aporeto.io/regolithe/spec"
)

// An Enum represents an enum.
type Enum struct {
	Type          string
	Values        map[string]string
	AttributeName string
}

func makeTemplate(p string) (*template.Template, error) {

	data, err := static.Asset(p)
	if err != nil {
		return nil, err
	}

	return template.New(path.Base(p)).Funcs(functions).Parse(string(data))
}

func writeFile(path string, data []byte) error {

	// fmt.Println(string(data))

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Unable to write file: %s", f.Name())
	}

	defer f.Close() // nolint: errcheck
	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("Unable to write file: %s", f.Name())
	}

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

func attrToField(attr *spec.Attribute, publicMode bool) string {

	json := attr.Name
	bson := strings.ToLower(attr.Name)

	if !attr.Exposed {
		json = "-"
	}

	if attr.OmitEmpty {
		json += ",omitempty"
	}

	if !attr.Stored {
		bson = "-"
	} else if attr.Identifier {
		bson = "_" + bson
	}

	descLines := strings.Split(attr.Description, "\n")
	for i := 0; i < len(descLines); i++ {
		descLines[i] = "// " + escapeBackticks(descLines[i])
	}

	return fmt.Sprintf(
		"%s\n    %s %s `json:\"%s\" bson:\"%s\" mapstructure:\"%s,omitempty\"`\n\n",
		strings.Join(descLines, "\n"),
		attr.ConvertedName,
		attr.ConvertedType,
		json,
		bson,
		json,
	)
}

func escapeBackticks(str string) string {
	return strings.Replace(str, "`", "`+\"`\"+`", -1)
}

func buildEnums(s spec.Specification, version string) []Enum {

	var enums []Enum
	attributes := s.Attributes(version)

	for _, attr := range attributes {

		if attr.Type != spec.AttributeTypeEnum {
			continue
		}

		attr.ConvertedType = fmt.Sprintf("%s%sValue", s.Model().EntityName, attr.ConvertedName)

		values := map[string]string{}
		for _, v := range attr.AllowedChoices {
			k := fmt.Sprintf("%s%s%s", s.Model().EntityName, attr.ConvertedName, v)
			values[k] = v
		}

		enums = append(enums, Enum{
			Type:          attr.ConvertedType,
			Values:        values,
			AttributeName: attr.Name,
		})
	}

	return enums
}

func shouldGenerateGetter(attr *spec.Attribute, publicMode bool) bool {

	if publicMode {
		return attr.Getter && attr.Exposed
	}

	return attr.Getter
}

func shouldGenerateSetter(attr *spec.Attribute, publicMode bool) bool {

	if publicMode {
		return attr.Setter && attr.Exposed
	}

	return attr.Setter
}

func shouldWriteInitializer(s spec.Specification, attrConvertedName string, version string, publicMode bool) bool {

	var attr *spec.Attribute
	for _, a := range s.Attributes(version) {
		if a.ConvertedName == attrConvertedName {
			attr = a
			break
		}
	}

	if publicMode {
		return attr.Exposed
	}

	return true
}

func shouldWriteAttributeMap(attr *spec.Attribute, publicMode bool) bool {

	if publicMode {
		return attr.Exposed
	}

	return true
}

func shouldRegisterSpecification(s spec.Specification, publicMode bool) bool {

	if publicMode {
		return !s.Model().Private
	}

	return true
}

func shouldRegisterRelationship(set spec.SpecificationSet, entityName string, publicMode bool) bool {

	var s spec.Specification
	for _, i := range set.Specifications() {
		if i.Model().EntityName == entityName {
			s = i
		}
	}

	if publicMode {
		return !s.Model().Private
	}

	return true
}

func shouldRegisterInnerRelationship(set spec.SpecificationSet, restName string, publicMode bool) bool {

	s := set.Specification(restName)

	if publicMode {
		return !s.Model().Private
	}

	return true
}
