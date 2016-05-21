// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRegistry_Registration(t *testing.T) {

	Convey("Given I have an Identity", t, func() {

		identity := MakeIdentity("parent", "parents")

		Convey("When I try to get it from name  before registering it", func() {

			i := IdentityFromName("parent")

			Convey("Then the result should be empty", func() {
				So(i.IsEmpty(), ShouldBeTrue)
			})

			Convey("Then the result should not be equal to identity", func() {
				So(i.IsEqual(identity), ShouldBeFalse)
			})
		})

		Convey("When I try to get from name it after registering it", func() {

			RegisterIdentity(identity)
			i := IdentityFromName("parent")

			Convey("Then the result should not be empty", func() {
				So(i.IsEmpty(), ShouldBeFalse)
			})

			Convey("Then the result should be equal to identity", func() {
				So(i.IsEqual(identity), ShouldBeTrue)
			})

			Convey("When I try to unregister it", func() {

				UnregisterIdentity(identity)
				i := IdentityFromName("parent")

				Convey("Then the result should be empty", func() {
					So(i.IsEmpty(), ShouldBeTrue)
				})

				Convey("Then the result should not be equal to identity", func() {
					So(i.IsEqual(identity), ShouldBeFalse)
				})
			})
		})

		Convey("When I try to get it from category  before registering it", func() {

			i := IdentityFromCategory("parents")

			Convey("Then the result should be empty", func() {
				So(i.IsEmpty(), ShouldBeTrue)
			})

			Convey("Then the result should not be equal to identity", func() {
				So(i.IsEqual(identity), ShouldBeFalse)
			})
		})

		Convey("When I try to get from category it after registering it", func() {

			RegisterIdentity(identity)
			i := IdentityFromCategory("parents")

			Convey("Then the result should not be empty", func() {
				So(i.IsEmpty(), ShouldBeFalse)
			})

			Convey("Then the result should be equal to identity", func() {
				So(i.IsEqual(identity), ShouldBeTrue)
			})
		})

		Convey("When I try to unregister it", func() {

			UnregisterIdentity(identity)
			i := IdentityFromCategory("parents")

			Convey("Then the result should be empty", func() {
				So(i.IsEmpty(), ShouldBeTrue)
			})

			Convey("Then the result should not be equal to identity", func() {
				So(i.IsEqual(identity), ShouldBeFalse)
			})
		})
	})
}
