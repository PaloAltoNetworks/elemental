// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPatch_NewPatch(t *testing.T) {

	Convey("Given I create a new Patch", t, func() {

		data := PatchData{
			"name": "patched",
		}

		a := NewPatch(PatchTypeSet, data)

		Convey("Then the patch should be correctly initialized", func() {
			So(a.Type, ShouldEqual, PatchTypeSet)
			So(a.Data, ShouldResemble, data)
		})

		Convey("Then the identity should be InternalPatchIdentity", func() {
			So(a.Identity(), ShouldResemble, internalPatchIdentity)
		})

		Convey("Then the identifier should be InternalPatchIdentity", func() {
			So(a.Identifier(), ShouldEqual, "__internal__")
		})
	})

	Convey("Given I create a new Patch with a bad operation", t, func() {

		a := NewPatch(42, nil)

		Convey("Then it not validate correctly", func() {
			So(a.Validate(), ShouldNotBeNil)
		})
	})

}

func TestPatch_String(t *testing.T) {

	data := PatchData{
		"name": "patched",
	}

	Convey("Given I create a new patch with operation set", t, func() {

		a := NewPatch(PatchTypeSet, data)

		Convey("When I convert it to string", func() {

			s := a.String()

			Convey("Then the string should be correct", func() {
				So(s, ShouldEqual, "<patch type:0 data:map[name:patched]>")
			})
		})
	})

	Convey("Given I create a new patch with operation setIfZero", t, func() {

		a := NewPatch(PatchTypeSetIfZero, data)

		Convey("When I convert it to string", func() {

			s := a.String()

			Convey("Then the string should be correct", func() {
				So(s, ShouldEqual, "<patch type:1 data:map[name:patched]>")
			})
		})
	})
}
