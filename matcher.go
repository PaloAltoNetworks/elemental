package elemental

import (
	"fmt"
	"reflect"
)

// MatchesFilter determines whether an identity matches a filter
func MatchesFilter(identifiable AttributeSpecifiable, filter *Filter, opts ...MatcherOption) (bool, error) {
	if identifiable == nil {
		panic(fmt.Errorf("elemental: identifiable cannot be nil"))
	}

	return matcher(identifiable, filter, make(map[*Filter]bool))
}

func matcher(identifiable AttributeSpecifiable, filter *Filter, seen map[*Filter]bool) (bool, error) {

	if filter == nil {
		panic(fmt.Errorf("elemental: filter cannot be nil"))
	}

	matched := true
	var err error

	for i, op := range filter.Operators() {
		if result, seen := seen[filter]; seen {
			return result, nil
		}

		switch op {
		case AndOperator:
			comparator := filter.Comparators()[i]
			attributeName := filter.Keys()[i]
			attributeValue := identifiable.ValueForAttribute(attributeName)

			// important: when adding support for a new comparator, you must only short circuit and return under two conditions:
			//    - an error occurred by your comparator
			//    - the comparator failed to find a match
			//
			// this is because we are dealing with AND semantics here; success is only guaranteed if all AndOperator's find a match

			switch comparator {
			case EqualComparator:
				seen[filter] = equals(attributeValue, filter.Values()[i][0])
				if !seen[filter] {
					return false, nil
				}
			case NotEqualComparator,
				GreaterComparator,
				GreaterOrEqualComparator,
				LesserComparator,
				LesserOrEqualComparator,
				InComparator,
				NotInComparator,
				ContainComparator,
				NotContainComparator,
				MatchComparator,
				NotMatchComparator,
				ExistsComparator,
				NotExistsComparator:
				return false, fmt.Errorf("elemental: unsuported comparator %q", translateComparator(filter.Comparators()[i]))
			default:
				panic(fmt.Errorf("elemental: unknown comparator %q", translateComparator(filter.Comparators()[i])))
			}
		case OrFilterOperator:
			for _, f := range filter.OrFilters()[i] {
				// only one 'or' filter must match for it to be considered a successful match
				if matched, err = matcher(identifiable, f, seen); err != nil || matched {
					break
				}
			}
		case AndFilterOperator:
			var subFilterMatched bool
			for _, f := range filter.AndFilters()[i] {
				// all 'and' filters must match for it to be considered a successful match
				subFilterMatched, err = matcher(identifiable, f, seen)
				matched = matched && subFilterMatched
				if err != nil || !matched {
					break
				}
			}
		}
	}

	return matched, err
}

// equals implements the elemental.EqualComparator behaviour by implementing the Go equivalent of
// https://docs.mongodb.com/manual/reference/operator/query/eq
//
// { <field>: { $eq: <value> } }
func equals(field, value interface{}) bool {

	// deal with the `nil` case before anything else
	// this is for queries that are checking whether an attribute does not exist, for example:
	//     db.getCollection('somecollection').find({ invalidAttribute: { $eq: null } })
	//     in the query above, all documents that DO NOT contain 'invalidAttribute' will be returned
	// so the equivalent translation of that for equality is to return true
	if field == nil && value == nil {
		return true
	}

	// check to see if we are dealing with an attribute that does not exist on the provided identifiable.
	// recall `field` will be nil in the event that `ValueForAttribute` returns an empty nil interface
	if field == nil {
		return false
	}

	fieldV, valueV := reflect.ValueOf(field), reflect.ValueOf(value)
	valueArrayLike, fieldArrayLike := isArrayLike(valueV), isArrayLike(fieldV)

	switch {
	case valueArrayLike && fieldArrayLike:

		// try handling first case:
		//   if the specified <value> is an array, MongoDB matches documents where the <field> matches the array exactly.
		//   the order of the elements matters.

		if valueV.Len() == fieldV.Len() {
			var failed bool
			for i := 0; i < valueV.Len(); i++ {
				ve, fe := valueV.Index(i), fieldV.Index(i)
				if !reflect.DeepEqual(ve.Interface(), fe.Interface()) {
					failed = true
					break
				}
			}

			if !failed {
				return true
			}
		}

		// try handling second case if first case failed to find a match:
		//   the <field> contains an element that matches the array exactly.

		fallthrough
	case fieldArrayLike:

		adaptedValue := valueV
		if valueArrayLike {

			// for convenience, try to coerce the comparator's key value from an array to a slice or vice versa to match
			// the type of the attribute elements. Any further type mismatches due to inner containers being different
			// are the callers responsibility

			// check to see if the attribute is an array/slice of arrays/slices
			switch fieldV.Type().Elem().Kind() {
			case reflect.Slice:
				// do we have a mismatch? (i.e. comparator key type is an array not a slice)
				// if so, let's change it to slice
				if valueV.Kind() != reflect.Slice {
					adaptedValue = reflect.MakeSlice(reflect.SliceOf(valueV.Type().Elem()), 0, valueV.Len())
					for i := 0; i < valueV.Len(); i++ {
						adaptedValue = reflect.Append(adaptedValue, valueV.Index(i))
					}
				}
			case reflect.Array:
				// do we have a mismatch? (i.e. comparator key type is a slice not an array)
				// if so, let's change it to an array
				if valueV.Kind() != reflect.Array {
					adaptedValue = reflect.New(reflect.ArrayOf(valueV.Len(), valueV.Type().Elem())).Elem()
					for i := 0; i < valueV.Len(); i++ {
						adaptedValue.Index(i).Set(valueV.Index(i))
					}
				}
			}
		}

		// calling the 'Interface()' method on an invalid value will result in a panic, exercise caution here
		valueI := value
		if adaptedValue.IsValid() {
			valueI = adaptedValue.Interface()
		}

		for i := 0; i < fieldV.Len(); i++ {
			if reflect.DeepEqual(fieldV.Index(i).Interface(), valueI) {
				return true
			}
		}
	default:
		// if our field and value are not arrays/slices, then we just do a recursive equality check using Go's `==` operator via
		// reflect.DeepEqual
		if reflect.DeepEqual(field, value) {
			return true
		}
	}

	return false
}

func isArrayLike(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	default:
		return false
	}
}
