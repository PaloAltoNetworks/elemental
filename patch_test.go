// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type attrSpecifiable struct {
	unexported string
}

func (attrSpecifiable) SpecificationForAttribute(string) AttributeSpecification {
	return AttributeSpecification{
		ConvertedName: "unexported",
	}
}
func (attrSpecifiable) AttributeSpecifications() map[string]AttributeSpecification {
	return nil
}

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

func TestPatch_Apply(t *testing.T) {

	Convey("Given I have a list and a patch", t, func() {

		now := time.Now()

		l1 := NewList()
		l1.CreationOnly = "not-patched"
		l1.Date = now
		l1.Description = "not-patched"
		l1.ID = "not-patched"
		l1.Name = "not-patched"
		l1.ParentID = "not-patched"
		l1.ParentType = "not-patched"
		l1.ReadOnly = "not-patched"
		l1.Slice = []string{"not-patched", "not-patched"}
		l1.Unexposed = "not-patched"

		Convey("When I apply the patch on name only", func() {

			err := NewPatch(0, PatchData{"name": "patched"}).Apply(l1)

			Convey("Then err should be bil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the patch should be applied", func() {
				So(l1.CreationOnly, ShouldEqual, "not-patched")
				So(l1.Date, ShouldResemble, now)
				So(l1.Description, ShouldEqual, "not-patched")
				So(l1.ID, ShouldEqual, "not-patched")
				So(l1.Name, ShouldEqual, "patched")
				So(l1.ParentID, ShouldEqual, "not-patched")
				So(l1.ParentType, ShouldEqual, "not-patched")
				So(l1.ReadOnly, ShouldEqual, "not-patched")
				So(l1.Slice, ShouldResemble, []string{"not-patched", "not-patched"})
				So(l1.Unexposed, ShouldResemble, "not-patched")
			})
		})

		Convey("When I apply the patch on everything", func() {

			err := NewPatch(0, PatchData{
				"creationOnly": "patched",
				"date":         time.Now(),
				"description":  "patched",
				"id":           "patched",
				"name":         "patched",
				"parentID":     "patched",
				"parentType":   "patched",
				"readOnly":     "patched",
				"slice":        []string{"patched"},
				"unexposed":    "patched",
			}).Apply(l1)

			Convey("Then err should be bil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the patch should be applied", func() {
				So(l1.CreationOnly, ShouldEqual, "patched")
				So(l1.Date, ShouldNotEqual, now)
				So(l1.Description, ShouldEqual, "patched")
				So(l1.ID, ShouldEqual, "patched")
				So(l1.Name, ShouldEqual, "patched")
				So(l1.ParentID, ShouldEqual, "patched")
				So(l1.ParentType, ShouldEqual, "patched")
				So(l1.ReadOnly, ShouldEqual, "patched")
				So(l1.Slice, ShouldResemble, []string{"patched"})
				So(l1.Unexposed, ShouldResemble, "patched")
			})
		})

		Convey("When I apply the patch on a field that does not exist", func() {

			err := NewPatch(0, PatchData{
				"woops": "patched",
			}).Apply(l1)

			Convey("Then err should be correct", func() {
				So(err.Error(), ShouldEqual, "error 422 (elemental): Validation Error: field 'woops' is invalid")
			})
		})

		Convey("When I apply the patch on a object that is not a pointer", func() {
			Convey("Then it should should panic", func() {
				So(func() { _ = NewPatch(0, PatchData{}).Apply(attrSpecifiable{}) }, ShouldPanicWith, "A pointer to elemental.AttributeSpecifiable must be passed to Apply")
			})
		})

		Convey("When I apply the patch on a object on a field that is not settable", func() {

			err := NewPatch(0, PatchData{
				"unexported": "patched",
			}).Apply(&attrSpecifiable{})

			Convey("Then err should be correct", func() {
				So(err.Error(), ShouldEqual, "error 422 (elemental): Validation Error: field 'unexported' cannot be set")
			})
		})
	})
}
