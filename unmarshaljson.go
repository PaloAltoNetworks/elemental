package elemental

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"sync"
	"unicode"
)

const (
	invalidJSON      = `Invalid JSON`
	invalidAttribute = `Attribute '%s' is invalid`
	readerError      = `Something went wrong in the server when reading the body of the request`
	typeError        = `Something went wrong in the server when analyzing the JSON`
	wrongType        = `Data '%v' of attribute '%s' should be a '%s'`
)

// UnmarshalJSON unmarshal the given reader and returns detailed errors if found
func UnmarshalJSON(r io.Reader, i interface{}) error {

	data, err := ioutil.ReadAll(r)

	if err != nil {
		return NewError("Validation Error", fmt.Sprintf(readerError), "elemental", http.StatusUnprocessableEntity)
	}

	errors := Errors{}

	var d map[string]interface{}
	err = json.Unmarshal(data, &d)

	if err != nil {
		errors = append(errors, NewError("Validation Error", fmt.Sprintf(invalidJSON), "elemental", http.StatusUnprocessableEntity))
		return errors
	}

	interfaceValue := reflect.ValueOf(i)

	if interfaceValue.Kind() == reflect.Ptr {
		interfaceValue = interfaceValue.Elem()
	}

	err = json.Unmarshal(data, i)

	if err == nil {
		return nil
	}

	// Check for type error here
	for k, v := range d {
		field, err := fieldForName(interfaceValue.Type(), k)

		// Invalid field already checked above
		if err != nil {
			continue
		}

		b, err := json.MarshalIndent(v, "", "")

		// Should not go there
		if err != nil {
			errors = append(errors, NewError("Validation Error", fmt.Sprintf(typeError), "elemental", http.StatusUnprocessableEntity))
			continue
		}

		j := fmt.Sprintf(`{ "%s" : %s}`, k, string(b))
		err = json.Unmarshal([]byte(j), i)

		if err == nil {
			continue
		}

		t := field.typ.String()

		if field.typ.Kind() == reflect.String {
			t = field.typ.Kind().String()
		}

		if err != nil || t != reflect.ValueOf(v).Type().String() {

			fieldType := field.typ.Kind().String()

			if field.typ.Kind() != reflect.String {
				fieldType = field.typ.String()
			}

			if field.typ.String() == "time.Time" {
				fieldType = "string in format YYYY-MM-DDTHH:MM:SSZ"
			}

			errors = append(errors, NewError("Validation Error", fmt.Sprintf(wrongType, v, k, fieldType), "elemental", http.StatusUnprocessableEntity))
			continue
		}
	}

	// Should not enter here, but in case we create a serve error if we don't find the json error
	if len(errors) == 0 {
		errors = append(errors, NewError("Validation Error", fmt.Sprintf(typeError), "elemental", http.StatusUnprocessableEntity))
	}

	return errors
}

func fieldForName(t reflect.Type, name string) (*field, error) {
	structFields := cachedTypeFields(t)

	for _, info := range structFields {
		if info.name == name {
			return &info, nil
		}
	}

	return nil, fmt.Errorf("Field %s not found", name)
}

type cache struct {
	sync.RWMutex
	m map[reflect.Type][]field
}

var fieldCache cache

// cachedTypeFields is like typeFields but uses a cache to avoid repeated work.
func cachedTypeFields(t reflect.Type) []field {

	fieldCache.RLock()
	f := fieldCache.m[t]
	fieldCache.RUnlock()

	if f != nil {
		return f
	}

	f = typeFields(t)

	fieldCache.Lock()

	if fieldCache.m == nil {
		fieldCache.m = map[reflect.Type][]field{}
	}

	fieldCache.m[t] = f
	fieldCache.Unlock()

	return f
}

// A field represents a single field found in a struct.
type field struct {
	name string

	tag       bool
	index     []int
	typ       reflect.Type
	omitEmpty bool
	quoted    bool
}

// byName sorts field by name, breaking ties with depth,
// then breaking ties with "name came from tag", then
// breaking ties with index sequence.
type byName []field

func (x byName) Len() int {
	return len(x)
}

func (x byName) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x byName) Less(i, j int) bool {

	if x[i].name != x[j].name {
		return x[i].name < x[j].name
	}

	if len(x[i].index) != len(x[j].index) {
		return len(x[i].index) < len(x[j].index)
	}

	if x[i].tag != x[j].tag {
		return x[i].tag
	}

	return byIndex(x).Less(i, j)
}

// byIndex sorts field by index sequence.
type byIndex []field

func (x byIndex) Len() int {
	return len(x)
}

func (x byIndex) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x byIndex) Less(i, j int) bool {

	for k, xik := range x[i].index {

		if k >= len(x[j].index) {
			return false
		}

		if xik != x[j].index[k] {
			return xik < x[j].index[k]
		}
	}

	return len(x[i].index) < len(x[j].index)
}

// typeFields returns a list of fields that should be recognized for the given type.
// The algorithm is breadth-first search over the set of structs to include - the top struct
// and then any reachable anonymous structs.
func typeFields(t reflect.Type) []field {

	// Anonymous fields to explore at the current level and the next.
	current := []field{}
	next := []field{{typ: t}}

	// Count of queued names for current level and the next.
	count := map[reflect.Type]int{}
	nextCount := map[reflect.Type]int{}

	// Types already visited at an earlier level.
	visited := map[reflect.Type]bool{}

	// Fields found.
	var fields []field

	for len(next) > 0 {

		current, next = next, current[:0]
		count, nextCount = nextCount, map[reflect.Type]int{}

		for _, f := range current {

			if visited[f.typ] {
				continue
			}

			visited[f.typ] = true

			// Scan f.typ for fields to include.
			for i := 0; i < f.typ.NumField(); i++ {

				sf := f.typ.Field(i)
				if sf.PkgPath != "" && !sf.Anonymous { // unexported
					continue
				}

				tag := tagForField(sf)
				if tag == "-" {
					continue
				}

				name, _ := parseTag(tag)
				if !isValidTag(name) {
					name = ""
				}

				index := make([]int, len(f.index)+1)
				copy(index, f.index)
				index[len(f.index)] = i

				ft := sf.Type
				if ft.Name() == "" && ft.Kind() == reflect.Ptr {
					// Follow pointer.
					ft = ft.Elem()
				}

				// Record found field and index sequence.
				if name != "" || !sf.Anonymous || ft.Kind() != reflect.Struct {

					tagged := name != ""
					if name == "" {
						name = sf.Name
					}
					fields = append(fields, field{
						name:  name,
						tag:   tagged,
						index: index,
						typ:   ft,
					})
					if count[f.typ] > 1 {
						// If there were multiple instances, add a second,
						// so that the annihilation code will see a duplicate.
						// It only cares about the distinction between 1 or 2,
						// so don't bother generating any more copies.
						fields = append(fields, fields[len(fields)-1])
					}
					continue
				}

				// Record new anonymous struct to explore in next round.
				nextCount[ft]++
				if nextCount[ft] == 1 {
					next = append(next, field{name: ft.Name(), index: index, typ: ft})
				}
			}
		}
	}

	sort.Sort(byName(fields))

	// Delete all fields that are hidden by the Go rules for embedded fields,
	// except that fields with valid tags are promoted.

	// The fields are sorted in primary order of name, secondary order
	// of field index length. Loop over names; for each name, delete
	// hidden fields by choosing the one dominant field that survives.
	out := fields[:0]
	for advance, i := 0, 0; i < len(fields); i += advance {
		// One iteration per name.
		// Find the sequence of fields with the name of this first field.
		fi := fields[i]
		name := fi.name
		for advance = 1; i+advance < len(fields); advance++ {
			fj := fields[i+advance]
			if fj.name != name {
				break
			}
		}
		if advance == 1 { // Only one field with this name
			out = append(out, fi)
			continue
		}
		dominant, ok := dominantField(fields[i : i+advance])
		if ok {
			out = append(out, dominant)
		}
	}

	fields = out
	sort.Sort(byIndex(fields))

	return fields
}

// dominantField looks through the fields, all of which are known to
// have the same name, to find the single field that dominates the
// others using Go's embedding rules, modified by the presence of
// valid tags. If there are multiple top-level fields, the boolean
// will be false: This condition is an error in Go and we skip all
// the fields.
func dominantField(fields []field) (field, bool) {
	// The fields are sorted in increasing index-length order. The winner
	// must therefore be one with the shortest index length. Drop all
	// longer entries, which is easy: just truncate the slice.
	length := len(fields[0].index)
	tagged := -1 // Index of first tagged field.

	for i, f := range fields {
		if len(f.index) > length {
			fields = fields[:i]
			break
		}
		if f.tag {
			if tagged >= 0 {
				// Multiple tagged fields at the same level: conflict.
				// Return no field.
				return field{}, false
			}
			tagged = i
		}
	}
	if tagged >= 0 {
		return fields[tagged], true
	}
	// All remaining fields have the same length. If there's more than one,
	// we have a conflict (two fields named "X" at the same level) and we
	// return no field.
	if len(fields) > 1 {
		return field{}, false
	}
	return fields[0], true
}

// TagName is the default tagName for this lib
const TagName = "json"

// tagOptions is the string following a comma in a struct field's
// tag, or the empty string. It does not include the leading comma.
type tagOptions string

func tagForField(sf reflect.StructField) string {
	return sf.Tag.Get(TagName)
}

// parseTag splits a struct field's tag into its name and
// comma-separated options.
func parseTag(tag string) (string, tagOptions) {

	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}

	return tag, tagOptions("")
}

func isValidTag(s string) bool {

	if s == "" {
		return false
	}

	for _, c := range s {

		switch {
		case strings.ContainsRune("!#$%&()*+-./:<=>?@[]^_{|}~ ", c):
			// Backslash and quote chars are reserved, but
			// otherwise any punctuation chars are allowed
			// in a tag name.
		default:
			if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
				return false
			}
		}
	}

	return true
}

// Contains returns whether checks that a comma-separated list of options
// contains a particular substr flag. substr must be surrounded by a
// string boundary or commas.
func (o tagOptions) Contains(optionName string) bool {

	if len(o) == 0 {
		return false
	}

	s := string(o)

	for s != "" {

		var next string
		i := strings.Index(s, ",")

		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}

		s = next
	}

	return false
}
