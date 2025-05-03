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

package elemental

import (
	"net/url"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPushConfig_FilterForIdentity(t *testing.T) {

	tests := map[string]struct {
		pushConfig     *PushConfig
		identity       string
		expectedFilter *Filter
		expectedFound  bool
	}{
		"should return the identity filter for the given identity and true if it exists": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": "namespace == /liverpool-fc and environment == production",
				},
			},
			identity: "identity_one",
			expectedFilter: NewFilterComposer().And(
				NewFilterComposer().WithKey("namespace").Equals("/liverpool-fc").Done(),
				NewFilterComposer().WithKey("environment").Equals("production").Done(),
			).Done(),
			expectedFound: true,
		},
		"should return a nil filter and false if the push config has no identity filters configured": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
			},
			identity:       "identity_one",
			expectedFilter: nil,
			expectedFound:  false,
		},
		"should return a nil filter and false if the push config has no identity filter for the provided identity": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": "namespace == /liverpool-fc and environment == production",
				},
			},
			identity:       "some_other_identity",
			expectedFilter: nil,
			expectedFound:  false,
		},
	}

	for scenario, testCase := range tests {
		t.Run(scenario, func(t *testing.T) {
			if len(testCase.pushConfig.IdentityFilters) > 0 {
				if err := testCase.pushConfig.ParseIdentityFilters(); err != nil {
					t.Fatalf("test setup invalid - failed to parse identity filters for the configured push config: %+v", err)
				}
			}

			filter, found := testCase.pushConfig.FilterForIdentity(testCase.identity)

			if found != testCase.expectedFound {
				t.Errorf("expectation failed\n"+
					"expected identity to be found: %t\n"+
					"actual: %t\n",
					testCase.expectedFound,
					found)
			}

			if !reflect.DeepEqual(filter, testCase.expectedFilter) {
				t.Errorf("returned filter does not match expected filter\n"+
					"expected: %+v\n"+
					"actual: %+v\n",
					testCase.expectedFilter,
					filter)
			}
		})
	}
}

func TestPushConfig_ParseIdentityFilters(t *testing.T) {

	tests := map[string]struct {
		pushConfig      *PushConfig
		expectedFilters map[string]*Filter
		expectedError   bool
	}{
		"should successfully parse the filters and populate the parsed filters attribute of push config": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
					"identity_two": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": "namespace == /liverpool-fc and environment == production",
					"identity_two": "namespace == /barcelona-fc and environment == staging",
				},
			},
			expectedFilters: map[string]*Filter{
				"identity_one": NewFilterComposer().And(
					NewFilterComposer().WithKey("namespace").Equals("/liverpool-fc").Done(),
					NewFilterComposer().WithKey("environment").Equals("production").Done(),
				).Done(),
				"identity_two": NewFilterComposer().And(
					NewFilterComposer().WithKey("namespace").Equals("/barcelona-fc").Done(),
					NewFilterComposer().WithKey("environment").Equals("staging").Done(),
				).Done(),
			},
			expectedError: false,
		},
		"should return an error in the event that an identity filter is being defined on an identity that is NOT declared in the identities attribute": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{},
				IdentityFilters: map[string]string{
					"identity_one": "namespace == /liverpool-fc and environment == production",
				},
			},
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should return an error in the event that an identity filter could not be parsed": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					// notice how this will result in a parsing error due to an invalid comparator "======="
					"identity_four": "namespace ======= /liverpool-fc",
				},
			},
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should zero the parsed identity filters attribute in the event that an error occurs due to an undeclared identity": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
					"identity_two": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
					"identity_three": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
					"identity_four": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					// notice how this is an undeclared identity as 'identity_five' is not declared in the PushConfig's
					// 'Identities' attribute
					"identity_five": "namespace == /liverpool-fc and environment == production",
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should zero the parsed identity filters attribute in the event that an error occurs due to a parsing error": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
					"identity_two": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
					"identity_three": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
					"identity_four": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					// notice how this will result in a parsing error due to an invalid comparator "======="
					"identity_four": "namespace ======= /liverpool-fc",
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should result in error if an identity filter uses the unsupported comparator '>'": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": NewFilterComposer().
						WithKey("someAttr").
						GreaterThan(1).
						Done().
						String(),
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should result in error if an identity filter uses the unsupported comparator '>='": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": NewFilterComposer().
						WithKey("someAttr").
						GreaterOrEqualThan(1).
						Done().
						String(),
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should result in error if an identity filter uses the unsupported comparator '<'": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": NewFilterComposer().
						WithKey("someAttr").
						LesserThan(1).
						Done().
						String(),
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should result in error if an identity filter uses the unsupported comparator '<='": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": NewFilterComposer().
						WithKey("someAttr").
						LesserOrEqualThan(1).
						Done().
						String(),
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should result in error if an identity filter uses the unsupported comparator 'in": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": NewFilterComposer().
						WithKey("someAttr").
						In(1, 2, 3).
						Done().
						String(),
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should result in error if an identity filter uses the unsupported comparator 'not in": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": NewFilterComposer().
						WithKey("someAttr").
						NotIn(1, 2, 3).
						Done().
						String(),
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should result in error if an identity filter uses the unsupported comparator 'contains": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": NewFilterComposer().
						WithKey("someAttr").
						Contains(1, 2, 3).
						Done().
						String(),
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
		"should result in error if an identity filter uses the unsupported comparator 'not contains": {
			pushConfig: &PushConfig{
				Identities: map[string][]EventType{
					"identity_one": {
						EventCreate,
						EventUpdate,
						EventDelete,
					},
				},
				IdentityFilters: map[string]string{
					"identity_one": NewFilterComposer().
						WithKey("someAttr").
						NotContains(1, 2, 3).
						Done().
						String(),
				},
			},
			// the push config's parsed identities attribute should be zero'd out
			expectedFilters: map[string]*Filter{},
			expectedError:   true,
		},
	}

	for scenario, testCase := range tests {
		t.Run(scenario, func(t *testing.T) {
			err := testCase.pushConfig.ParseIdentityFilters()

			if (err != nil) != testCase.expectedError {
				t.Errorf("\n"+
					"error expectation failed\n"+
					"test case expected an error: %t\n"+
					"an error actually occur: %t\n"+
					"actual error: %+v\n",
					testCase.expectedError,
					err != nil,
					err,
				)
			}

			if !reflect.DeepEqual(testCase.pushConfig.parsedIdentityFilters, testCase.expectedFilters) {
				t.Errorf("the parsed filters did not match what was expected\n"+
					"expected: %+v\n"+
					"actual: %+v\n",
					testCase.expectedFilters,
					testCase.pushConfig.parsedIdentityFilters)
			}
		})
	}
}

func TestPushConfig_NewPushConfig(t *testing.T) {

	Convey("Given I create a new PushConfig", t, func() {

		f := NewPushConfig()

		Convey("Then it should be correctly initialized", func() {
			So(f.Identities, ShouldNotBeNil)
			So(f.IdentityFilters, ShouldNotBeNil)
			So(f.parsedIdentityFilters, ShouldNotBeNil)
			So(f.Params, ShouldBeNil)
		})
	})
}

// just keeping this for backwards compatibility so we don't break the API by accident by removing the old constructor API
func TestPushConfig_NewPushFilter(t *testing.T) {

	Convey("Given I create a new NewPushFilter", t, func() {

		f := NewPushFilter()

		Convey("Then it should be correctly initialized", func() {
			So(f.Identities, ShouldNotBeNil)
			So(f.IdentityFilters, ShouldNotBeNil)
			So(f.parsedIdentityFilters, ShouldNotBeNil)
			So(f.Params, ShouldBeNil)
		})
	})
}

func TestPushConfig_Duplicate(t *testing.T) {

	Convey("Given I create a new PushConfig", t, func() {

		f := NewPushConfig()

		f.SetParameter("key", "values")

		f.FilterIdentity("i1", EventCreate, EventDelete)
		f.FilterIdentity("i2", EventCreate, EventDelete)

		testFilter := NewFilterComposer().
			And(
				NewFilterComposer().
					WithKey("propertyA").
					Equals("someValue").
					Done(),
				NewFilterComposer().
					WithKey("propertyB").
					Equals("someValue").
					Done(),
			).Done()

		f.IdentityFilters = map[string]string{
			"i1": testFilter.String(),
		}

		// parse the identity filters so the private attribute 'IdentityFilters' gets populated
		So(f.ParseIdentityFilters(), ShouldBeNil)

		Convey("When I call Duplicate", func() {

			dup := f.Duplicate()

			Convey("Then it should be correctly duplicated", func() {
				So(dup.Identities, ShouldResemble, f.Identities)
				So(dup.Identities, ShouldNotEqual, f.Identities)

				So(dup.IdentityFilters, ShouldResemble, f.IdentityFilters)
				So(dup.IdentityFilters, ShouldNotEqual, f.IdentityFilters)

				So(dup.parsedIdentityFilters, ShouldResemble, f.parsedIdentityFilters)
				So(dup.parsedIdentityFilters, ShouldNotEqual, f.parsedIdentityFilters)

				So(dup.Params, ShouldResemble, f.Params)
				So(dup.Params, ShouldNotEqual, f.Params)
			})
		})
	})
}

func TestPushConfig_Parameters(t *testing.T) {

	Convey("Given I create a new PushConfig", t, func() {

		f := NewPushConfig()

		Convey("When I call SetParameter", func() {

			f.SetParameter("key1", "v1", "v2")
			f.SetParameter("key2", "v3")

			Convey("Then the parameter should be set", func() {

				So(f.Parameters(), ShouldResemble, url.Values{
					"key1": []string{"v1", "v2"},
					"key2": []string{"v3"},
				})

				So(f.Parameters(), ShouldNotEqual, f.Params)
			})
		})
	})

	Convey("Given I have a push filter with no parameters", t, func() {

		f := NewPushConfig()

		Convey("When I call Parameters", func() {

			p := f.Parameters()

			Convey("Then p should be nil", func() {
				So(p, ShouldBeNil)
			})
		})
	})
}

func TestPushConfig_IsFilteredOut(t *testing.T) {

	Convey("Given I create a new PushConfig", t, func() {

		f := NewPushConfig()

		Convey("When I check if i1 is filtered with a nil value for identities", func() {

			f.Identities = nil

			filtered1 := f.IsFilteredOut("i1", EventDelete)
			filtered2 := f.IsFilteredOut("i2", EventDelete)

			Convey("Then filtered1 should be false", func() {
				So(filtered1, ShouldBeFalse)
			})

			Convey("Then filtered2 should be false", func() {
				So(filtered2, ShouldBeFalse)
			})
		})

		Convey("When I check if i1 is filtered with an empty identities list", func() {

			filtered1 := f.IsFilteredOut("i1", EventDelete)
			filtered2 := f.IsFilteredOut("i2", EventDelete)

			Convey("Then filtered1 should be false", func() {
				So(filtered1, ShouldBeFalse)
			})

			Convey("Then filtered2 should be false", func() {
				So(filtered2, ShouldBeFalse)
			})
		})

		Convey("When I check if i1 is filtered", func() {

			f.FilterIdentity("i1")

			filtered1 := f.IsFilteredOut("i1", EventDelete)
			filtered2 := f.IsFilteredOut("i2", EventDelete)

			Convey("Then filtered1 should be false", func() {
				So(filtered1, ShouldBeFalse)
			})

			Convey("Then filtered2 should be false", func() {
				So(filtered2, ShouldBeTrue)
			})
		})

		Convey("When I add a filter for i1 on Create and Delete", func() {

			f.FilterIdentity("i1", EventCreate, EventDelete)
			f.FilterIdentity("i2")

			Convey("Then create and delete should not be filtered out on i1", func() {
				So(f.IsFilteredOut("i1", EventCreate), ShouldBeFalse)
				So(f.IsFilteredOut("i1", EventDelete), ShouldBeFalse)
			})

			Convey("Then update should be filtered out on i1", func() {
				So(f.IsFilteredOut("i1", EventUpdate), ShouldBeTrue)
			})

			Convey("Then nothing should be filtered out on i2", func() {
				So(f.IsFilteredOut("i2", EventCreate), ShouldBeFalse)
				So(f.IsFilteredOut("i2", EventUpdate), ShouldBeFalse)
				So(f.IsFilteredOut("i2", EventDelete), ShouldBeFalse)
			})
		})

		Convey("When I add a filter for i1 on nothing", func() {

			f.FilterIdentity("i1")
			f.FilterIdentity("i2")

			Convey("Then everything should not be filtered out on i1", func() {
				So(f.IsFilteredOut("i1", EventCreate), ShouldBeFalse)
				So(f.IsFilteredOut("i1", EventDelete), ShouldBeFalse)
				So(f.IsFilteredOut("i1", EventUpdate), ShouldBeFalse)
			})

			Convey("Then nothing should be filtered out on i2", func() {
				So(f.IsFilteredOut("i2", EventCreate), ShouldBeFalse)
				So(f.IsFilteredOut("i2", EventUpdate), ShouldBeFalse)
				So(f.IsFilteredOut("i2", EventDelete), ShouldBeFalse)
			})
		})
	})
}

func TestPushConfig_String(t *testing.T) {

	Convey("Given I create a new PushConfig", t, func() {

		f := NewPushConfig()

		f.FilterIdentity("i1", EventCreate, EventDelete)

		testFilter := NewFilterComposer().
			And(
				NewFilterComposer().
					WithKey("propertyA").
					Equals("someValue").
					Done(),
				NewFilterComposer().
					WithKey("propertyB").
					Equals("someValue").
					Done(),
			).Done()

		f.IdentityFilters = map[string]string{
			"i1": testFilter.String(),
		}

		Convey("When I call the String Method", func() {
			s := f.String()

			Convey("Then it should be correctly printed", func() {
				So(s, ShouldEqual, `<pushconfig identities:map[i1:[create delete]] identityfilters:map[i1:((propertyA == "someValue") and (propertyB == "someValue"))]>`)
			})
		})
	})
}
