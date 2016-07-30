package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVerify_ValidateAdvancedSpecification(t *testing.T) {

	Convey("Given I have two lists", t, func() {

		l1 := NewList()
		l2 := NewList()

		Convey("When I verify 2 objects that are ok on a create operation", func() {

			l1.ReadOnly = ""
			l1.CreationOnly = "cvalue"

			errs := ValidateAdvancedSpecification(l1, nil, OperationCreate)

			Convey("Then errs should not nil", func() {
				So(errs, ShouldBeNil)
			})
		})

		Convey("When I try to set a readonly attribute on a create operation", func() {

			l1.ReadOnly = "value"

			errs := ValidateAdvancedSpecification(l1, nil, OperationCreate)

			Convey("Then errs should not be nil", func() {
				So(errs, ShouldNotBeNil)
				So(len(errs), ShouldEqual, 1)
			})
		})

		Convey("When I try to modify a readonly attribute on a update operation", func() {

			l1.ReadOnly = "value"
			l2.ReadOnly = "not value"

			errs := ValidateAdvancedSpecification(l1, l2, OperationUpdate)

			Convey("Then errs should not be nil", func() {
				So(errs, ShouldNotBeNil)
				So(len(errs), ShouldEqual, 1)
			})
		})

		Convey("When I try to modify a creationonly attribute on a create operation", func() {

			l1.CreationOnly = "value"
			l2.CreationOnly = "not value"

			errs := ValidateAdvancedSpecification(l1, l2, OperationCreate)

			Convey("Then errs should be nil", func() {
				So(errs, ShouldBeNil)
			})
		})

		Convey("When I try to modify a creationonly attribute on a create update", func() {

			l1.CreationOnly = "value"
			l2.CreationOnly = "not value"

			errs := ValidateAdvancedSpecification(l1, l2, OperationUpdate)

			Convey("Then errs should not be nil", func() {
				So(errs, ShouldNotBeNil)
				So(len(errs), ShouldEqual, 1)
			})
		})

		Convey("When I try to modify a creationonly and a readonly attribute on a create update", func() {

			l1.ReadOnly = "value"
			l2.ReadOnly = "not value"

			l1.CreationOnly = "value"
			l2.CreationOnly = "not value"

			errs := ValidateAdvancedSpecification(l1, l2, OperationUpdate)

			Convey("Then errs should not be nil", func() {
				So(errs, ShouldNotBeNil)
				So(len(errs), ShouldEqual, 2)
			})
		})
	})
}
