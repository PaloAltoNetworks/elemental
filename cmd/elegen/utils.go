// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"sort"
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

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("unable to write file: %s", f.Name())
	}

	// #nosec G307
	defer f.Close() // nolint: errcheck
	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("unable to write file: %s", f.Name())
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
		// @TODO: pass subtype as a spec.AttributeType so we can make a recursive call here.
		if subtype == "" || subtype == "object" {
			return "[]interface{}", ""
		}
		if subtype == string(spec.AttributeTypeInt) {
			return "[]int", ""
		}
		if subtype == string(spec.AttributeTypeBool) {
			return "[]bool", ""
		}
		if subtype == string(spec.AttributeTypeFloat) {
			return "[]float64", ""
		}
		if subtype == string(spec.AttributeTypeTime) {
			return "[]time.Time", "time"
		}
		return "[]" + subtype, ""

	default:
		return "interface{}", ""
	}
}

func attributeNameConverter(attrName string) string {

	return strings.Title(attrName)
}

func attrToType(set spec.SpecificationSet, shadow bool, attr *spec.Attribute) string {

	var pointer string
	if mode, ok := attr.Extensions["refMode"]; ok && mode == "pointer" {
		pointer = "*"
	}

	var pointerShadow string
	if shadow {
		pointerShadow = "*"
	}

	var convertedType string
	switch attr.Type {
	case spec.AttributeTypeRef:
		convertedType = pointerShadow + pointer + set.Specification(attr.SubType).Model().EntityName
	case spec.AttributeTypeRefList:
		remoteSpec := set.Specification(attr.SubType)
		if remoteSpec.Model().Detached {
			convertedType = pointerShadow + "[]" + pointer + remoteSpec.Model().EntityName
		} else {
			convertedType = pointerShadow + pointer + remoteSpec.Model().EntityNamePlural + "List"
		}
	case spec.AttributeTypeRefMap:
		convertedType = pointerShadow + "map[string]" + pointer + set.Specification(attr.SubType).Model().EntityName
	default:
		convertedType = pointerShadow + attr.ConvertedType
	}

	if strings.HasPrefix(convertedType, "**") {
		convertedType = convertedType[1:]
	}

	return convertedType
}

func attrToField(set spec.SpecificationSet, shadow bool, attr *spec.Attribute) string {

	exposedName := attr.Name
	if attr.ExposedName != "" {
		exposedName = attr.ExposedName
	}

	json := exposedName
	msgpack := exposedName
	bson := strings.ToLower(attr.Name)

	if extname, ok := attr.Extensions["bson_name"].(string); ok {
		bson = extname
	}

	if !attr.Exposed {
		json = "-"
		msgpack = "-"
	}

	if attr.Exposed && (attr.OmitEmpty || shadow) {
		json += ",omitempty"
		msgpack += ",omitempty"
	}

	if !attr.Stored {
		bson = "-"
	} else if attr.Identifier {
		bson = "-"
	} else if shadow || attr.OmitEmpty {
		bson += ",omitempty"
	}

	descLines := strings.Split(attr.Description, "\n")
	for i := 0; i < len(descLines); i++ {
		descLines[i] = "// " + descLines[i]
	}

	convertedType := attrToType(set, shadow, attr)

	return fmt.Sprintf(
		"%s\n    %s %s `json:\"%s\" msgpack:\"%s\" bson:\"%s\" mapstructure:\"%s,omitempty\"`\n\n",
		strings.Join(descLines, "\n"),
		attr.ConvertedName,
		convertedType,
		json,
		msgpack,
		bson,
		strings.Replace(json, ",omitempty", "", 1),
	)
}

func attrToMongoField(set spec.SpecificationSet, shadow bool, attr *spec.Attribute) string {

	if !attr.Stored {
		panic(fmt.Sprintf("cannot use attrToMongoField on a non stored attribute: %s", attr.Name))
	}

	bson := strings.ToLower(attr.Name)

	if extname, ok := attr.Extensions["bson_name"].(string); ok {
		bson = extname
	}

	if attr.Identifier {
		bson = "_id,omitempty"
	} else if shadow || attr.OmitEmpty {
		bson += ",omitempty"
	}

	var convertedType string

	if attr.Identifier {
		convertedType = "bson.ObjectId"
	} else {
		convertedType = attrToType(set, shadow, attr)
	}

	return fmt.Sprintf(
		"%s %s `bson:\"%s\"`",
		attr.ConvertedName,
		convertedType,
		bson,
	)
}

func escapeBackticks(str string) string {
	return strings.Replace(str, "`", "`+\"`\"+`", -1)
}

func buildEnums(s spec.Specification, version string) []Enum {

	var enums []Enum // nolint
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

	if s.Model().Detached {
		return false
	}

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

	if s.Model().Detached {
		return false
	}

	if publicMode {
		return !s.Model().Private
	}

	return true
}

func shouldRegisterInnerRelationship(set spec.SpecificationSet, restName string, publicMode bool) bool {

	s := set.Specification(restName)

	if s.Model().Detached {
		return false
	}

	if publicMode {
		return !s.Model().Private
	}

	return true
}

func writeInitializer(set spec.SpecificationSet, s spec.Specification, attr *spec.Attribute) string {

	if attr.Initializer == "" &&
		attr.DefaultValue == nil &&
		attr.Type != spec.AttributeTypeList &&
		attr.Type != spec.AttributeTypeRef &&
		attr.Type != spec.AttributeTypeRefList &&
		attr.Type != spec.AttributeTypeRefMap {
		return ""
	}

	if ok1, ok2 := attr.Extensions["noInit"].(bool); ok2 && ok1 {
		return ""
	}

	return fmt.Sprintf("%s: %s,", attr.ConvertedName, writeDefaultValue(set, s, attr))
}

func writeDefaultValue(set spec.SpecificationSet, s spec.Specification, attr *spec.Attribute) string {

	if attr.Initializer != "" {
		return attr.Initializer
	}

	var pointer string
	var ref string
	if mode, ok := attr.Extensions["refMode"]; ok && mode == "pointer" {
		pointer = "*"
		ref = "New"
	}

	switch attr.Type {

	case spec.AttributeTypeList:
		if attr.DefaultValue == nil {
			return ref + attr.ConvertedType + "{}"
		}

	case spec.AttributeTypeRef:
		return ref + set.Specification(attr.SubType).Model().EntityName + "()"

	case spec.AttributeTypeRefList:
		remoteSpec := set.Specification(attr.SubType)
		if remoteSpec.Model().Detached {
			return "[]" + pointer + remoteSpec.Model().EntityName + "{}"
		}
		return pointer + remoteSpec.Model().EntityNamePlural + "List{}"

	case spec.AttributeTypeRefMap:
		return "map[string]" + pointer + set.Specification(attr.SubType).Model().EntityName + "{}"
	}

	var prefix string
	if attr.Type == spec.AttributeTypeEnum {
		prefix = s.Model().EntityName + attr.ConvertedName
	}

	return crawl(reflect.ValueOf(attr.DefaultValue), prefix)
}

func crawl(val reflect.Value, prefix string) string {

	switch val.Kind() {

	case reflect.Bool:
		if val.Bool() {
			return "true"
		}
		return "false"

	case reflect.String:
		if prefix != "" {
			return fmt.Sprintf(`%s%s`, prefix, val.String())
		}
		return fmt.Sprintf(`"%s"`, val.String())

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		return fmt.Sprintf(`%d`, val.Int())

	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf(`%f`, val.Float())

		// case reflect.Map:

	case reflect.Slice:

		out := "[]" + val.Index(0).Elem().Kind().String() + "{\n"
		for i := 0; i < val.Len(); i++ {
			out += fmt.Sprintf("%s,\n", crawl(val.Index(i).Elem(), prefix))
		}
		out += "}"

		return out
	}

	return ""
}

func sortAttributes(attrs []*spec.Attribute) []*spec.Attribute {

	out := make([]*spec.Attribute, len(attrs))
	copy(out, attrs)

	sort.Slice(out, func(i int, j int) bool {
		return strings.Compare(attrs[i].ConvertedName, attrs[j].ConvertedName) == -1
	})

	return out
}
