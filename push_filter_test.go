package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPushFilter_NewPushFilter(t *testing.T) {

	Convey("Given I create a new PushFilter", t, func() {
		f := NewPushFilter()

		Convey("Then it should be correctly initialized", func() {
			So(f.RequestID, ShouldNotBeEmpty)
			So(f.EventTypes, ShouldNotBeNil)
			So(f.Identities, ShouldNotBeNil)
		})
	})
}

func TestPushFilter_Duplicate(t *testing.T) {

	Convey("Given I create a new PushFilter", t, func() {

		f := NewPushFilter()
		i1 := MakeIdentity("i1", "i1")
		i2 := MakeIdentity("i2", "i2")

		f.Identities = []Identity{i1, i2}
		f.EventTypes = []EventType{EventCreate, EventDelete}

		Convey("When I duplicate it", func() {
			dup := f.Duplicate()

			Convey("Then it should be correctly duplicated", func() {
				So(dup.Identities, ShouldResemble, f.Identities)
				So(dup.Identities, ShouldNotEqual, f.Identities)
				So(dup.EventTypes, ShouldResemble, f.EventTypes)
				So(dup.EventTypes, ShouldNotEqual, f.EventTypes)
				So(dup.RequestID, ShouldNotEqual, f.RequestID)
				So(dup.RequestID, ShouldNotBeEmpty)
			})
		})
	})
}

func TestPushFilter_IsOperationFiltered(t *testing.T) {

	Convey("Given I create a new PushFilter", t, func() {

		f := NewPushFilter()
		f.EventTypes = []EventType{EventCreate, EventUpdate}

		Convey("When I check if OperationDelete is filtered", func() {

			filtered := f.IsEventTypeFiltered(EventDelete)

			Convey("Then filtered should be true", func() {
				So(filtered, ShouldBeTrue)
			})
		})

		Convey("When I check if OperationCreate is filtered", func() {

			filtered := f.IsEventTypeFiltered(EventCreate)

			Convey("Then filtered should be false", func() {
				So(filtered, ShouldBeFalse)
			})
		})
	})

	Convey("Given I create a new PushFilter with no type", t, func() {

		f := NewPushFilter()

		Convey("When I check if OperationCreate is filtered", func() {

			filtered := f.IsEventTypeFiltered(EventCreate)

			Convey("Then filtered should be false", func() {
				So(filtered, ShouldBeFalse)
			})
		})
	})
}

func TestPushFilter_IsIdentityFiltered(t *testing.T) {

	Convey("Given I create a new PushFilter", t, func() {

		f := NewPushFilter()
		i1 := MakeIdentity("i1", "i1")
		i2 := MakeIdentity("i2", "i2")

		f.Identities = []Identity{i1, i2}

		Convey("When I check if i3 is filtered", func() {

			filtered := f.IsIdentityFiltered("i3")

			Convey("Then filtered should be true", func() {
				So(filtered, ShouldBeTrue)
			})
		})

		Convey("When I check if i2 is filtered", func() {

			filtered := f.IsIdentityFiltered("i2")

			Convey("Then filtered should be false", func() {
				So(filtered, ShouldBeFalse)
			})
		})
	})

	Convey("Given I create a new PushFilter with no identities", t, func() {

		f := NewPushFilter()

		Convey("When I check if i3 is filtered", func() {

			filtered := f.IsIdentityFiltered("i3")

			Convey("Then filtered should be false", func() {
				So(filtered, ShouldBeFalse)
			})
		})
	})
}

func TestPushFilter_String(t *testing.T) {

	Convey("Given I create a new PushFilter", t, func() {

		f := NewPushFilter()
		i1 := MakeIdentity("i1", "i1")
		i2 := MakeIdentity("i2", "i2")

		f.Identities = []Identity{i1, i2}
		f.EventTypes = []EventType{EventCreate, EventUpdate}
		f.RequestID = "toto"

		Convey("When I call the String Method", func() {
			s := f.String()

			Convey("Then it should be correctly printed", func() {
				So(s, ShouldEqual, "<pushfilter id:toto event-types:[create update] identities:[<Identity i1|i1> <Identity i2|i2>]>")
			})
		})
	})
}
