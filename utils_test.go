// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtils_extractFieldNames(t *testing.T) {

	Convey("Given I have a list", t, func() {

		l1 := NewList()

		Convey("When I extract the fields", func() {

			fields := extractFieldNames(l1)

			Convey("Then all fields should be present", func() {
				So(len(fields), ShouldEqual, 9)
				So(fields, ShouldContain, "ID")
				So(fields, ShouldContain, "Description")
				So(fields, ShouldContain, "Name")
				So(fields, ShouldContain, "ParentID")
				So(fields, ShouldContain, "ParentType")
				So(fields, ShouldContain, "CreationOnly")
				So(fields, ShouldContain, "ReadOnly")
				So(fields, ShouldContain, "Unexposed")
				So(fields, ShouldContain, "Date")
			})
		})
	})
}

func TestUtils_areFieldValuesEqual(t *testing.T) {

	Convey("Given I have 2 list", t, func() {

		l1 := NewList()
		l2 := NewList()

		Convey("When I set the same name", func() {

			l1.Name = "list1"
			l2.Name = "list1"

			Convey("Then the values should be equal", func() {
				So(areFieldValuesEqual("Name", l1, l2), ShouldBeTrue)
			})
		})

		Convey("When I set a different name", func() {

			l1.Name = "list1"
			l2.Name = "list2"

			Convey("Then the values should not be equal", func() {
				So(areFieldValuesEqual("Name", l1, l2), ShouldBeFalse)
			})
		})

		Convey("When I set a same time", func() {

			l1.Date = time.Date(2009, time.November, 10, 10, 0, 0, 0, time.UTC)
			l2.Date = time.Date(2009, time.November, 10, 5, 0, 0, 0, time.FixedZone("Eastern", -5*3600))

			Convey("Then the values should not be equal", func() {
				So(areFieldValuesEqual("Date", l1, l2), ShouldBeTrue)
			})
		})

		Convey("When I set a different time", func() {

			l1.Date = time.Now()
			l2.Date = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

			Convey("Then the values should not be equal", func() {
				So(areFieldValuesEqual("Date", l1, l2), ShouldBeFalse)
			})
		})
	})
}

func TestUtils_isFieldValueZero(t *testing.T) {

	Convey("Given I have a struct", t, func() {

		type S struct {
			S   string
			B   bool
			I   int
			F   float32
			A   []string
			M   map[string]string
			Sub *S
		}

		Convey("When I set all zero values", func() {

			s := &S{"", false, 0, 0.0, nil, nil, &S{}}

			Convey("Then isFieldValueZero on S should return true", func() {
				So(isFieldValueZero("S", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on B should return true", func() {
				So(isFieldValueZero("B", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on I should return true", func() {
				So(isFieldValueZero("I", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on F should return true", func() {
				So(isFieldValueZero("F", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on A should return true", func() {
				So(isFieldValueZero("A", s), ShouldBeTrue)
			})

			Convey("Then isFieldValueZero on M should return true", func() {
				So(isFieldValueZero("M", s), ShouldBeTrue)
			})
		})

		Convey("When I set all non zero values", func() {

			s := &S{"hello", true, 1, 1.0, []string{"a"}, map[string]string{"a": "b"}, &S{S: "nope"}}

			Convey("Then isFieldValueZero on S should return false", func() {
				So(isFieldValueZero("S", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on B should return false", func() {
				So(isFieldValueZero("B", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on I should return false", func() {
				So(isFieldValueZero("I", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on F should return false", func() {
				So(isFieldValueZero("F", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on A should return false", func() {
				So(isFieldValueZero("A", s), ShouldBeFalse)
			})

			Convey("Then isFieldValueZero on M should return false", func() {
				So(isFieldValueZero("M", s), ShouldBeFalse)
			})
		})
	})
}
