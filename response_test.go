package elemental

import (
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
