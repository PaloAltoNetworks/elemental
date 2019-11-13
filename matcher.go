package elemental

import (
	"fmt"
	"reflect"
)

// MatchesFilter determines whether an identity matches a filter
func MatchesFilter(identifiable AttributeSpecifiable, filter *Filter, opts ...MatcherOption) (bool, error) {

	if filter == nil {
		panic(fmt.Errorf("elemental: filter cannot be nil"))
	}

	if identifiable == nil {
		panic(fmt.Errorf("elemental: identifiable cannot be nil"))
	}

	matched := true
	var err error

	for i, op := range filter.Operators() {
		switch op {
		case AndOperator:
			comparator := filter.Comparators()[i]
			attributeName := filter.Keys()[i]
			attributeValue := identifiable.ValueForAttribute(attributeName)

			// important: when adding support for a new comparator, you must only short circuit and return under two conditions:
			//    - an error occurred by your comparator
			//    - the comparator failed to find a match
			//
			// this is because we are dealing with AND semantics here; success is only possible when all AndOperator's find a match

			switch comparator {
			case EqualComparator:
				if !equals(attributeValue, filter.Values()[i][0]) {
					return false, nil
				}
			case NotEqualComparator:
				if !notEquals(attributeValue, filter.Values()[i][0]) {
					return false, nil
				}
			case ExistsComparator:
				if !exists(attributeName, identifiable.AttributeSpecifications()) {
					return false, nil
				}
			case NotExistsComparator:
				if !notExists(attributeName, identifiable.AttributeSpecifications()) {
					return false, nil
				}
			case
				MatchComparator,
				NotMatchComparator,
				GreaterComparator,
				GreaterOrEqualComparator,
				LesserComparator,
				LesserOrEqualComparator,
				InComparator,
				NotInComparator,
				ContainComparator,
				NotContainComparator:
				return false, fmt.Errorf("elemental: unsuported comparator %q", translateComparator(filter.Comparators()[i]))
			default:
				panic(fmt.Errorf("elemental: unknown comparator %q", translateComparator(filter.Comparators()[i])))
			}
		case OrFilterOperator:
			for _, f := range filter.OrFilters()[i] {
				// only one 'or' filter must match for it to be considered a successful match
				if matched, err = MatchesFilter(identifiable, f); err != nil || matched {
					break
				}
			}
		case AndFilterOperator:
			var subFilterMatched bool
			for _, f := range filter.AndFilters()[i] {
				// all 'and' filters must match for it to be considered a successful match
				subFilterMatched, err = MatchesFilter(identifiable, f)
				matched = matched && subFilterMatched
				if err != nil || !matched {
					break
				}
			}
		}
	}

	return matched, err
}

// exists implements the elemental.ExistsComparator behaviour by implementing the Go equivalent of
// https://docs.mongodb.com/manual/reference/operator/query/exists/ where the value of the boolean is TRUE
//
// { field: { $exists: <boolean> (true) } }
//  Quote from docs:
//     When <boolean> is true, $exists matches the documents that contain the field, including documents where
//     the field value is null.
//
// exists will return true as long as the identifiable has the attribute irrespective of its value (even if it is nil)
func exists(attributeName string, attributes map[string]AttributeSpecification) bool {
	// check to see if we are dealing with an attribute that does not exist on the provided identifiable.
	_, exists := attributes[attributeName]
	return exists
}

// notExists implements the elemental.NotExistsComparator by implementing the Go equivalent of
// https://docs.mongodb.com/manual/reference/operator/query/exists/ where the value of the boolean is FALSE
//
// { field: { $exists: <boolean> (false) } }
//
// Quote from docs:
//     if <boolean> is false, the query returns only the documents that do not contain the field.
func notExists(attributeName string, attributes map[string]AttributeSpecification) bool {
	_, exists := attributes[attributeName]
	return !exists
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
	//     so the equivalent translation of that for equality is to return true
	if field == nil && value == nil {
		return true
	}

	return equalsCommon(field, value)
}

// notEquals implements the elemental.NotEqualComparator behaviour by implementing the Go equivalent of
// https://docs.mongodb.com/manual/reference/operator/query/ne
//
// {field: {$ne: value} }
func notEquals(field, value interface{}) bool {

	// deals with the 'nil' case where an attribute that does not exist with a null value has been specified
	// Example query:
	//     db.getCollection('someCollection').find({ invalidAttribute: { $ne: null } })
	//     in the query above, no match will ever be possible, therefore the equivalent behavioural translation for this
	//     is to simply return false
	if field == nil && value == nil {
		return false
	}

	return !equalsCommon(field, value)
}

func equalsCommon(field, value interface{}) bool {

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
