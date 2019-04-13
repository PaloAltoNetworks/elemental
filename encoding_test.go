package elemental

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncodeDecode(t *testing.T) {

	test := func(encoding EncodingType) {

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

		// Convey(fmt.Sprintf("Given I encode an unmarshalable object into the request using encoding %s", encoding), t, func() {

		// 	o := &UnmarshalableList{}

		// 	data, err := Encode(encoding, o)

		// 	Convey("Then err should not be nil", func() {
		// 		So(err, ShouldNotBeNil)
		// 	})

		// 	Convey("Then data should be empty", func() {
		// 		So(len(data), ShouldEqual, 0)
		// 	})
		// })

		Convey(fmt.Sprintf("Given I decode an unmarshalable object into the request using encoding %s", encoding), t, func() {

			o := &UnmarshalableList{}

			err := Decode(encoding, nil, o)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	}

	test(EncodingTypeJSON)
	test(EncodingTypeMSGPACK)
}

func TestConvert(t *testing.T) {

	Convey("Given I have some data encoded in JSON", t, func() {

		l := NewList()
		l.Name = "hello"
		data, err := Encode(EncodingTypeJSON, l)
		if err != nil {
			panic(err)
		}

		Convey("When I call Convert to make it MSGPACK", func() {

			cdata, err := Convert(EncodingTypeJSON, EncodingTypeMSGPACK, data)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then converted data should be correct", func() {
				l2 := NewList()
				Decode(EncodingTypeMSGPACK, cdata, l2) // nolint
				So(l2, ShouldResemble, l)
			})
		})
	})

	Convey("Given I have some data encoded in MSGPACK", t, func() {

		l := NewList()
		l.Name = "hello"
		data, err := Encode(EncodingTypeMSGPACK, l)
		if err != nil {
			panic(err)
		}

		Convey("When I call Convert to make it JSON", func() {

			cdata, err := Convert(EncodingTypeMSGPACK, EncodingTypeJSON, data)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then converted data should be correct", func() {
				l2 := NewList()
				Decode(EncodingTypeJSON, cdata, l2) // nolint
				So(l2, ShouldResemble, l)
			})
		})
	})

	Convey("Given I have some data encoded in MSGPACK", t, func() {

		l := NewList()
		l.Name = "hello"
		data, err := Encode(EncodingTypeMSGPACK, l)
		if err != nil {
			panic(err)
		}

		Convey("When I call Convert to make it MSGPACK", func() {

			cdata, err := Convert(EncodingTypeMSGPACK, EncodingTypeMSGPACK, data)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then converted data should be correct", func() {
				l2 := NewList()
				Decode(EncodingTypeMSGPACK, cdata, l2) // nolint
				So(l2, ShouldResemble, l)
			})
		})
	})

	Convey("Given I have some data encoded in JSON", t, func() {

		l := NewList()
		l.Name = "hello"
		data, err := Encode(EncodingTypeJSON, l)
		if err != nil {
			panic(err)
		}

		Convey("When I call Convert to make it JSON", func() {

			cdata, err := Convert(EncodingTypeJSON, EncodingTypeJSON, data)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then converted data should be correct", func() {
				l2 := NewList()
				Decode(EncodingTypeJSON, cdata, l2) // nolint
				So(l2, ShouldResemble, l)
			})
		})
	})

	Convey("Given I have some invalid JSON", t, func() {

		_, err := Convert(EncodingTypeJSON, EncodingTypeMSGPACK, []byte("woops"))

		Convey("Then err should be correct", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, `unable to decode application/json: json decode error [pos 1]: read map - expect char '{' but got char 'w'`)
		})
	})

	Convey("Given I have some invalid MSGPACK", t, func() {

		_, err := Convert(EncodingTypeMSGPACK, EncodingTypeJSON, []byte("woops"))

		Convey("Then err should be correct", func() {
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, `unable to decode application/msgpack: msgpack decode error [pos 1]: cannot read container length: unrecognized descriptor byte: hex: 77, decimal: 119`)
		})
	})
}

func TestEncodingFromHeader(t *testing.T) {

	Convey("Given I have no header", t, func() {

		Convey("When I call EncodingFromHeaders", func() {

			r, w, err := EncodingFromHeaders(nil)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r, ShouldEqual, EncodingTypeJSON)
			})

			Convey("Then w should be correct", func() {
				So(w, ShouldEqual, EncodingTypeJSON)
			})
		})
	})

	Convey("Given I have good header", t, func() {

		h := http.Header{}
		h.Set("Content-Type", "application/msgpack; something=cool")
		h.Set("Accept", "application/msgpack")

		Convey("When I call EncodingFromHeaders", func() {

			r, w, err := EncodingFromHeaders(h)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r, ShouldEqual, EncodingTypeMSGPACK)
			})

			Convey("Then w should be correct", func() {
				So(w, ShouldEqual, EncodingTypeMSGPACK)
			})
		})
	})

	Convey("Given I have wrong content-type header", t, func() {

		h := http.Header{}
		h.Set("Content-Type", "NA99sf9sdf99 dfgjdhfgjh")
		h.Set("Accept", "application/msgpack")

		Convey("When I call EncodingFromHeaders", func() {

			_, _, err := EncodingFromHeaders(h)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Bad Request: Invalid Content-Type header: mime: expected slash after first token`)
			})
		})
	})

	Convey("Given I have wrong accept header", t, func() {

		h := http.Header{}
		h.Set("Content-Type", "application/msgpack")
		h.Set("Accept", "NA99sf9sdf99 dfgjdhfgjh")

		Convey("When I call EncodingFromHeaders", func() {

			_, _, err := EncodingFromHeaders(h)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Bad Request: Invalid Accept header: mime: expected slash after first token`)
			})
		})
	})
}
