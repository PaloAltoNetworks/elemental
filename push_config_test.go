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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
