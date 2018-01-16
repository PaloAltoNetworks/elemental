// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestError_NewError(t *testing.T) {

	Convey("Given I create an Error", t, func() {

		e := NewError("bad", "something bad", "containers", 42)

		Convey("Then the Error should be correctly initialized", func() {
			So(e.Code, ShouldEqual, 42)
			So(e.Description, ShouldEqual, "something bad")
			So(e.Subject, ShouldEqual, "containers")
			So(e.Title, ShouldEqual, "bad")
		})

		Convey("Then the string representation should be correct", func() {
			So(e.Error(), ShouldEqual, "error 42 (containers): bad: something bad")
		})
	})

	Convey("Given I create an Error with data", t, func() {

		e := NewErrorWithData("bad", "something bad", "containers", 42, map[string]string{"test": "coucou"})

		Convey("Then the Error should be correctly initialized", func() {
			So(e.Code, ShouldEqual, 42)
			So(e.Description, ShouldEqual, "something bad")
			So(e.Subject, ShouldEqual, "containers")
			So(e.Title, ShouldEqual, "bad")
			So(e.Data, ShouldResemble, map[string]string{"test": "coucou"})
		})

		Convey("Then the string representation should be correct", func() {
			So(e.Error(), ShouldEqual, "error 42 (containers): bad: something bad")
		})
	})
}

func TestError_Error(t *testing.T) {

	Convey("Given I create an Error", t, func() {

		e := NewError("bad", "something bad", "containers", 42)

		Convey("When I use the Error interface", func() {
			s := e.Error()

			Convey("Then string should be correct", func() {
				So(s, ShouldEqual, "error 42 (containers): bad: something bad")
			})
		})
	})
}

func TestError_NewErrors(t *testing.T) {

	Convey("Given I create an elemental.Errors with some errors", t, func() {

		e1 := NewError("bad", "something bad", "containers", 42)
		e2 := NewError("bad1", "something bad1", "containers1", 43)

		errs := NewErrors(e1, e2)

		Convey("Then the Error should be correctly initialized", func() {
			So(errs, ShouldResemble, Errors{e1, e2})
			So(errs.Error(), ShouldResemble, "error 42 (containers): bad: something bad, error 43 (containers1): bad1: something bad1")
			So(errs.Code(), ShouldEqual, 42)
		})
	})

	Convey("Given I create an elemental.Errors without any error", t, func() {

		errs := NewErrors()

		Convey("Then the Error should be correctly initialized", func() {
			So(errs, ShouldResemble, Errors{})
			So(errs.Code(), ShouldEqual, -1)
		})
	})

	Convey("Given I create an elemental.Errors with a standard error", t, func() {

		e := fmt.Errorf("wesh")
		errs := NewErrors(e)

		Convey("Then the Error should be correctly initialized", func() {
			So(errs, ShouldResemble, Errors{e})
			So(errs.Code(), ShouldEqual, -1)
		})
	})
}

func TestError_At(t *testing.T) {

	Convey("Given I create an elemental.Errors with some errors", t, func() {

		e1 := NewError("bad", "something bad", "containers", 42)
		e2 := fmt.Errorf("not good")

		errs := NewErrors(e1, e2)

		Convey("Then the Error should be correctly initialized", func() {
			So(errs.At(0), ShouldResemble, e1)
			So(errs.At(1).Code, ShouldEqual, -1)
		})
	})
}

func TestError_IsValidationError(t *testing.T) {

	Convey("Given I have a list of one validation error", t, func() {

		err := NewError("the title", "the description", "http.test", http.StatusUnprocessableEntity)
		err.Data = map[string]interface{}{"attribute": "theattr"}

		errs := NewErrors(err)

		Convey("When I call IsValidationError with expected title and attribute", func() {

			ok := IsValidationError(errs, "the title", "theattr")

			Convey("Then is should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsValidationError with expected title and non expected attribute", func() {

			ok := IsValidationError(errs, "the title", "not-theattr")

			Convey("Then is should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsValidationError with non expected title and expected attribute", func() {

			ok := IsValidationError(errs, "not the title", "theattr")

			Convey("Then is should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have a list of multiple validation errors", t, func() {

		err := NewError("the title", "the description", "http.test", http.StatusUnprocessableEntity)
		err.Data = map[string]interface{}{"attribute": "theattr"}

		errs := NewErrors(err, err)

		Convey("When I call IsValidationError with expected title and attribute", func() {

			ok := IsValidationError(errs, "the title", "theattr")

			Convey("Then is should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have a single validation error", t, func() {

		err := NewError("the title", "the description", "http.test", http.StatusUnprocessableEntity)
		err.Data = map[string]interface{}{"attribute": "theattr"}

		Convey("When I call IsValidationError with expected title and attribute", func() {

			ok := IsValidationError(err, "the title", "theattr")

			Convey("Then is should be ok", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have a classic error", t, func() {

		err := errors.New("not elemental")

		Convey("When I call IsValidationError", func() {

			ok := IsValidationError(err, "the title", "theattr")

			Convey("Then is should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have a list of non validation elemental error", t, func() {

		err := NewError("the title", "the description", "http.test", http.StatusNotFound)
		err.Data = map[string]interface{}{"attribute": "theattr"}

		errs := NewErrors(err)

		Convey("When I call IsValidationError with expected title and attribute", func() {

			ok := IsValidationError(errs, "the title", "theattr")

			Convey("Then is should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have a non validation elemental error", t, func() {

		err := NewError("the title", "the description", "http.test", http.StatusNotFound)
		err.Data = map[string]interface{}{"attribute": "theattr"}

		Convey("When I call IsValidationError with expected title and attribute", func() {

			ok := IsValidationError(err, "the title", "theattr")

			Convey("Then is should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have a validation elemental error with no data", t, func() {

		err := NewError("the title", "the description", "http.test", http.StatusUnprocessableEntity)

		Convey("When I call IsValidationError with expected title and attribute", func() {

			ok := IsValidationError(err, "the title", "theattr")

			Convey("Then is should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have a validation elemental error with no bad data", t, func() {

		err := NewError("the title", "the description", "http.test", http.StatusUnprocessableEntity)
		err.Data = "zob"

		Convey("When I call IsValidationError with expected title and attribute", func() {

			ok := IsValidationError(err, "the title", "theattr")

			Convey("Then is should not be ok", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}
