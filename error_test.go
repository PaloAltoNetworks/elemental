// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"fmt"
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
			So(errs.Error(), ShouldResemble, "error 0: error 42 (containers): bad: something bad\nerror 1: error 43 (containers1): bad1: something bad1\n")
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
