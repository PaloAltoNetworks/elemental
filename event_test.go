// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
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

func TestError_NewEvent(t *testing.T) {

	Convey("Given I create an Event", t, func() {

		tag := &Tag{}
		e := NewEvent(EventCreate, tag)

		Convey("Then the Error should be correctly initialized", func() {
			So(e.Identity, ShouldEqual, "tag")
			So(e.Type, ShouldEqual, EventCreate)
			So(e.Entity, ShouldEqual, tag)
		})
	})
}

func TestError_NewEvents(t *testing.T) {

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
