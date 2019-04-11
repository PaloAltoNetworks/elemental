package elemental

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncodeDecode(t *testing.T) {

	test := func(encoding string) {

		Convey(fmt.Sprintf("Given I encode an object into the request using encoding %s", encoding), t, func() {

			o := &List{ID: "1", Name: "hello"}

			data, err := Encode(encoding, o)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then data should not be nil", func() {
				So(len(data), ShouldNotBeEmpty)
			})

			Convey("When I decode it", func() {
				o1 := &List{}

				err := Decode(encoding, data, o1)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then o2 should resemble to o", func() {
					So(o1, ShouldResemble, o)
				})
			})
		})

		Convey(fmt.Sprintf("Given I encode a list of objects into the request using encoding %s", encoding), t, func() {

			o := ListsList{
				&List{ID: "1", Name: "hello1"},
				&List{ID: "2", Name: "hello2"},
			}

			data, err := Encode(encoding, &o)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then data should not be nil", func() {
				So(len(data), ShouldNotBeEmpty)
			})

			Convey("When I decode it", func() {

				o1 := ListsList{}

				err := Decode(encoding, data, &o1)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then o2 should resemble to o", func() {
					So(o1, ShouldResemble, o)
				})
			})
		})

		Convey(fmt.Sprintf("Given I encode an unmarshalable object into the request using encoding %s", encoding), t, func() {

			var o interface{}
			if encoding == "application/json" {
				o = &UnmarshalableList{}
			}

			data, err := Encode(encoding, o)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then data should be empty", func() {
				So(len(data), ShouldEqual, 0)
			})
		})

		Convey(fmt.Sprintf("Given I decode an unmarshalable object into the request using encoding %s", encoding), t, func() {

			var o interface{}
			if encoding == "application/json" {
				o = &UnmarshalableList{}
			}

			err := Decode(encoding, nil, o)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	}

	test("application/json")
	test("application/gob")
}
