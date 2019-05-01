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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIdentity_AllIdentity(t *testing.T) {

	Convey("Given I retrieve the AllIdentity", t, func() {
		i := AllIdentity

		Convey("Then Name should *", func() {
			So(i.Name, ShouldEqual, "*")
		})

		Convey("Then Category should *", func() {
			So(i.Category, ShouldEqual, "*")
		})
	})
}

func TestIdentity_MakeIdentity(t *testing.T) {

	Convey("Given I create a new identity", t, func() {
		i := MakeIdentity("n", "c")

		Convey("Then Name should n", func() {
			So(i.Name, ShouldEqual, "n")
		})

		Convey("Then Category should c", func() {
			So(i.Category, ShouldEqual, "c")
		})
	})
}

func TestIdentity_String(t *testing.T) {

	Convey("Given I create a new identity", t, func() {
		i := MakeIdentity("n", "c")

		Convey("Then String should <Identity n|c>", func() {
			So(i.String(), ShouldEqual, "<Identity n|c>")
		})
	})
}

func TestIdentity_IsEmpty(t *testing.T) {

	Convey("Given I create a new emtpty identity", t, func() {
		i := Identity{}

		Convey("Then IsEmpty should return true", func() {
			So(i.IsEmpty(), ShouldBeTrue)
		})
	})

	Convey("Given I create a new non emtpty identity", t, func() {
		i := MakeIdentity("a", "b")

		Convey("Then IsEmpty should return false", func() {
			So(i.IsEmpty(), ShouldBeFalse)
		})
	})
}

func TestIdentity_Identity_Copy(t *testing.T) {

	Convey("Given I create I have a Identifiables with 2 Identifiable", t, func() {

		l1 := NewList()
		l1.ID = "x"

		l2 := NewList()
		l2.ID = "y"

		lst1 := ListsList{l1, l2}

		Convey("When I create copy", func() {

			lst2 := lst1.Copy()

			Convey("Then the copy should be correct", func() {
				So(len(lst1.List()), ShouldEqual, len(lst2.List()))
				So(lst1.List()[0].Identifier(), ShouldEqual, lst2.List()[0].Identifier())
				So(lst1.List()[1].Identifier(), ShouldEqual, lst2.List()[1].Identifier())
			})
		})
	})
}
