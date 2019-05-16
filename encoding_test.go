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

		Convey(fmt.Sprintf("Encode will throw an error if passed in a nil interface or an interface holding a nil pointer using encoding %s", encoding), t, func() {

			var nilPointer *int
			cases := map[string]interface{}{
				"nil interface": nil,
				"nil pointer":   nilPointer,
			}

			for description, object := range cases {
				Convey(fmt.Sprintf("passing in a %s should result in an error", description), func() {
					data, err := Encode(encoding, object)
					So(data, ShouldBeNil)
					So(err, ShouldNotBeNil)
				})
			}
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

	Convey("Given I have good msgpack header", t, func() {

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

	Convey("Given I have good json header", t, func() {

		h := http.Header{}
		h.Set("Content-Type", "application/json; something=cool")
		h.Set("Accept", "application/json")

		Convey("When I call EncodingFromHeaders", func() {

			r, w, err := EncodingFromHeaders(h)

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

	Convey("Given I have a classic browser Accept header", t, func() {

		h := http.Header{}
		h.Set("Content-Type", "application/msgpack; something=cool")
		h.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")

		Convey("When I call EncodingFromHeaders", func() {

			r, w, err := EncodingFromHeaders(h)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r, ShouldEqual, EncodingTypeMSGPACK)
			})

			Convey("Then w should be correct", func() {
				So(w, ShouldEqual, EncodingTypeJSON)
			})
		})
	})

	Convey("Given I have an unaccepatble content-type header", t, func() {

		h := http.Header{}
		h.Set("Content-Type", "application/ppt")
		h.Set("Accept", "application/msgpack")

		Convey("When I call EncodingFromHeaders", func() {

			_, _, err := EncodingFromHeaders(h)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 415 (elemental): Unsupported Media Type: Cannot find any acceptable Content-Type media type in provided header: application/ppt`)
			})
		})
	})

	Convey("Given I have an unaccepatble accept header", t, func() {

		h := http.Header{}
		h.Set("Content-Type", "application/msgpack")
		h.Set("Accept", "application/ppt,application/toto")

		Convey("When I call EncodingFromHeaders", func() {

			_, _, err := EncodingFromHeaders(h)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 415 (elemental): Unsupported Media Type: Cannot find any acceptable Accept media type in provided header: application/ppt,application/toto`)
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

	Convey("Given I have an registered content-type header", t, func() {

		h := http.Header{}
		h.Set("Content-Type", "application/aaaa")
		h.Set("Accept", "application/msgpack")
		RegisterSupportedContentType("application/aaaa")

		Convey("When I call EncodingFromHeaders", func() {

			_, _, err := EncodingFromHeaders(h)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})

	Convey("Given I have an registered accept header", t, func() {

		h := http.Header{}
		h.Set("Content-Type", "application/msgpack")
		h.Set("Accept", "application/ppt,application/toto,application/aaaa")
		RegisterSupportedAcceptType("application/aaaa")

		Convey("When I call EncodingFromHeaders", func() {

			_, _, err := EncodingFromHeaders(h)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
