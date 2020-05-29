package elemental

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefaultNamespacer(t *testing.T) {
	Convey("Given a default namespaer", t, func() {
		d := &defaultNamespacer{}

		Convey("When I pass a request with the namespace parameter, it should return the right value", func() {
			r := &http.Request{
				Header: http.Header{},
			}
			r.Header.Add("X-Namespace", "mynamespace")
			result, err := d.Extract(r)
			So(err, ShouldBeNil)
			So(result, ShouldResemble, "mynamespace")
		})

		Convey("When I set the namespace to a value, Inject should add the right header", func() {
			r := &http.Request{}
			err := d.Inject(r, "injectednamespace")
			So(err, ShouldBeNil)

			So(r.Header.Get("X-Namespace"), ShouldResemble, "injectednamespace")
		})
	})
}

func TestSetNamespacer(t *testing.T) {
	Convey("When I set the namespacer it should take effect", t, func() {
		newNamespacer := &defaultNamespacer{}

		SetNamespacer(newNamespacer)
		So(namespacer, ShouldEqual, newNamespacer)
	})
}
