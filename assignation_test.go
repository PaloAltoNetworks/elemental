// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAssignation_NewAssignation(t *testing.T) {

	Convey("Given I create a new Assignation", t, func() {

		l1 := NewList()
		l1.ID = "1"
		l2 := NewList()
		l2.ID = "2"

		a := NewAssignation(AssignationTypeSet, ListIdentity, l1, l2)

		Convey("Then the assignation should be correctly initialized", func() {
			So(a.Type, ShouldEqual, AssignationTypeSet)
			So(a.IDs, ShouldResemble, []string{"1", "2"})
			So(a.MembersIdentity, ShouldResemble, ListIdentity)
		})

		Convey("Then the identity should be InternalAssignationIdentity", func() {
			So(a.Identity(), ShouldResemble, internalAssignationIdentity)
		})

		Convey("Then the identifier should be InternalAssignationIdentity", func() {
			So(a.Identifier(), ShouldEqual, "__internal__")
		})
	})

	Convey("Given I create a new Assignation with object with no Identifier", t, func() {

		l1 := NewList()
		l1.Name = "l1"

		Convey("Then it should panic", func() {
			So(func() { NewAssignation(AssignationTypeSet, ListIdentity, l1) }, ShouldPanic)
		})
	})

	Convey("Given I create a new Assignation with a bad operation", t, func() {

		a := NewAssignation(42, ListIdentity)

		Convey("Then it not validate correctly", func() {
			So(a.Validate(), ShouldNotBeNil)
		})
	})

}

func TestAssignation_String(t *testing.T) {

	l1 := NewList()
	l1.ID = "1"

	Convey("Given I create a new Assignation with operation set", t, func() {

		a := NewAssignation(AssignationTypeSet, ListIdentity, l1)

		Convey("When I convert it to string", func() {

			s := a.String()

			Convey("Then the string should be correct", func() {
				So(s, ShouldEqual, "<Assignation type:1 identity:list ids:[1]>")
			})
		})
	})

	Convey("Given I create a new Assignation with operation substractive", t, func() {

		a := NewAssignation(AssignationTypeSubstract, ListIdentity, l1)

		Convey("When I convert it to string", func() {

			s := a.String()

			Convey("Then the string should be correct", func() {
				So(s, ShouldEqual, "<Assignation type:3 identity:list ids:[1]>")
			})
		})
	})

	Convey("Given I create a new Assignation with operation additive", t, func() {

		a := NewAssignation(AssignationTypeAdd, ListIdentity, l1)

		Convey("When I convert it to string", func() {

			s := a.String()

			Convey("Then the string should be correct", func() {
				So(s, ShouldEqual, "<Assignation type:2 identity:list ids:[1]>")
			})
		})
	})
}
