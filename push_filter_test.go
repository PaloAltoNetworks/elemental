package elemental

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPushFilter_NewPushFilter(t *testing.T) {

	Convey("Given I create a new PushFilter", t, func() {
		f := NewPushFilter()

		Convey("Then it should be correctly initialized", func() {
			So(f.Identities, ShouldNotBeNil)
		})
	})
}

func TestPushFilter_Duplicate(t *testing.T) {

	Convey("Given I create a new PushFilter", t, func() {

		f := NewPushFilter()

		f.FilterIdentity("i1", EventCreate, EventDelete)
		f.FilterIdentity("i2", EventCreate, EventDelete)

		Convey("When I duplicate it", func() {
			dup := f.Duplicate()

			Convey("Then it should be correctly duplicated", func() {
				So(dup.Identities, ShouldResemble, f.Identities)
				So(dup.Identities, ShouldNotEqual, f.Identities)
			})
		})
	})
}

func TestPushFilter_IsFilteredOut(t *testing.T) {

	Convey("Given I create a new PushFilter", t, func() {

		f := NewPushFilter()

		Convey("When I check if i1 is filtered", func() {

			filtered1 := f.IsFilteredOut("i1", EventDelete)
			filtered2 := f.IsFilteredOut("i2", EventDelete)

			Convey("Then filtered1 should be false", func() {
				So(filtered1, ShouldBeFalse)
			})

			Convey("Then filtered2 should be false", func() {
				So(filtered2, ShouldBeFalse)
			})
		})

		Convey("When I add a filter for i1 on Create and Delete", func() {

			f.FilterIdentity("i1", EventCreate, EventDelete)
			f.FilterIdentity("i2")

			Convey("Then create and delete should not be filtered out on i1", func() {
				So(f.IsFilteredOut("i1", EventCreate), ShouldBeFalse)
				So(f.IsFilteredOut("i1", EventDelete), ShouldBeFalse)
			})

			Convey("Then update should be filtered out on i1", func() {
				So(f.IsFilteredOut("i1", EventUpdate), ShouldBeTrue)
			})

			Convey("Then nothing should be filtered out on i2", func() {
				So(f.IsFilteredOut("i2", EventCreate), ShouldBeFalse)
				So(f.IsFilteredOut("i2", EventUpdate), ShouldBeFalse)
				So(f.IsFilteredOut("i2", EventDelete), ShouldBeFalse)
			})
		})

		Convey("When I add a filter for i1 on nothing", func() {

			f.FilterIdentity("i1")
			f.FilterIdentity("i2")

			Convey("Then everything should not be filtered out on i1", func() {
				So(f.IsFilteredOut("i1", EventCreate), ShouldBeFalse)
				So(f.IsFilteredOut("i1", EventDelete), ShouldBeFalse)
				So(f.IsFilteredOut("i1", EventUpdate), ShouldBeFalse)
			})

			Convey("Then nothing should be filtered out on i2", func() {
				So(f.IsFilteredOut("i2", EventCreate), ShouldBeFalse)
				So(f.IsFilteredOut("i2", EventUpdate), ShouldBeFalse)
				So(f.IsFilteredOut("i2", EventDelete), ShouldBeFalse)
			})
		})
	})
}

func TestPushFilter_String(t *testing.T) {

	Convey("Given I create a new PushFilter", t, func() {

		f := NewPushFilter()

		f.FilterIdentity("i1", EventCreate, EventDelete)

		Convey("When I call the String Method", func() {
			s := f.String()

			Convey("Then it should be correctly printed", func() {
				So(s, ShouldEqual, "<pushfilter identities:map[i1:[create delete]]>")
			})
		})
	})
}
