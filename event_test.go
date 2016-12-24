// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// UserIdentity represents the Identity of the object
var TagIdentity = Identity{
	Name:     "tag",
	Category: "tag",
}

type Tag struct {
	ID          string `cql:"id"`
	Description string `cql:"description"`
	Name        string `cql:"name"`
	Type        int    `cql:"type"`
}

func (t *Tag) Identifier() string {
	return t.ID
}

// Identity returns the Identity of the object.
func (t *Tag) Identity() Identity {

	return TagIdentity
}

// SetIdentifier sets the value of the object's unique identifier.
func (t *Tag) SetIdentifier(ID string) {
	t.ID = ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (t *Tag) Validate() Errors {
	return nil
}

func TestEvent_NewEvent(t *testing.T) {

	Convey("Given I create an Event", t, func() {

		tag := &Tag{}
		e := NewEvent(EventCreate, tag)

		Convey("Then the Error should be correctly initialized", func() {
			d, _ := json.Marshal(tag)
			So(e.Identity, ShouldEqual, "tag")
			So(e.Type, ShouldEqual, EventCreate)
			So(e.Entity, ShouldResemble, json.RawMessage(d))
		})
	})

	Convey("Given I create an Event with an unmarshalable entity", t, func() {

		tag := &UnmarshalableList{}

		Convey("Then it should panic", func() {
			So(func() { NewEvent(EventCreate, tag) }, ShouldPanic)
		})
	})
}

func TestEvent_Decode(t *testing.T) {

	Convey("Given I create an Event", t, func() {

		tag := &Tag{Name: "t1"}
		e := NewEvent(EventCreate, tag)
		d, _ := json.Marshal(tag)
		e.Entity = d

		Convey("When I decode the data", func() {
			t2 := &Tag{}

			e.Decode(t2)

			Convey("Then t2 should ressemble to tag", func() {
				So(t2, ShouldResemble, tag)
			})
		})
	})
}

func TestEvent_String(t *testing.T) {

	Convey("Given I create an Event", t, func() {

		tag := &Tag{Name: "t1"}
		e := NewEvent(EventCreate, tag)

		Convey("When I use String", func() {
			str := e.String()

			Convey("Then the string representatipn should be correct", func() {
				So(str, ShouldEqual, "<event type: create identity: tag>")
			})
		})
	})
}

func TestEvent_NewEvents(t *testing.T) {

	Convey("Given I create an Events", t, func() {

		tag := &Tag{}
		e1 := NewEvent(EventCreate, tag)
		e2 := NewEvent(EventDelete, tag)

		evts := NewEvents(e1, e2)

		Convey("Then the Error should be correctly initialized", func() {
			So(len(evts), ShouldEqual, 2)
		})
	})
}
