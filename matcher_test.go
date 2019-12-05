package elemental_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"go.aporeto.io/elemental"
	"go.aporeto.io/elemental/internal"
	testmodel "go.aporeto.io/elemental/test/model"
)

// this unit test suite tests the functionality of the EqualComparator when used in conjunction with the helper
// MatchesFilter for filtering an AttributeSpecifiable using the supplied filter
func TestEqualComparator(t *testing.T) {

	testAttributeName := "someAttribute"
	tests := map[string]struct {
		filter        *elemental.Filter
		mockSetupFunc func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable
		expectedMatch bool
		expectedError bool
	}{
		// https://docs.mongodb.com/manual/reference/operator/query/eq/#equals-an-array-value
		// quote: "If the specified <value> is an array, MongoDB matches documents where the <field> matches the array exactly"

		"should return true if the filter comparator key value is an ARRAY and the attribute is an array that matches the key value exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals([4]string{
					"a",
					"b",
					"c",
					"d",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([4]string{
						"a",
						"b",
						"c",
						"d",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},

		// https://docs.mongodb.com/manual/reference/operator/query/eq/#equals-an-array-value
		// quote: "If the specified <value> is an array, MongoDB matches documents where the <field> matches the array exactly"

		// notice that the ACTUAL attribute value here is a string ARRAY of length 4, but the filter key was a slice
		// this is just our way of making the API developer friendly as the mongo reference manual obviously has no concept of
		// a slice.

		"should return true if the filter key value is a SLICE and the attribute is an array that matches the key value exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals([]string{
					"a",
					"b",
					"c",
					"d",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([4]string{
						"a",
						"b",
						"c",
						"d",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},

		// https://docs.mongodb.com/manual/reference/operator/query/eq/#equals-an-array-value
		//
		// as per stated in the reference manual, in the event that both the value and field are arrays, they should only
		// be considered a match if there is an EXACT match...ORDER MATTERS

		"should return false if the order of the array value does not match the attribute's value even if they both contain the same elements": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals([]string{
					"d",
					"b",
					"c",
					"a",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{

						// attention: notice how the order of elements here is different!
						// the filter wanted a match on "d,b,c,a", however the actual value of the attribute was "a,b,c,d"

						"a",
						"b",
						"c",
						"d",
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},

		// https://docs.mongodb.com/manual/reference/operator/query/eq/#equals-an-array-value
		// quote: "...or the <field> contains an element that matches the array exactly."
		//
		// in this case the attribute value is a slice that CONTAINS the comparator's key value which happens to be a slice

		"should return true if the attribute being filtered on is a slice that contains the comparator's key value (which is a slice) as an element": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals([]string{
					"amir",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([][]string{
						{"eric"},
						{"antoine"},
						{"chris"},
						{"amir"},
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},

		// https://docs.mongodb.com/manual/reference/operator/query/eq/#equals-an-array-value
		// quote: "...or the <field> contains an element that matches the array exactly."
		//
		// in this case the attribute value is an array that CONTAINS the comparator's key value which happens to be a slice

		"should return true if the attribute being filtered on is an array of arrays that contains the comparator's key value (which is a slice) as an element": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals([]string{
					"henry",
					"jessica",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).

					// notice how this is an array that contains the filter key as an element

					Return(interface{}([3][2]string{
						{"yoda", "bear"},
						{"alpha", "master"},
						{"henry", "jessica"},
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},

		// https://docs.mongodb.com/manual/reference/operator/query/eq/#equals-an-array-value
		// quote: "...or the <field> contains an element that matches the array exactly."
		//
		// in this case the attribute value is an array of slices that CONTAINS the comparator's key value which is also an array

		"should return true if the attribute being filtered on is an array of slices that contains the comparator's key value (which is an array) as an element": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).

				// notice how this is an array [1]string and not a slice

				Equals([1]string{
					"frank",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).

					// notice how this is an array of slices of length 4 that contains the filter key value as an element

					Return(interface{}([4][]string{
						{"bill"},
						{"bob"},
						{"john"},
						{"frank"},
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		// https://docs.mongodb.com/manual/reference/operator/query/eq/#array-element-equals-a-value
		"should return true if the attribute being filtered on (a slice) contains the comparator key's value": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals("amir").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).

					// the attribute contains "amir", so we expect a match to occur

					Return(interface{}([]string{
						"billy",
						"bob",
						"aaron",
						"john",
						"muhammad",
						"amir",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		// https://docs.mongodb.com/manual/reference/operator/query/eq/#array-element-equals-a-value
		"should return false if the attribute being filtered on (a slice) does not contain the comparator key's value": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals("amir").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).

					// the attribute does NOT contain "amir", so we do NOT expect a match to occur

					Return(interface{}([]string{
						"billy",
						"bob",
						"aaron",
						"john",
						"muhammad",
						"zlatan",
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		// https://docs.mongodb.com/manual/reference/operator/query/eq/#array-element-equals-a-value
		"should return true if the attribute being filtered on (an array) contains the comparator key's value": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals("amir").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).

					// the attribute contains "amir", so we expect a match to occur

					Return(interface{}([...]string{
						"billy",
						"bob",
						"aaron",
						"john",
						"muhammad",
						"amir",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		// https://docs.mongodb.com/manual/reference/operator/query/eq/#array-element-equals-a-value
		"should return false if the attribute being filtered on (an array) does not contain the comparator key's value": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals("amir").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).

					// the attribute does NOT contain "amir", so we do NOT expect a match to occur

					Return(interface{}([...]string{
						"billy",
						"bob",
						"aaron",
						"john",
						"muhammad",
						"zlatan",
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		// https://docs.mongodb.com/manual/reference/operator/query/eq/#array-element-equals-a-value
		"should do deep equality check via Go's == operator if both field and array ": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals(struct {
					name string
					age  int
				}{
					name: "john",
					age:  98,
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(struct {
						name string
						age  int
					}{
						name: "john",
						age:  98,
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},

		// this is for queries that are checking whether an attribute does not exist on an identifiable, for example:
		//     db.getCollection('somecollection').find({ invalidAttribute: { $eq: null } })
		//     in the query above, all documents that DO NOT contain 'invalidAttribute' will be returned

		"should result in a match when attempting to match on an attribute that does not exist with the value of 'nil'": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals(nil).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nil))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return false if attempting to match on an attribute that does not exist for the given identifiable": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Equals("some value").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).

					// fyi: if 'ValueForAttribute' returns the `nil` interface, this means that the attribute did not exist
					// on the identifiable that was being queried with the equality comparator

					Return(interface{}(nil))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should be able to handle a nil value even if the attribute exists": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				// passing in nil for the value
				Equals(nil).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{"amir"}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
	}

	for description, tc := range tests {
		t.Run(description, func(t *testing.T) {

			identity := tc.mockSetupFunc(t, gomock.NewController(t))
			matched, err := elemental.MatchesFilter(identity, tc.filter)

			if (err != nil) != tc.expectedError {
				t.Errorf("\n"+
					"error expectation failued:\n"+
					"expected an error: %t\n"+
					"actual error: %+v\n",
					tc.expectedError,
					err)
			}

			if matched != tc.expectedMatch {
				t.Errorf("\n"+
					"match expectation failed:\n"+
					"expected a match: %t\n"+
					"matched occurred: %+v\n",
					tc.expectedMatch,
					matched)
			}
		})
	}
}

func TestUnexportedTypeEmbedded_EqualComparator(t *testing.T) {

	// a type with a field that has a type that contains unexported fields!
	type EmbeddedType struct {
		Time time.Time
	}

	t1, t2 := EmbeddedType{}, EmbeddedType{}
	t1.Time = time.Now()

	timeData, err := t1.Time.MarshalJSON()
	if err != nil {
		t.Fatalf("failed to setup test fixture, error: %+v", err)
	}

	timeCopy := time.Time{}
	if err := timeCopy.UnmarshalJSON(timeData); err != nil {
		t.Fatalf("failed to setup test fixture, error: %+v", err)
	}

	t2.Time = timeCopy

	testAttributeName := "someAttribute"
	ctrl := gomock.NewController(t)

	mockIdentity := internal.NewMockAttributeSpecifiable(ctrl)
	mockIdentity.EXPECT().
		ValueForAttribute(testAttributeName).
		Return(interface{}(t1))

	filter := elemental.NewFilterComposer().
		WithKey(testAttributeName).
		Equals(t2).
		Done()

	// why doesn't this match? aren't they technically referring to the same time?
	// it's because reflect.DeepEquals is also traversing unexported fields of time.Time; it could not compare time.Time
	// well as it is comparing the unexported monotonic timestamps fields

	if matched, err := elemental.MatchesFilter(mockIdentity, filter); err != nil {
		t.Errorf("did not expect to get an error, but received: %+v\n", err)
	} else if matched {
		t.Errorf("did not expect a match because the private/unexported fields are different\n")
	}
}

// TestMatchesFilter validates the correctness of the matcher algorithm by exercising all the various filter permutations
//
// note: this test is NOT verifying the logic of the supported comparators, individual test suites should exist for each comparator
// implementation.
func TestMatchesFilter(t *testing.T) {

	testCases := map[string]struct {
		filter        *elemental.Filter
		mockSetupFunc func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable
		expectedPanic bool
		expectedMatch bool
		expectedError bool
	}{
		"should panic if you pass in a nil filter": {
			filter: nil,
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				// our mock should never be called in this case
				mockAS.
					EXPECT().
					ValueForAttribute(gomock.Any()).
					Times(0)
				return mockAS
			},
			expectedPanic: true,
			expectedMatch: false,
			expectedError: false,
		},
		"should panic if you pass in a nil identifiable": {
			filter: elemental.NewFilterComposer().Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				return nil
			},
			expectedPanic: true,
			expectedMatch: false,
			expectedError: false,
		},

		// checks whether AND semantics are working as expected

		"should return false if a sub-filter within the AND operator fails to find a match": {
			filter: elemental.NewFilterComposer().
				And(
					// sub-filter one
					elemental.NewFilterComposer().
						WithKey("attr1").
						Equals(true).
						Done(),

					// sub-filter two
					elemental.NewFilterComposer().
						WithKey("attr2").
						Equals(true).
						Done(),
				).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				// this will match the first sub-filter
				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return(interface{}(true)).
					Times(1)

				// this will NOT match the second sub-filter as we are now returning false when true is expected
				mockAS.EXPECT().
					ValueForAttribute("attr2").
					Return(interface{}(false)).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},

		// checks whether AND semantics are working as expected

		"should return true if ALL sub-filter within the AND operator are successful in finding a match": {
			filter: elemental.NewFilterComposer().
				And(
					elemental.NewFilterComposer().
						WithKey("attr1").
						Equals(true).
						Done(),
					elemental.NewFilterComposer().
						WithKey("attr2").
						Equals(true).
						Done(),
					elemental.NewFilterComposer().
						WithKey("attr3").
						Equals(true).
						Done(),
				).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return(interface{}(true)).
					Times(1)

				mockAS.EXPECT().
					ValueForAttribute("attr2").
					Return(interface{}(true)).
					Times(1)

				mockAS.EXPECT().
					ValueForAttribute("attr3").
					Return(interface{}(true)).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},

		// checks whether OR semantics are working as expected

		"should return false only if ALL sub-filters of the OR operator fail to find a match": {
			filter: elemental.NewFilterComposer().
				Or(
					// sub-filter one
					elemental.NewFilterComposer().
						WithKey("attr1").
						Equals(true).
						Done(),

					// sub-filter two
					elemental.NewFilterComposer().
						WithKey("attr2").
						Equals(true).
						Done(),
				).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				// setup our mock so none of our filters match

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return(interface{}(false)).
					Times(1)

				mockAS.EXPECT().
					ValueForAttribute("attr2").
					Return(interface{}(false)).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},

		// checks whether OR semantics are working as expected

		"should return true as long as one sub-filter of the OR operator is able to find a match": {
			filter: elemental.NewFilterComposer().
				Or(
					// sub-filter one
					elemental.NewFilterComposer().
						WithKey("attr1").
						Equals(true).
						Done(),

					// sub-filter two
					elemental.NewFilterComposer().
						WithKey("attr2").
						Equals(true).
						Done(),

					// sub-filter three
					elemental.NewFilterComposer().
						WithKey("attr3").
						Equals(true).
						Done(),
				).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				// setup mocks such that all sub-filters fail EXPECT ONE just to demonstrate that OR semantics are working as expected

				// make sub-filter one fail
				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return(interface{}(false)).
					Times(1)

				// make sub-filter two fail
				mockAS.EXPECT().
					ValueForAttribute("attr2").
					Return(interface{}(false)).
					Times(1)

				// make sub-filter three PASS
				mockAS.EXPECT().
					ValueForAttribute("attr3").
					Return(interface{}(true)).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true only if ALL AndOperator's were evaluated and found a match for with their respective comparator": {
			// this is a single filter with three AndOperator's
			filter: elemental.NewFilterComposer().
				WithKey("attr1").
				Equals(true).
				WithKey("attr2").
				Equals(true).
				WithKey("attr3").
				Equals(true).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return(interface{}(true)).
					Times(1)

				mockAS.EXPECT().
					ValueForAttribute("attr2").
					Return(interface{}(true)).
					Times(1)

				mockAS.EXPECT().
					ValueForAttribute("attr3").
					Return(interface{}(true)).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return false if as long as one AndOperator's comparator fails to find a match": {
			// this is a single filter with three AndOperator's
			filter: elemental.NewFilterComposer().
				WithKey("attr1").
				Equals(true).
				WithKey("attr2").
				Equals(true).
				WithKey("attr3").
				Equals(true).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return(interface{}(true))

				mockAS.EXPECT().
					ValueForAttribute("attr2").
					Return(interface{}(true))

				// make the third operator fail
				mockAS.EXPECT().
					ValueForAttribute("attr3").
					Return(interface{}(false))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return true only if the comparators for ALL AndOperators were able to find a match": {
			// this is a single filter with three AndOperator's
			filter: elemental.NewFilterComposer().
				WithKey("attr1").
				Equals(true).
				WithKey("attr2").
				Equals(true).
				WithKey("attr3").
				Equals(true).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return(interface{}(true))

				mockAS.EXPECT().
					ValueForAttribute("attr2").
					Return(interface{}(true))

				mockAS.EXPECT().
					ValueForAttribute("attr3").
					Return(interface{}(true))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should be able to handle a long chain of OrFilterOperator's - match case": {

			filter: elemental.NewFilterComposer().
				Or(
					elemental.NewFilterComposer().
						Or(
							elemental.NewFilterComposer().
								Or(
									elemental.NewFilterComposer().
										Or(
											elemental.NewFilterComposer().
												Or(
													elemental.NewFilterComposer().
														Or(
															elemental.NewFilterComposer().
																Or(
																	elemental.NewFilterComposer().
																		Or(elemental.NewFilterComposer().
																			WithKey("attr1").
																			Equals(true).
																			Done()).
																		Done()).
																Done()).
														Done()).
												Done()).
										Done()).
								Done()).
						Done(),
				).Done(),

			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return(interface{}(true)).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should be able to handle a long chain of OrFilterOperator's - no match case": {

			filter: elemental.NewFilterComposer().
				Or(
					elemental.NewFilterComposer().
						Or(
							elemental.NewFilterComposer().
								Or(
									elemental.NewFilterComposer().
										Or(
											elemental.NewFilterComposer().
												Or(
													elemental.NewFilterComposer().
														Or(
															elemental.NewFilterComposer().
																Or(
																	elemental.NewFilterComposer().
																		Or(elemental.NewFilterComposer().
																			WithKey("attr1").
																			Equals(false).
																			Done()).
																		Done()).
																Done()).
														Done()).
												Done()).
										Done()).
								Done()).
						Done(),
				).Done(),

			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return(interface{}(true)).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should be able to handle a long chain of AndFilterOperator's - match case": {

			filter: elemental.NewFilterComposer().
				And(
					elemental.NewFilterComposer().
						WithKey("attr1").
						Equals(true).
						Done(),
					elemental.NewFilterComposer().
						And(
							elemental.NewFilterComposer().
								WithKey("attr1").
								Equals(true).
								Done(),
							elemental.NewFilterComposer().
								And(
									elemental.NewFilterComposer().
										WithKey("attr1").
										Equals(true).
										Done(),
									elemental.NewFilterComposer().
										And(
											elemental.NewFilterComposer().
												WithKey("attr1").
												Equals(true).
												Done(),
											elemental.NewFilterComposer().
												And(
													elemental.NewFilterComposer().
														WithKey("attr1").
														Equals(false).
														Done(),
													elemental.NewFilterComposer().
														And(
															elemental.NewFilterComposer().
																WithKey("attr1").
																Equals(true).
																Done(),
															elemental.NewFilterComposer().
																And(
																	elemental.NewFilterComposer().
																		And(elemental.NewFilterComposer().
																			WithKey("attr1").
																			Equals(true).
																			Done()).
																		Done()).
																Done()).
														Done()).
												Done()).
										Done()).
								Done()).
						Done(),
				).Done(),

			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return([]bool{
						false,
						false,
						false,
						false,
						false,
						true,
					}).
					Times(7)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should be able to handle a long chain of AndFilterOperator's - no match case": {

			filter: elemental.NewFilterComposer().
				And(
					elemental.NewFilterComposer().
						WithKey("attr1").
						Equals(true).
						Done(),
					elemental.NewFilterComposer().
						And(
							elemental.NewFilterComposer().
								WithKey("attr1").
								Equals(true).
								Done(),
							elemental.NewFilterComposer().
								And(
									elemental.NewFilterComposer().
										WithKey("attr1").
										Equals(true).
										Done(),
									elemental.NewFilterComposer().
										And(
											elemental.NewFilterComposer().
												WithKey("attr1").
												Equals(true).
												Done(),
											elemental.NewFilterComposer().
												And(
													elemental.NewFilterComposer().
														WithKey("attr1").
														Equals(false).
														Done(),
													elemental.NewFilterComposer().
														And(
															elemental.NewFilterComposer().
																WithKey("attr1").
																Equals(true).
																Done(),
															elemental.NewFilterComposer().
																And(
																	elemental.NewFilterComposer().
																		And(elemental.NewFilterComposer().
																			WithKey("attr1").
																			Equals("WILL NOT MATCH").
																			Done()).
																		Done()).
																Done()).
														Done()).
												Done()).
										Done()).
								Done()).
						Done(),
				).Done(),

			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return([]bool{
						false,
						false,
						false,
						false,
						false,
						true,
					}).
					Times(7)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should be able to handle a mix of AndFilterOperator & OrFilterOperator": {
			filter: elemental.NewFilterComposer().
				And(
					elemental.NewFilterComposer().
						Or(
							elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
								"hello",
								[]bool{
									true,
									true,
								},
								"amir",
								"cristiano",
							}).Done(),
							elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
								"hello",
								[]bool{
									true,
									true,
								},
								"amir",
								"cristiano",
							}).Done(),
							elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
								"hello",
								[]bool{
									true,
									true,
								},
								"amir",
								"cristiano",
							}).Done(),
							elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
								"hello",
								[]bool{
									true,
									true,
								},
								"amir",
								"cristiano",
							}).Done(),
							elemental.NewFilterComposer().WithKey("attr1").Equals(true).Done(),
							elemental.NewFilterComposer().WithKey("attr1").Equals(true).Done(),
							elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
								false,
								false,
								false,
								"cristiano",
							}).Done(),
						).
						Done(),
				).And(
				elemental.NewFilterComposer().
					WithKey("attr1").Equals("cristiano").
					WithKey("attr1").Equals("cristiano").
					WithKey("attr1").Equals("cristiano").
					WithKey("attr1").Equals("cristiano").
					WithKey("attr1").Equals([]interface{}{false, false, false, "cristiano"}).
					WithKey("attr1").Equals("cristiano").
					WithKey("attr1").Equals("cristiano").
					WithKey("attr1").Equals("cristiano").
					WithKey("attr1").Equals("cristiano").
					WithKey("attr1").Equals("cristiano").
					Done(),
				elemental.NewFilterComposer().
					Or(
						elemental.NewFilterComposer().WithKey("attr1").Equals(false).Done(),
						elemental.NewFilterComposer().WithKey("attr1").Equals("amir").Done(),
						elemental.NewFilterComposer().WithKey("attr1").Equals("cristiano").Done(),
					).
					Done(),
			).Done(),

			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)

				mockAS.EXPECT().
					ValueForAttribute("attr1").
					Return([]interface{}{
						false,
						false,
						false,
						"cristiano",
					}).
					AnyTimes()

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && !tc.expectedPanic {
					t.Fatalf("did not expect a panic, but test case paniced with: %+v", err)
				}
			}()

			identity := tc.mockSetupFunc(t, gomock.NewController(t))
			matched, err := elemental.MatchesFilter(identity, tc.filter)

			if (err != nil) != tc.expectedError {
				t.Errorf("\n"+
					"error expectation failued:\n"+
					"expected an error: %t\n"+
					"actual error: %+v\n",
					tc.expectedError,
					err)
			}

			if matched != tc.expectedMatch {
				t.Errorf("\n"+
					"match expectation failed:\n"+
					"expected a match: %t\n"+
					"matched occurred: %+v\n",
					tc.expectedMatch,
					matched)
			}
		})
	}
}

func TestUnsupportedComparators(t *testing.T) {

	identifiable := testmodel.NewUser()
	testAttribute := "firstName"

	tests := map[string]struct {
		filter     *elemental.Filter
		comparator string
	}{
		"greater than": {
			filter:     elemental.NewFilterComposer().WithKey(testAttribute).GreaterThan("").Done(),
			comparator: ">",
		},
		"greater than equal": {
			filter:     elemental.NewFilterComposer().WithKey(testAttribute).GreaterOrEqualThan("").Done(),
			comparator: ">=",
		},
		"lesser than": {
			filter:     elemental.NewFilterComposer().WithKey(testAttribute).LesserThan("").Done(),
			comparator: "<",
		},
		"lesser than or equal": {
			filter:     elemental.NewFilterComposer().WithKey(testAttribute).LesserOrEqualThan("").Done(),
			comparator: "<=",
		},
		"in": {
			filter:     elemental.NewFilterComposer().WithKey(testAttribute).In("a", "b").Done(),
			comparator: "in",
		},
		"not in": {
			filter:     elemental.NewFilterComposer().WithKey(testAttribute).NotIn("a", "b").Done(),
			comparator: "not in",
		},
		"contains": {
			filter:     elemental.NewFilterComposer().WithKey(testAttribute).Contains("a", "b", "c").Done(),
			comparator: "contains",
		},
		"not contains": {
			filter:     elemental.NewFilterComposer().WithKey(testAttribute).NotContains("a", "b", "c").Done(),
			comparator: "not contains",
		},
	}

	for description, test := range tests {
		t.Run(fmt.Sprintf("elemental.MatchesFilter should return an error when using unsupported comparator %s", description), func(t *testing.T) {
			matched, err := elemental.MatchesFilter(identifiable, test.filter)

			if err == nil {
				t.Fatalf("expected an error to occur when using unsupported comparator\n" +
					"Hint: if you just added support for a new comparator, you can now remove the failing test case for that comparator")
			}

			if matched {
				t.Errorf("a match should never occur when using an unsupported operator")
			}

			var me *elemental.MatcherError
			if ok := errors.As(err, &me); !ok {
				t.Fatalf("expected underlying error type to be: *elemental.MatcherError\n"+
					"actual error type was: %s\n"+
					"WARNING: this is a major breaking change as you could break client error handling logic",
					reflect.TypeOf(err))
			}

			// this verifies that the matcher error chain contains the expected error

			if !errors.Is(err, elemental.ErrUnsupportedComparator) {
				t.Errorf("expected the matcher error to contain an 'elemental.ErrUnsupportedComparator'\n"+
					"actual error type: %s\n"+
					"WARNING: this is a major breaking change as you could break client error handling logic",
					reflect.TypeOf(err),
				)
			}

			expectedErrCopy := fmt.Sprintf("elemental: %s - comparator: %q", elemental.ErrUnsupportedComparator, test.comparator)
			if me.Error() != expectedErrCopy {
				t.Errorf("expected the error copy to equal: %s\n"+
					"actual error copy: %s",
					expectedErrCopy,
					me.Error())
			}
		})
	}
}

// this unit test suite tests the functionality of the NotEqualComparator when used in conjunction with the helper
// MatchesFilter for filtering an AttributeSpecifiable using the supplied filter
func TestNotEqualComparator(t *testing.T) {

	testAttributeName := "someAttribute"
	tests := map[string]struct {
		filter        *elemental.Filter
		mockSetupFunc func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable
		expectedMatch bool
		expectedError bool
	}{
		"should return true if the attribute value does NOT match the desired value": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals("john").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}("amir"))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return false if the attribute value DOES match the desired value": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals("amir").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}("amir"))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return true if the attribute does not exist on the provided identifiable": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals("john").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nil))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},

		// deals with db.getCollection('collection').find( { invalidAttribute: { $eq: null } } )
		// in such a query, no match will ever be possible

		"should return false if the specified attribute is missing in the identifiable and the value provided is nil": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals(nil).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nil))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return true if the specified attribute (a slice) does not contain the value ": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals("Thierry Henry").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"Zinedine Zidane",
						"Dennis Bergkamp",
						"Patrick Vieira",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the specified attribute (an array) does not contain the value ": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals("Thierry Henry").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...]string{
						"Zinedine Zidane",
						"Dennis Bergkamp",
						"Patrick Vieira",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return false if the specified attribute (a slice) DOES contain the value ": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals("Thierry Henry").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"Zinedine Zidane",
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the specified attribute (an array) DOES contain the value ": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals("Thierry Henry").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...]string{
						"Zinedine Zidane",
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the specified attribute (an array) matches the value (an array) exactly": {
			filter: elemental.NewFilterComposer().

				// note: order matters here!

				WithKey(testAttributeName).
				NotEquals([...]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...]string{
						"Zinedine Zidane",
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the specified attribute (a slice) matches the value (an array) exactly": {
			filter: elemental.NewFilterComposer().

				// note: order matters here!

				WithKey(testAttributeName).
				NotEquals([...]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"Zinedine Zidane",
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the specified attribute (a slice) matches the value (a slice) exactly": {
			filter: elemental.NewFilterComposer().

				// note: order matters here!

				WithKey(testAttributeName).
				NotEquals([]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"Zinedine Zidane",
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the specified attribute (an array) matches the value (a slice) exactly": {
			filter: elemental.NewFilterComposer().

				// note: order matters here!

				WithKey(testAttributeName).
				NotEquals([]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...]string{
						"Zinedine Zidane",
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return true if the specified attribute (an array) does NOT match the value (an array) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([...]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...]string{
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Zinedine Zidane", // notice how Zinedine is now at a different index
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the specified attribute (a slice) does NOT match the value (an array) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([...]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Zinedine Zidane", // notice how Zinedine is now at a different index
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the specified attribute (an array) does NOT match the value (a slice) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...]string{
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Zinedine Zidane", // notice how Zinedine is now at a different index
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the specified attribute (a slice) does NOT match the value (a slice) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"Dennis Bergkamp",
						"Patrick Vieira",
						"Zinedine Zidane", // notice how Zinedine is now at a different index
						"Thierry Henry",
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return false if the specified attribute (a slice of slice) contains an element that matches the value (a slice) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([][]string{
						{
							"a",
						},
						{
							"b",
						},
						{
							"c",
						},
						{
							"Zinedine Zidane",
							"Dennis Bergkamp",
							"Patrick Vieira",
							"Thierry Henry",
						},
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the specified attribute (a slice of slice) contains an element that matches the value (an array) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([...]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([][]string{
						{
							"a",
						},
						{
							"b",
						},
						{
							"c",
						},
						{
							"Zinedine Zidane",
							"Dennis Bergkamp",
							"Patrick Vieira",
							"Thierry Henry",
						},
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the specified attribute (an array of slice) contains an element that matches the value (a slice) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...][]string{
						{
							"a",
						},
						{
							"b",
						},
						{
							"c",
						},
						{
							"Zinedine Zidane",
							"Dennis Bergkamp",
							"Patrick Vieira",
							"Thierry Henry",
						},
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the specified attribute (an array of slice) contains an element that matches the value (an array) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([...]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...][]string{
						{
							"a",
						},
						{
							"b",
						},
						{
							"c",
						},
						{
							"Zinedine Zidane",
							"Dennis Bergkamp",
							"Patrick Vieira",
							"Thierry Henry",
						},
					}))

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return true if the specified attribute (a slice of slice) does NOT contain an element that matches the value (a slice) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([][]string{
						{
							"a",
						},
						{
							"b",
						},
						{
							"c",
						},
						{
							"Dennis Bergkamp",
							"Zinedine Zidane", // notice the index of zinedine is different on the attribute
							"Patrick Vieira",
							"Thierry Henry",
						},
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the specified attribute (a slice of slice) does NOT contain an element that matches the value (an array) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([...]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([][]string{
						{
							"a",
						},
						{
							"b",
						},
						{
							"c",
						},
						{
							"Dennis Bergkamp",
							"Zinedine Zidane", // notice the index of zinedine is different on the attribute
							"Patrick Vieira",
							"Thierry Henry",
						},
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return false if the specified attribute (an array of slice) does NOT contain an element that matches the value (a slice) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...][]string{
						{
							"a",
						},
						{
							"b",
						},
						{
							"c",
						},
						{
							"Dennis Bergkamp",
							"Zinedine Zidane", // notice the index of zinedine is different on the attribute
							"Patrick Vieira",
							"Thierry Henry",
						},
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the specified attribute (an array of slice) does NOT contain an element that matches the value (an array) exactly": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotEquals([...]string{
					"Zinedine Zidane",
					"Dennis Bergkamp",
					"Patrick Vieira",
					"Thierry Henry",
				}).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([...][]string{
						{
							"a",
						},
						{
							"b",
						},
						{
							"c",
						},
						{
							"Dennis Bergkamp",
							"Zinedine Zidane", // notice the index of zinedine is different on the attribute
							"Patrick Vieira",
							"Thierry Henry",
						},
					}))

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
	}

	for description, tc := range tests {
		t.Run(description, func(t *testing.T) {

			identity := tc.mockSetupFunc(t, gomock.NewController(t))
			matched, err := elemental.MatchesFilter(identity, tc.filter)

			if (err != nil) != tc.expectedError {
				t.Errorf("\n"+
					"error expectation failued:\n"+
					"expected an error: %t\n"+
					"actual error: %+v\n",
					tc.expectedError,
					err)
			}

			if matched != tc.expectedMatch {
				t.Errorf("\n"+
					"match expectation failed:\n"+
					"expected a match: %t\n"+
					"matched occurred: %+v\n",
					tc.expectedMatch,
					matched)
			}
		})
	}
}

// this unit test suite tests the functionality of the ExistsComparator when used in conjunction with the helper
// MatchesFilter for filtering an AttributeSpecifiable using the supplied filter
func TestExistsComparator(t *testing.T) {

	// nil fixtures
	var (
		nilPointer   *interface{}          = nil
		nilMap       map[struct{}]struct{} = nil
		nilFunc      func()                = nil
		nilChan      chan struct{}         = nil
		nilInterface interface{}           = nil
		nilSlice     []interface{}         = nil
	)

	testAttributeName := "someAttribute"
	tests := map[string]struct {
		filter        *elemental.Filter
		mockSetupFunc func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable
		expectedMatch bool
		expectedError bool
	}{
		"should return false if the attribute does not exist on the identifiable": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Exists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nil)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().

					// IMPORTANT: notice how this is returning a map which does not contain the attribute specified in the filter

					Return(map[string]elemental.AttributeSpecification{
						"someOtherAttribute": {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return true if the attribute exists, but is a nil pointer": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Exists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilPointer)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the attribute exists, but is a nil map": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Exists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilMap)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the attribute exists, but is a nil func": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Exists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilFunc)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the attribute exists, but is a nil channel": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Exists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilChan)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the attribute exists, but is a nil interface": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Exists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					// nolint - i want to be explicit about the fact that ValueForAttribute is returning an interface{}
					Return(interface{}(nilInterface)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return true if the attribute exists, but is a nil slice": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Exists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilSlice)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
	}

	for description, tc := range tests {
		t.Run(description, func(t *testing.T) {

			identity := tc.mockSetupFunc(t, gomock.NewController(t))
			matched, err := elemental.MatchesFilter(identity, tc.filter)

			if (err != nil) != tc.expectedError {
				t.Errorf("\n"+
					"error expectation failued:\n"+
					"expected an error: %t\n"+
					"actual error: %+v\n",
					tc.expectedError,
					err)
			}

			if matched != tc.expectedMatch {
				t.Errorf("\n"+
					"match expectation failed:\n"+
					"expected a match: %t\n"+
					"matched occurred: %+v\n",
					tc.expectedMatch,
					matched)
			}
		})
	}
}

// this unit test suite tests the functionality of the NotExists when used in conjunction with the helper
// MatchesFilter for filtering an AttributeSpecifiable using the supplied filter
func TestNotExistsComparator(t *testing.T) {

	// nil fixtures
	var (
		nilPointer   *interface{}          = nil
		nilMap       map[struct{}]struct{} = nil
		nilFunc      func()                = nil
		nilChan      chan struct{}         = nil
		nilInterface interface{}           = nil
		nilSlice     []interface{}         = nil
	)

	testAttributeName := "someAttribute"
	tests := map[string]struct {
		filter        *elemental.Filter
		mockSetupFunc func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable
		expectedMatch bool
		expectedError bool
	}{
		"should return true if the attribute does not exist on the identifiable": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotExists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nil)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().

					// IMPORTANT: notice how this is returning a map which does not contain the attribute specified in the filter

					Return(map[string]elemental.AttributeSpecification{
						"someOtherAttribute": {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return false if the attribute exists, but is a nil pointer": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotExists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilPointer)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the attribute exists, but is a nil map": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotExists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilMap)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the attribute exists, but is a nil func": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotExists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilFunc)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the attribute exists, but is a nil channel": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotExists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilChan)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the attribute exists, but is a nil interface": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotExists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					// nolint - i want to be explicit about the fact that ValueForAttribute is returning an interface{}
					Return(interface{}(nilInterface)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the attribute exists, but is a nil slice": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				NotExists().
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nilSlice)).
					Times(1)

				mockAS.
					EXPECT().
					AttributeSpecifications().
					Return(map[string]elemental.AttributeSpecification{
						testAttributeName: {},
					}).
					Times(1)

				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
	}

	for description, tc := range tests {
		t.Run(description, func(t *testing.T) {

			identity := tc.mockSetupFunc(t, gomock.NewController(t))
			matched, err := elemental.MatchesFilter(identity, tc.filter)

			if (err != nil) != tc.expectedError {
				t.Errorf("\n"+
					"error expectation failued:\n"+
					"expected an error: %t\n"+
					"actual error: %+v\n",
					tc.expectedError,
					err)
			}

			if matched != tc.expectedMatch {
				t.Errorf("\n"+
					"match expectation failed:\n"+
					"expected a match: %t\n"+
					"matched occurred: %+v\n",
					tc.expectedMatch,
					matched)
			}
		})
	}
}

// this unit test suite tests the functionality of the MatchComparator comparator when used in conjunction with the helper
// MatchesFilter for filtering an AttributeSpecifiable using a supplied filter
func TestMatchComparator(t *testing.T) {

	// some type that has a string has its underlying type
	type someCustomStringType string
	testAttributeName := "someAttribute"
	tests := map[string]struct {
		filter        *elemental.Filter
		mockSetupFunc func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable
		expectedMatch bool
		expectedError bool
	}{
		"should return true if the attribute (a string) matches the pattern supplied to the comparator": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches("Vancouver").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}("It rains a lot in Vancouver")).
					Times(1)
				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should return false if the attribute (a string) does not match the pattern supplied to the comparator": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches("Vancouver").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}("It's sunny in Phuket")).
					Times(1)
				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should support pattern matching the elements of an attribute that is a slice using the supplied pattern (match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches(`\$identity=.*`).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"$identity=automation",
						"$name=Test Automation",
						"$namespace=/account-9bb41b0a-8703-40ca-b3e6-975d727fcba7",
						"$id=5d5208a1339d6456750010b5",
					})).
					Times(1)
				return mockAS
			},

			// expecting a match because the attribute contains an element ($identity=automation) matching the pattern provided (^\\$identity=.*)

			expectedMatch: true,
			expectedError: false,
		},
		"should support pattern matching the elements of an attribute that is a slice using the supplied pattern (no match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches("WILL NOT MATCH ANYTHING").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"$identity=automation",
						"$name=Test Automation",
						"$namespace=/account-9bb41b0a-8703-40ca-b3e6-975d727fcba7",
						"$id=5d5208a1339d6456750010b5",
					})).
					Times(1)
				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should support pattern matching the elements of an attribute that is an array using the supplied pattern (match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches(`\$identity=.*`).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([4]string{
						"$identity=automation",
						"$name=Test Automation",
						"$namespace=/account-9bb41b0a-8703-40ca-b3e6-975d727fcba7",
						"$id=5d5208a1339d6456750010b5",
					})).
					Times(1)
				return mockAS
			},

			// expecting a match because the attribute contains an element ($identity=automation) matching the pattern provided (^\\$identity=.*)

			expectedMatch: true,
			expectedError: false,
		},
		"should support pattern matching the elements of an attribute that is an array using the supplied pattern (no match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches("WILL NOT MATCH ANYTHING").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([4]string{
						"$identity=automation",
						"$name=Test Automation",
						"$namespace=/account-9bb41b0a-8703-40ca-b3e6-975d727fcba7",
						"$id=5d5208a1339d6456750010b5",
					})).
					Times(1)
				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if none of the attribute elements is a string or is type based off a string": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches("yolo").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]interface{}{
						123,
						[]string{"yolo"},
					})).
					Times(1)
				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false is the attribute being matched on is NOT a string": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches("123456789").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(123456789)).
					Times(1)
				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if no patterns have been provided to the comparator to match on": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).

				// notice how this is being passed nil

				Matches(nil).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"$identity=automation",
						"$name=Test Automation",
						"$namespace=/account-9bb41b0a-8703-40ca-b3e6-975d727fcba7",
						"$id=5d5208a1339d6456750010b5",
					})).
					Times(1)
				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should return false if the attribute being matched on does not exist on the identifiable": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches(".*").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}(nil)).
					Times(1)
				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should ignore non-string patterns passed into the match comparator parameters (match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches(
					// nope, not a string, should be skipped
					1234,
					// nope, not a string, should be skipped
					[]string{"yolo"},
					// this should result in a match on ("$identity=automation")
					`^\$identity=.*`,
				).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}("$identity=policy")).
					Times(1)
				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should ignore non-string patterns passed into the match comparator parameters (no match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches(
					// nope, not a string, should be skipped
					1234,
					// nope, not a string, should be skipped
					[]string{"yolo"},
					// nope, not a string, should be skipped
					interface{}(123),
				).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}("yolo")).
					Times(1)
				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
		"should support matching on attribute types that have a string as their underlying type (match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches(".*amir$").
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).

					// notice the type of the attribute - someCustomStringType - which has a string as its underlying type!

					Return(interface{}(
						someCustomStringType("###############amir"),
					)).
					Times(1)
				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},

		// in this test, we verify the behaviour that for each pattern provided, it should be matched against each of the
		// attribute values to determine if a match is possible. this also verifies that we are dealing with OR semantics because
		// a match is found as long as just one of the provided patterns matches

		"should support matching on multiple patterns for an attribute that is a slice (match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches(
					".*somepattern$",
					"nope$",
					// this should match the last attribute element! ($id=5d5208a1339d6456750010b5)
					`^\$id=5d5208.*`).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"$identity=automation",
						"$name=Test Automation",
						"$namespace=/account-9bb41b0a-8703-40ca-b3e6-975d727fcba7",
						"$id=5d5208a1339d6456750010b5",
					})).
					Times(1)
				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should ignore regex patterns that result in errors (match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches(
					// missing argument to repetition operator: `*`
					"*",
					// missing argument to repetition operator: `+`
					"+",
					// unexpected ): `abc)`
					"abc)",
					// this is a valid regex and should match the first attribute element ($identity=automation)
					`^\$identity=`,
				).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"$identity=automation",
						"$name=Test Automation",
						"$namespace=/account-9bb41b0a-8703-40ca-b3e6-975d727fcba7",
						"$id=5d5208a1339d6456750010b5",
					})).
					Times(1)
				return mockAS
			},
			expectedMatch: true,
			expectedError: false,
		},
		"should ignore regex patterns that result in errors (no match case)": {
			filter: elemental.NewFilterComposer().
				WithKey(testAttributeName).
				Matches(
					// missing argument to repetition operator: `*`
					"*",
					// missing argument to repetition operator: `+`
					"+",
					// unexpected ): `abc)`
					"abc)",
				).
				Done(),
			mockSetupFunc: func(t *testing.T, ctrl *gomock.Controller) elemental.AttributeSpecifiable {
				t.Helper()
				mockAS := internal.NewMockAttributeSpecifiable(ctrl)
				mockAS.
					EXPECT().
					ValueForAttribute(testAttributeName).
					Return(interface{}([]string{
						"$identity=automation",
						"$name=Test Automation",
						"$namespace=/account-9bb41b0a-8703-40ca-b3e6-975d727fcba7",
						"$id=5d5208a1339d6456750010b5",
					})).
					Times(1)
				return mockAS
			},
			expectedMatch: false,
			expectedError: false,
		},
	}

	for description, tc := range tests {
		t.Run(description, func(t *testing.T) {

			identity := tc.mockSetupFunc(t, gomock.NewController(t))
			matched, err := elemental.MatchesFilter(identity, tc.filter)

			if (err != nil) != tc.expectedError {
				t.Errorf("\n"+
					"error expectation failued:\n"+
					"expected an error: %t\n"+
					"actual error: %+v\n",
					tc.expectedError,
					err)
			}

			if matched != tc.expectedMatch {
				t.Errorf("\n"+
					"match expectation failed:\n"+
					"expected a match: %t\n"+
					"matched occurred: %+v\n",
					tc.expectedMatch,
					matched)
			}
		})
	}
}

//func TestMatcherError_String(t *testing.T) {
//
//	tests := map[string]struct {
//		kind elemental.MatcherErrorKind
//		want string
//	}{
//		"KindUnsupportedComparator": {
//			kind: elemental.KindUnsupportedComparator,
//			want: "KindUnsupportedComparator",
//		},
//		"UnknownMatcherErrorKind": {
//			kind: elemental.MatcherErrorKind(-1),
//			want: "UnknownMatcherErrorKind",
//		},
//	}
//
//	for scenario, test := range tests {
//		t.Run(scenario, func(t *testing.T) {
//			if actual := test.kind.String(); actual != test.want {
//				t.Errorf("expected: %s\n"+
//					"got: %s",
//					test.want,
//					actual)
//			}
//		})
//	}
//}

type benchFixture struct {
}

func (f *benchFixture) SpecificationForAttribute(string) elemental.AttributeSpecification {
	return elemental.AttributeSpecification{}
}

func (f *benchFixture) AttributeSpecifications() map[string]elemental.AttributeSpecification {
	return nil
}

func (f *benchFixture) ValueForAttribute(name string) interface{} {
	switch name {
	case "attr1":
		return []interface{}{
			false,
			false,
			false,
			"cristiano",
		}
	}

	return nil
}

// this is to avoid compiler optimization which may ruin the benchmark
var matchResult bool

func BenchmarkMatchesFilterStress(b *testing.B) {
	b.ReportAllocs()

	testFixture := &benchFixture{}
	testFilter := elemental.NewFilterComposer().
		And(
			elemental.NewFilterComposer().
				Or(
					elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
						"hello",
						[]bool{
							true,
							true,
						},
						"amir",
						"cristiano",
					}).Done(),
					elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
						"hello",
						[]bool{
							true,
							true,
						},
						"amir",
						"cristiano",
					}).Done(),
					elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
						"hello",
						[]bool{
							true,
							true,
						},
						"amir",
						"cristiano",
					}).Done(),
					elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
						"hello",
						[]bool{
							true,
							true,
						},
						"amir",
						"cristiano",
					}).Done(),
					elemental.NewFilterComposer().WithKey("attr1").Equals(true).Done(),
					elemental.NewFilterComposer().WithKey("attr1").Equals(true).Done(),
					elemental.NewFilterComposer().WithKey("attr1").Equals([]interface{}{
						false,
						false,
						false,
						"cristiano",
					}).Done(),
				).
				Done(),
		).And(
		elemental.NewFilterComposer().
			WithKey("attr1").Equals("cristiano").
			WithKey("attr1").Equals("cristiano").
			WithKey("attr1").Equals("cristiano").
			WithKey("attr1").Equals("cristiano").
			WithKey("attr1").Equals([]interface{}{false, false, false, "cristiano"}).
			WithKey("attr1").Equals("cristiano").
			WithKey("attr1").Equals("cristiano").
			WithKey("attr1").Equals("cristiano").
			WithKey("attr1").Equals("cristiano").
			WithKey("attr1").Equals("cristiano").
			Done(),
		elemental.NewFilterComposer().
			Or(
				elemental.NewFilterComposer().WithKey("attr1").Equals(false).Done(),
				elemental.NewFilterComposer().WithKey("attr1").Equals("amir").Done(),
				elemental.NewFilterComposer().WithKey("attr1").Equals("cristiano").Done(),
			).
			Done(),
	).Done()

	var matched bool
	var err error
	for n := 0; n < b.N; n++ {
		// recording the result so the compiler doesn't avoid the function call
		matched, err = elemental.MatchesFilter(testFixture, testFilter)
		if err != nil {
			b.Errorf("benchmark test case should not have resulted in an error - %s", err)
		}

		if !matched {
			b.Error("benchmark test case should have resulted in a match")
		}
	}

	matchResult = matched
}
