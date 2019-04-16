package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResponse_NewResponse(t *testing.T) {

	Convey("Given I create a new response", t, func() {

		r := NewResponse(&Request{RequestID: "x"})

		Convey("Then it should be correctly initialized", func() {
			So(r, ShouldNotBeNil)
			So(r.RequestID, ShouldEqual, "x")
		})
	})
}

func TestEncode(t *testing.T) {

	Convey("Given I have a list and a request that Accepts JSON", t, func() {

		req := &Request{
			Accept: EncodingTypeJSON,
		}
		resp := NewResponse(req)

		Convey("When I call Encode", func() {

			lst := NewList()
			err := resp.Encode(lst)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then it should be correctly encoded", func() {
				l, _ := Encode(EncodingTypeJSON, lst)
				So(resp.Data, ShouldResemble, l)
			})
		})
	})

	Convey("Given I have a list and a request that Accepts MSGPACK", t, func() {

		req := &Request{
			Accept: EncodingTypeMSGPACK,
		}
		resp := NewResponse(req)

		Convey("When I call Encode", func() {

			lst := NewList()
			err := resp.Encode(lst)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then it should be correctly encoded", func() {
				l, _ := Encode(EncodingTypeMSGPACK, lst)
				So(resp.Data, ShouldResemble, l)
			})
		})
	})
}
