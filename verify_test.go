package elemental

import (
	"net/http"
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

			Convey("Then errs should be nil", func() {
				So(errs, ShouldBeNil)
			})
		})

		Convey("When I try to set a readonly attribute on a create operation", func() {

			l1.ReadOnly = "value"

			errs := ValidateAdvancedSpecification(l1, nil, OperationCreate).(Errors)

			Convey("Then errs should not be nil", func() {
				So(errs, ShouldNotBeNil)
				So(len(errs), ShouldEqual, 1)
				So(errs.Code(), ShouldEqual, http.StatusUnprocessableEntity)
			})
		})

		Convey("When I try to set a readonly attribute on a create operation that has the same value as the pristine", func() {

			l1.ReadOnly = "value"
			l2.ReadOnly = "value"

			errs := ValidateAdvancedSpecification(l1, l2, OperationCreate)

			Convey("Then errs should be nil", func() {
				So(errs, ShouldBeNil)
			})
		})

		Convey("When I try to modify a readonly attribute on a update operation", func() {

			l1.ReadOnly = "value"
			l2.ReadOnly = "not value"

			errs := ValidateAdvancedSpecification(l1, l2, OperationUpdate).(Errors)

			Convey("Then errs should not be nil", func() {
				So(errs, ShouldNotBeNil)
				So(len(errs), ShouldEqual, 1)
				So(errs.Code(), ShouldEqual, http.StatusUnprocessableEntity)
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

			errs := ValidateAdvancedSpecification(l1, l2, OperationUpdate).(Errors)

			Convey("Then errs should not be nil", func() {
				So(errs, ShouldNotBeNil)
				So(len(errs), ShouldEqual, 1)
				So(errs.Code(), ShouldEqual, http.StatusUnprocessableEntity)
			})
		})

		Convey("When I try to modify a creationonly and a readonly attribute on a create update", func() {

			l1.ReadOnly = "value"
			l2.ReadOnly = "not value"

			l1.CreationOnly = "value"
			l2.CreationOnly = "not value"

			errs := ValidateAdvancedSpecification(l1, l2, OperationUpdate).(Errors)

			Convey("Then errs should not be nil", func() {
				So(errs, ShouldNotBeNil)
				So(len(errs), ShouldEqual, 2)
				So(errs.At(0).Code, ShouldEqual, http.StatusUnprocessableEntity)
				So(errs.At(1).Code, ShouldEqual, http.StatusUnprocessableEntity)
			})
		})
	})
}

func TestVerify_BackportUnexposedFields(t *testing.T) {

	Convey("Given have to objects with unexposed fields", t, func() {

		l1 := NewList()
		l2 := NewList()

		l1.Name = "l1"
		l2.Name = "l2"

		l1.Unexposed = "u1"
		l2.Unexposed = "u2"

		Convey("When I backport unexposed fields from l1 to l2", func() {

			BackportUnexposedFields(l1, l2)

			Convey("Then the name should be different", func() {
				So(l1.Name, ShouldEqual, "l1")
				So(l2.Name, ShouldEqual, "l2")
			})

			Convey("Then the Unexposed attribute should be equal", func() {
				So(l1.Unexposed, ShouldEqual, "u1")
				So(l2.Unexposed, ShouldEqual, "u1")
			})
		})
	})

	Convey("Given have to objects with secret fields", t, func() {

		l1 := NewList()
		l2 := NewList()

		Convey("When I backport secrets fields from l1 to l2 with no change in l2", func() {

			l1.Secret = "u1"
			l2.Secret = "u1"

			BackportUnexposedFields(l1, l2)

			Convey("Then the Unexposed attribute should be equal", func() {
				So(l1.Secret, ShouldEqual, "u1")
				So(l2.Secret, ShouldEqual, "u1")
			})
		})

		Convey("When I backport secrets fields from l1 to l2 with empty changes in l2", func() {

			l1.Secret = "u1"
			l2.Secret = ""

			BackportUnexposedFields(l1, l2)

			Convey("Then the Unexposed attribute should be equal", func() {
				So(l1.Secret, ShouldEqual, "u1")
				So(l2.Secret, ShouldEqual, "u1")
			})
		})

		Convey("When I backport secrets fields from l1 to l2 with changes in l2", func() {

			l1.Secret = "u1"
			l2.Secret = "u2"

			BackportUnexposedFields(l1, l2)

			Convey("Then the Unexposed attribute should be equal", func() {
				So(l1.Secret, ShouldEqual, "u1")
				So(l2.Secret, ShouldEqual, "u2")
			})
		})
	})
}

func TestVerify_ResetDefaultForZeroValues(t *testing.T) {

	Convey("Given I have a task with an empty string as status", t, func() {

		task := NewTask()
		task.Status = ""

		Convey("When I call ResetDefaultForZeroValues", func() {

			ResetDefaultForZeroValues(task)

			Convey("Then the stats should be reset", func() {
				So(task.Status, ShouldEqual, "TODO")
			})
		})

	})
}
