// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRequest_NewRequest(t *testing.T) {

	Convey("Given I create a new request", t, func() {
		r := NewRequest("ns", OperationCreate, ListIdentity)

		Convey("Then it should be correctly initialized", func() {
			So(r.RequestID, ShouldNotBeEmpty)
			So(r.Namespace, ShouldEqual, "ns")
			So(r.Operation, ShouldEqual, OperationCreate)
			So(r.Identity, ShouldResemble, ListIdentity)
		})
	})
}

func TestRequest_EncodeDecode(t *testing.T) {

	Convey("Given I create a new request", t, func() {
		r := NewRequest("ns", OperationCreate, ListIdentity)

		Convey("When I encode an object into the request", func() {

			o := &List{ID: "1", Name: "hello"}
			err := r.Encode(o)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then data should not be nil", func() {
				So(len(r.Data), ShouldNotBeEmpty)
			})

			Convey("When I Decode it", func() {
				o1 := &List{}

				err := r.Decode(&o1)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then o2 should resemble to o", func() {
					So(o1, ShouldResemble, o)
				})
			})
		})

		Convey("When I encode an unmarshallable object into the request", func() {

			o := &UnmarshalableList{}
			err := r.Encode(o)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then data should be empty", func() {
				So(len(r.Data), ShouldEqual, 0)
			})

			Convey("When I Decode it", func() {
				o1 := &List{}

				err := r.Decode(&o1)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestResponse_NewReponse(t *testing.T) {

	Convey("Given I create a new reponse", t, func() {

		r := NewResponse()

		Convey("Then it should be correctly initialized", func() {
			So(r, ShouldNotBeNil)
		})
	})
}

func TestResponse_EncodeDecode(t *testing.T) {

	Convey("Given I create a new response", t, func() {
		r := NewResponse()

		Convey("When I encode an object into the response", func() {

			o := &List{ID: "1", Name: "hello"}
			err := r.Encode(o)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then data should not be nil", func() {
				So(len(r.Data), ShouldNotBeEmpty)
			})

			Convey("When I Decode it", func() {
				o1 := &List{}

				err := r.Decode(&o1)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then o2 should resemble to o", func() {
					So(o1, ShouldResemble, o)
				})
			})
		})

		Convey("When I encode an unmarshallable object into the response", func() {

			o := &UnmarshalableList{}
			err := r.Encode(o)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then data should be empty", func() {
				So(len(r.Data), ShouldEqual, 0)
			})

			Convey("When I Decode it", func() {
				o1 := &List{}

				err := r.Decode(&o1)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}
