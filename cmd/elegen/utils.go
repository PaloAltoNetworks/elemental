package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/aporeto-inc/elemental/cmd/elegen/static"
	"github.com/aporeto-inc/regolithe/spec"
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

func attrToField(attr *spec.Attribute) string {

	json := attr.Name
	bson := strings.ToLower(attr.Name)

	if !attr.Exposed {
		json = "-"
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

func buildEnum(entityName string, attr *spec.Attribute) Enum {

	attr.ConvertedType = fmt.Sprintf("%s%sValue", entityName, attr.ConvertedName)

	values := map[string]string{}
	for _, v := range attr.AllowedChoices {
		k := fmt.Sprintf("%s%s%s", entityName, attr.ConvertedName, strings.Title(strings.ToLower(v)))
		values[k] = v
	}

	return Enum{
		Type:          attr.ConvertedType,
		Values:        values,
		AttributeName: attr.Name, // TODO: put converted name
	}
}
