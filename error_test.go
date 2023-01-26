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
	"errors"
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestError_NewError(t *testing.T) {

	Convey("Given I create an Error", t, func() {

		e := NewError("bad", "something bad", "containers", 42)
		e.Trace = "xyz"

		Convey("Then the Error should be correctly initialized", func() {
			So(e.Code, ShouldEqual, 42)
			So(e.Description, ShouldEqual, "something bad")
			So(e.Subject, ShouldEqual, "containers")
			So(e.Title, ShouldEqual, "bad")
			So(e.Trace, ShouldEqual, "xyz")
		})

		Convey("Then the string representation should be correct", func() {
			So(e.Error(), ShouldEqual, "error 42 (containers): bad: something bad [trace: xyz]")
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
}

func TestError_Append(t *testing.T) {

	Convey("Given I append to an elemental.Errors some other errors", t, func() {

		e0 := NewError("bad", "something bad", "containers", 42)
		errs := NewErrors(e0)

		e1 := NewError("bad", "something bad", "containers", 42)
		e2 := NewError("bad", "something bad", "containers", 42)

		errs2 := errs.Append(e1, e2, errs, fmt.Errorf("boom"))

		Convey("Then out should be correct", func() {
			So(errs2[0], ShouldResemble, e0)
			So(errs2[1], ShouldResemble, e1)
			So(errs2[2], ShouldResemble, e2)
			So(errs2[3], ShouldResemble, e0)
			So(errs2[4], ShouldResemble, NewError("Internal Server Error", "boom", "elemental", 500))
		})
	})
}

func TestError_SetTraceID(t *testing.T) {

	Convey("Given I append to an elemental.Errors some other errors", t, func() {

		e0 := NewError("bad", "something bad", "containers", 42)
		e1 := NewError("bad", "something bad", "containers", 42)
		e2 := NewError("bad", "something bad", "containers", 42)
		errs := NewErrors(e0, e1, e2)

		Convey("When I call setTraceID", func() {

			errs = errs.Trace("trace")

			Convey("Then the trace is set in all errors", func() {
				So(errs[0].Trace, ShouldEqual, "trace")
				So(errs[1].Trace, ShouldEqual, "trace")
				So(errs[2].Trace, ShouldEqual, "trace")
			})
		})
	})
}

func TestError_IsValidationError(t *testing.T) {

	Convey("Given I have a list of one validation error", t, func() {

		err := NewError("the title", "the description", "http.test", http.StatusUnprocessableEntity)
		err.Data = map[string]any{"attribute": "theattr"}

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
		err.Data = map[string]any{"attribute": "theattr"}

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
		err.Data = map[string]any{"attribute": "theattr"}

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
		err.Data = map[string]any{"attribute": "theattr"}

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
		err.Data = map[string]any{"attribute": "theattr"}

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

func TestError_DecodeError(t *testing.T) {

	Convey("Given I have some valid json data", t, func() {

		data := `[{"title":"t1","code":42,"description":"coucou"},{"title":"t1","code":42,"description":"coucou"}]`

		Convey("When I call DecodeErrors", func() {

			errs, err := DecodeErrors([]byte(data))

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then errs should be correct", func() {
				So(len(errs), ShouldEqual, 2)
			})
		})
	})

	Convey("Given I have some invalid json data", t, func() {

		data := `[{"title":"t1"]`

		Convey("When I call DecodeErrors", func() {

			errs, err := DecodeErrors([]byte(data))

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "invalid character ']' after object key:value pair")
			})

			Convey("Then errs should be correct", func() {
				So(errs, ShouldBeNil)
			})
		})
	})
}

func TestIsErrorWithCode(t *testing.T) {
	type args struct {
		err  error
		code int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"positive test with elemental.Error",
			args{
				NewError("title", "description", "subject", 404),
				404,
			},
			true,
		},
		{
			"negative test with elemental.Error",
			args{
				NewError("title", "description", "subject", 404),
				500,
			},
			false,
		},
		{
			"positive test with elemental.Errors",
			args{
				NewErrors(NewError("title", "description", "subject", 404)),
				404,
			},
			true,
		},
		{
			"negative test with elemental.Errors",
			args{
				NewErrors(NewError("title", "description", "subject", 404)),
				500,
			},
			false,
		},
		{
			"test with classic error",
			args{
				fmt.Errorf("boom"),
				500,
			},
			false,
		},
		{
			"test with nil",
			args{
				nil,
				500,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsErrorWithCode(tt.args.err, tt.args.code); got != tt.want {
				t.Errorf("IsErrorWithCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
