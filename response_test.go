package elemental

import (
	"net/http"
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

func TestResponse_EncodeDecode(t *testing.T) {

	Convey("Given I create a new response", t, func() {
		r := NewResponse(&Request{RequestID: "x"})

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

			Convey("When I Decode it but the response code is set to 204", func() {

				r.Data = nil
				r.StatusCode = http.StatusNoContent
				err := r.Decode(nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("When I Decode an nil object but the response code is not set to 204", func() {

				r.Data = nil
				err := r.Decode(nil)

				Convey("Then err should be nil", func() {
					So(err, ShouldNotBeNil)
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

// Encode 1000 identifiables
// json:              1539060 ns/op
// jsoniter:           925631 ns/op
func Benchmark_RequestEncode(b *testing.B) {

	r := NewResponse(&Request{RequestID: "x"})
	o := make(ListsList, 1000)
	for i := 0; i < 1000; i++ {
		o[i] = &List{ID: "1", Name: "hello"}
	}

	for i := 0; i < b.N; i++ {
		_ = r.Encode(o)
	}
}

// Decode 1000 identifiables
// json:                3890001 ns/op
// jsoniter:             946838 ns/op
func Benchmark_RequestDecode(b *testing.B) {

	r := NewResponse(&Request{RequestID: "x"})
	o := make(ListsList, 1000)
	for i := 0; i < 1000; i++ {
		o[i] = &List{ID: "1", Name: "hello"}
	}
	_ = r.Encode(o)

	o = make(ListsList, 1000)
	for i := 0; i < b.N; i++ {
		r.Decode(&o)
	}
}
