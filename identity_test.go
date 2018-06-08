// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

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
