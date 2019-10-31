package elemental_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	testmodel "go.aporeto.io/elemental/test/model"

	"github.com/golang/mock/gomock"
	"go.aporeto.io/elemental"
	"go.aporeto.io/elemental/internal"
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
		filter *elemental.Filter
	}{
		"not equals": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).NotEquals("").Done(),
		},
		"greater than": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).GreaterThan("").Done(),
		},
		"greater than equal": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).GreaterOrEqualThan("").Done(),
		},
		"lesser than": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).LesserThan("").Done(),
		},
		"lesser than or equal": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).LesserOrEqualThan("").Done(),
		},
		"in": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).In("a", "b").Done(),
		},
		"not in": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).NotIn("a", "b").Done(),
		},
		"contains": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).Contains("a", "b", "c").Done(),
		},
		"not contains": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).NotContains("a", "b", "c").Done(),
		},
		"matches": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).Matches(".*").Done(),
		},
		"exists": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).Exists().Done(),
		},
		"not exists": {
			filter: elemental.NewFilterComposer().WithKey(testAttribute).NotExists().Done(),
		},
	}

	for description, test := range tests {
		t.Run(fmt.Sprintf("elemental.MatchesFilter should return an error when using unsupported comparator %s", description), func(t *testing.T) {
			matched, err := elemental.MatchesFilter(identifiable, test.filter)

			if err == nil {
				t.Fatalf("expected an error to occur when using unsupported comparator")
			}

			if err != nil {
				if !strings.Contains(err.Error(), "elemental: unsuported comparator") {
					t.Errorf("expected the error to be due to using an unsupported comparator, but it was: %s", err)
				}
			}

			if matched {
				t.Errorf("a match should never occur when using an unsupported operator")
			}
		})
	}
}

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
