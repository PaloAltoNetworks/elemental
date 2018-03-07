// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestEvent_NewEvent(t *testing.T) {
	Convey("Given I create an Event", t, func() {

		list := &List{}
		e := NewEvent(EventCreate, list)

		Convey("Then the Error should be correctly initialized", func() {
			d, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(list)
			So(e.Identity, ShouldEqual, "list")
			So(e.Type, ShouldEqual, EventCreate)
			So(e.Entity, ShouldResemble, jsoniter.RawMessage(d))
		})
	})

	Convey("Given I create an Event with an unmarshalable entity", t, func() {

		list := &UnmarshalableList{}

		Convey("Then it should panic", func() {
			So(func() { NewEvent(EventCreate, list) }, ShouldPanic)
		})
	})
}

func TestEvent_Decode(t *testing.T) {

	Convey("Given I create an Event", t, func() {

		list := &List{Name: "t1"}
		e := NewEvent(EventCreate, list)
		d, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(list)
		e.Entity = d

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
		e := NewEvent(EventCreate, list)

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

		Convey("Then the Error should be correctly initialized", func() {
			So(len(evts), ShouldEqual, 2)
		})
	})
}
