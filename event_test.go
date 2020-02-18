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
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEvent_NewEvent(t *testing.T) {

	Convey("Given I create an Event using EncodingTypeJSON", t, func() {

		list := &List{}
		e := NewEventWithEncoding(EventCreate, list, EncodingTypeJSON)

		Convey("Then the Event should be correctly initialized", func() {
			d, _ := Encode(EncodingTypeJSON, list)
			So(e.Identity, ShouldEqual, "list")
			So(e.Type, ShouldEqual, EventCreate)
			So(e.Encoding, ShouldEqual, EncodingTypeJSON)
			So(e.JSONData, ShouldResemble, json.RawMessage(d))
			So(e.RawData, ShouldBeNil)
			So(e.Entity(), ShouldResemble, []byte(e.JSONData))
		})
	})

	Convey("Given I create an Event using EncodingTypeMSGPACK", t, func() {

		list := &List{}
		e := NewEventWithEncoding(EventCreate, list, EncodingTypeMSGPACK)

		Convey("Then the Event should be correctly initialized", func() {
			d, _ := Encode(EncodingTypeMSGPACK, list)
			So(e.Identity, ShouldEqual, "list")
			So(e.Type, ShouldEqual, EventCreate)
			So(e.Encoding, ShouldEqual, EncodingTypeMSGPACK)
			So(e.JSONData, ShouldBeNil)
			So(e.RawData, ShouldResemble, d)
			So(e.Entity(), ShouldResemble, e.RawData)
		})
	})

	Convey("Given I create an Event with an unmarshalable entity", t, func() {

		Convey("Then it should panic", func() {
			So(func() { NewEvent(EventCreate, nil) }, ShouldPanicWith, "unable to create new event: encode received a nil object")
		})
	})
}

func TestNewErrorEvent(t *testing.T) {

	testErr := Error{
		Description: "some description",
		Subject:     "some subject",
		Title:       "some title",
		Data: map[string]interface{}{
			"someKey": "someValue",
		},
	}

	Convey("Given I create an error Event using EncodingTypeJSON", t, func() {

		e := NewErrorEvent(testErr, EncodingTypeJSON)

		Convey("Then the error Event should be correctly initialized", func() {
			d, _ := Encode(EncodingTypeJSON, testErr)
			So(e.Identity, ShouldEqual, "error")
			So(e.Type, ShouldEqual, EventError)
			So(e.Encoding, ShouldEqual, EncodingTypeJSON)
			So(e.JSONData, ShouldResemble, json.RawMessage(d))
			So(e.RawData, ShouldBeNil)
			So(e.Entity(), ShouldResemble, []byte(e.JSONData))
		})
	})

	Convey("Given I create an error Event using EncodingTypeMSGPACK", t, func() {

		e := NewErrorEvent(testErr, EncodingTypeMSGPACK)

		Convey("Then the error Event should be correctly initialized", func() {
			d, _ := Encode(EncodingTypeMSGPACK, testErr)
			So(e.Identity, ShouldEqual, "error")
			So(e.Type, ShouldEqual, EventError)
			So(e.Encoding, ShouldEqual, EncodingTypeMSGPACK)
			So(e.JSONData, ShouldBeNil)
			So(e.RawData, ShouldResemble, d)
			So(e.Entity(), ShouldResemble, e.RawData)
		})
	})
}

func TestEvent_Decode(t *testing.T) {

	Convey("Given I create an Event using EncodingTypeJSON", t, func() {

		list := &List{Name: "t1"}
		e := NewEventWithEncoding(EventCreate, list, EncodingTypeJSON)

		Convey("When I decode the data", func() {
			l2 := &List{}

			_ = e.Decode(l2)

			Convey("Then t2 should resemble to tag", func() {
				So(l2, ShouldResemble, list)
			})
		})
	})

	Convey("Given I create an Event using EncodingTypeMSGPACK", t, func() {

		list := &List{Name: "t1"}
		e := NewEventWithEncoding(EventCreate, list, EncodingTypeMSGPACK)

		Convey("When I decode the data", func() {
			l2 := &List{}

			_ = e.Decode(l2)

			Convey("Then t2 should resemble to tag", func() {
				So(l2, ShouldResemble, list)
			})
		})
	})
}

func TestEvent_String(t *testing.T) {

	Convey("Given I create an Event", t, func() {

		list := &List{Name: "t1"}
		e := NewEventWithEncoding(EventCreate, list, EncodingTypeJSON)

		Convey("When I use String", func() {
			str := e.String()

			Convey("Then the string representatipn should be correct", func() {
				So(str, ShouldEqual, "<event type: create identity: list>")
			})
		})
	})
}

func TestEvent_NewEvents(t *testing.T) {

	Convey("Given I create an Events", t, func() {

		list := &List{}
		e1 := NewEvent(EventCreate, list)
		e2 := NewEvent(EventDelete, list)

		evts := NewEvents(e1, e2)

		Convey("Then the Event should be correctly initialized", func() {
			So(len(evts), ShouldEqual, 2)
		})
	})
}

func TestEvent_Convert(t *testing.T) {

	Convey("Given I have an Event with EncodingTypeJSON data", t, func() {

		list := &List{
			Name: "hello",
		}

		e := NewEventWithEncoding(EventCreate, list, EncodingTypeJSON)

		Convey("When I Convert to EncodingTypeMSGPACK", func() {

			err := e.Convert(EncodingTypeMSGPACK)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the converted event should be correct", func() {
				// Here there is some changes in the way time is econded, making the bytes not completely equals
				l2 := &List{}
				_ = Decode(EncodingTypeMSGPACK, e.Entity(), l2)
				So(e.JSONData, ShouldBeNil)
				So(e.Encoding, ShouldEqual, EncodingTypeMSGPACK)
				So(list, ShouldResemble, l2)
			})
		})
	})

	Convey("Given I have an Event with EncodingTypeMSGPACK data", t, func() {

		list := &List{
			Name: "hello",
		}
		e := NewEventWithEncoding(EventCreate, list, EncodingTypeMSGPACK)

		Convey("When I Convert to EncodingTypeJSON", func() {

			err := e.Convert(EncodingTypeJSON)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the converted event should be correct", func() {
				d, _ := Encode(EncodingTypeJSON, list)
				So(string(e.Entity()), ShouldResemble, string(d))
				So(e.JSONData, ShouldResemble, json.RawMessage(d))
				So(e.RawData, ShouldBeNil)
				So(e.Encoding, ShouldEqual, EncodingTypeJSON)
			})
		})

		Convey("When I Convert to EncodingTypeMSGPACK", func() {

			err := e.Convert(EncodingTypeMSGPACK)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the converted event should be correct", func() {
				d, _ := Encode(EncodingTypeMSGPACK, list)
				So(string(e.Entity()), ShouldResemble, string(d))
				So(e.JSONData, ShouldBeNil)
				So(e.RawData, ShouldResemble, d)
				So(e.Encoding, ShouldEqual, EncodingTypeMSGPACK)
			})
		})
	})

	Convey("Given I have an Event with invalid msgpack data", t, func() {

		list := &List{
			Name: "hello",
		}
		e := NewEventWithEncoding(EventCreate, list, EncodingTypeMSGPACK)
		e.RawData = []byte("not-good")

		Convey("When I Convert to EncodingTypeMSGPACK", func() {

			err := e.Convert(EncodingTypeJSON)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then the event should still be correct", func() {
				So(e.JSONData, ShouldBeNil)
				So(e.RawData, ShouldNotBeNil)
				So(e.Encoding, ShouldEqual, EncodingTypeMSGPACK)
			})
		})
	})

	Convey("Given I have an Event with invalid json data", t, func() {

		list := &List{
			Name: "hello",
		}
		e := NewEventWithEncoding(EventCreate, list, EncodingTypeJSON)
		e.JSONData = []byte("not-good")

		Convey("When I Convert to EncodingTypeMSGPACK", func() {

			err := e.Convert(EncodingTypeMSGPACK)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then the event should still be correct", func() {
				So(e.JSONData, ShouldNotBeNil)
				So(e.RawData, ShouldBeNil)
				So(e.Encoding, ShouldEqual, EncodingTypeJSON)
			})
		})
	})
}

func TestEvent_Duplicate(t *testing.T) {

	Convey("Given I have an Event with encoding set to json", t, func() {

		list := &List{}
		e1 := NewEventWithEncoding(EventCreate, list, EncodingTypeJSON)

		Convey("When I Duplicate ", func() {

			e2 := e1.Duplicate()

			Convey("Then the duplicated event should be correct", func() {
				So(e2.Type, ShouldEqual, e1.Type)
				So(e2.Entity(), ShouldResemble, e1.Entity())
				So(e2.RawData, ShouldBeNil)
				So(e2.JSONData, ShouldResemble, e1.JSONData)
				So(fmt.Sprintf("%p", e2.JSONData), ShouldNotEqual, fmt.Sprintf("%p", e1.JSONData))
				So(e2.Identity, ShouldEqual, e1.Identity)
				So(e2.Timestamp, ShouldEqual, e1.Timestamp)
				So(e2.Encoding, ShouldEqual, e1.Encoding)
			})
		})
	})

	Convey("Given I have an Event with encoding set to msgpack", t, func() {

		list := &List{}
		e1 := NewEventWithEncoding(EventCreate, list, EncodingTypeMSGPACK)

		Convey("When I Duplicate ", func() {

			e2 := e1.Duplicate()

			Convey("Then the duplicated event should be correct", func() {
				So(e2.Type, ShouldEqual, e1.Type)
				So(e2.Entity(), ShouldResemble, e1.Entity())
				So(e2.RawData, ShouldResemble, e1.RawData)
				So(fmt.Sprintf("%p", e2.RawData), ShouldNotEqual, fmt.Sprintf("%p", e1.RawData))
				So(e2.JSONData, ShouldBeNil)
				So(e2.Identity, ShouldEqual, e1.Identity)
				So(e2.Timestamp, ShouldEqual, e1.Timestamp)
				So(e2.Encoding, ShouldEqual, e1.Encoding)
			})
		})
	})
}
