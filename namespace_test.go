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

func TestNamespaces_NamespaceAncestorsNames(t *testing.T) {

	Convey("Given I have a namespace", t, func() {

		ns := "/hello/world/wesh/ta/vu"

		Convey("When I call NamespaceAncestorsNames", func() {

			nss := NamespaceAncestorsNames(ns)

			Convey("Then the array should have 5 elements", func() {
				So(len(nss), ShouldEqual, 5)
			})

			Convey("Then the first namespace should be correct", func() {
				So(nss[0], ShouldEqual, "/hello/world/wesh/ta")
			})

			Convey("Then the second namespace should be correct", func() {
				So(nss[1], ShouldEqual, "/hello/world/wesh")
			})

			Convey("Then the third namespace should be correct", func() {
				So(nss[2], ShouldEqual, "/hello/world")
			})

			Convey("Then the fourth namespace should be correct", func() {
				So(nss[3], ShouldEqual, "/hello")
			})

			Convey("Then the fifth namespace should be correct", func() {
				So(nss[4], ShouldEqual, "/")
			})
		})
	})

	Convey("Given I have a / namespace", t, func() {

		ns := "/"

		Convey("When I call NamespaceAncestorsNames", func() {

			nss := NamespaceAncestorsNames(ns)

			Convey("Then the array should have 0 elements", func() {
				So(len(nss), ShouldEqual, 0)
			})
		})
	})

	Convey("Given I have an empty namespace", t, func() {

		ns := ""

		Convey("When I call NamespaceAncestorsNames", func() {

			nss := NamespaceAncestorsNames(ns)

			Convey("Then the array should have 0 elements", func() {
				So(len(nss), ShouldEqual, 0)
			})
		})
	})

	Convey("Given a valid namespace hierarchy /a/b/c", t, func() {

		ns := "/a/b/c"

		Convey("When I try to get the namespace hieararchy ", func() {

			nss := NamespaceAncestorsNames(ns)

			Convey("I should get the right namespaces", func() {
				So(len(nss), ShouldEqual, 3)
			})
		})
	})

	Convey("Given a valid namespace / ", t, func() {

		ns := "/"

		Convey("When I try to get the namespace hieararchy ", func() {

			nss := NamespaceAncestorsNames(ns)

			Convey("I should get empty array ", func() {
				So(len(nss), ShouldEqual, 0)
			})
		})
	})

	Convey("Given a valid namespace /a ", t, func() {

		ns := "/a"

		Convey("When I try to get the namespace hieararchy ", func() {

			nss := NamespaceAncestorsNames(ns)

			Convey("I should get empty array of 1", func() {
				So(len(nss), ShouldEqual, 1)
			})
		})
	})

	Convey("Given a namespace with extra / like /a/  ", t, func() {

		ns := "/a/"

		Convey("When I try to get the namespace hieararchy ", func() {

			nss := NamespaceAncestorsNames(ns)

			Convey("I should get an  array of 2 ", func() {
				So(len(nss), ShouldEqual, 2)
			})
		})
	})
}

func TestNamespaces_ParentNamespaceFromString(t *testing.T) {

	Convey("Given I have a namespace", t, func() {

		ns := "/hello/world"

		Convey("When I call ParentNamespaceFromString", func() {

			s, err := ParentNamespaceFromString(ns)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then s should be correct", func() {
				So(s, ShouldEqual, "/hello")
			})
		})
	})

	Convey("Given I have a / namespace", t, func() {

		ns := "/"

		Convey("When I call ParentNamespaceFromString", func() {

			s, err := ParentNamespaceFromString(ns)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("Then s should be correct", func() {
				So(s, ShouldEqual, "")
			})
		})
	})

	Convey("Given I have a bad namespace", t, func() {

		ns := "asdasdasd"

		Convey("When I call ParentNamespaceFromString", func() {

			s, err := ParentNamespaceFromString(ns)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then s should be correct", func() {
				So(s, ShouldEqual, "")
			})
		})
	})

	Convey("Given I have an empty namespace", t, func() {

		ns := ""

		Convey("When I call ParentNamespaceFromString", func() {

			s, err := ParentNamespaceFromString(ns)

			Convey("Then err should be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then s should be correct", func() {
				So(s, ShouldEqual, "")
			})
		})
	})

	Convey("Given a valid namespace string /parent/child ", t, func() {

		ns := "/parent/child"

		Convey("When I try to get the parent", func() {

			parent, err := ParentNamespaceFromString(ns)

			Convey("Then err should be nil ", func() {
				So(err, ShouldBeNil)
			})

			Convey("It should provide /parent ", func() {
				So(parent, ShouldResemble, "/parent")
			})
		})
	})

	Convey("Given some invalid namespace strings ", t, func() {

		Convey("When I try to get the parent of parent*child ", func() {

			ns := "parent*child"

			Convey("When I try to get the parent", func() {

				parent, err := ParentNamespaceFromString(ns)

				Convey("Then err should not be nil ", func() {
					So(err, ShouldNotBeNil)
				})

				Convey("It should return an empty string", func() {
					So(parent, ShouldResemble, "")
				})
			})
		})

		Convey("When I try to get the parent of /child  ", func() {

			ns := "/child"

			Convey("When I try to get the parent", func() {

				parent, err := ParentNamespaceFromString(ns)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("It should return /", func() {
					So(parent, ShouldResemble, "/")
				})
			})
		})
	})
}

func TestNamespace_IsNamespaceChildrenOfNamespace(t *testing.T) {

	Convey("Given I have a namespace", t, func() {
		ns := "/a/b/c"

		Convey("When I call IsNamespaceChildrenOfNamespace on /a/b", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/a/b")

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /a", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/a")

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/")

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /z", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/z")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /a/c", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, "/a/c")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceChildrenOfNamespace on /a/b/c", func() {

			ok := IsNamespaceChildrenOfNamespace(ns, ns)

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I have an empty namespace", t, func() {

		Convey("When I call IsNamespaceChildrenOfNamespace on empty string", func() {

			ok := IsNamespaceChildrenOfNamespace("", "")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestNamespace_IsNamespaceParentOfNamespace(t *testing.T) {

	Convey("Given I have a namespace", t, func() {
		ns := "/a/b/c"

		Convey("When I call IsNamespaceParentOfNamespace on /a/b", func() {

			ok := IsNamespaceParentOfNamespace("/a/b", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /a", func() {

			ok := IsNamespaceParentOfNamespace("/a", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /", func() {

			ok := IsNamespaceParentOfNamespace("/", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /z", func() {

			ok := IsNamespaceParentOfNamespace(ns, "/z")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /a/c", func() {

			ok := IsNamespaceParentOfNamespace(ns, "/a/c")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceParentOfNamespace on /a/b/c", func() {

			ok := IsNamespaceParentOfNamespace(ns, ns)

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})

	Convey("Given I check if /aa/b is a children of /a", t, func() {

		ok := IsNamespaceParentOfNamespace("/aa/b", "/a")

		Convey("Then ok should be false", func() {
			So(ok, ShouldBeFalse)
		})
	})

	Convey("Given I check if /a is a children of /a", t, func() {

		ok := IsNamespaceParentOfNamespace("/a", "/a")

		Convey("Then ok should be false", func() {
			So(ok, ShouldBeFalse)
		})
	})

	Convey("Given I check if /a is a children of /aa/b", t, func() {

		ok := IsNamespaceParentOfNamespace("/a", "/aa/b")

		Convey("Then ok should be false", func() {
			So(ok, ShouldBeFalse)
		})
	})

	Convey("Given I have an empty namespace", t, func() {
		ns := ""

		Convey("When I call IsNamespaceChildrenOfNamespace on empty string", func() {

			ok := IsNamespaceParentOfNamespace(ns, "/a/b")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}

func TestNamespace_IsNamespaceRelatedToNamesapce(t *testing.T) {

	Convey("Given I have a namespace", t, func() {
		ns := "/a/b/c"

		Convey("When I call IsNamespaceRelatedToNamesapce on /a/b", func() {

			ok := IsNamespaceRelatedToNamespace("/a/b", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /a", func() {

			ok := IsNamespaceRelatedToNamespace("/a", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /", func() {

			ok := IsNamespaceRelatedToNamespace("/", ns)

			Convey("Then ok should be true", func() {
				So(ok, ShouldBeTrue)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /z", func() {

			ok := IsNamespaceRelatedToNamespace(ns, "/z")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /a/c", func() {

			ok := IsNamespaceRelatedToNamespace(ns, "/a/c")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})

		Convey("When I call IsNamespaceRelatedToNamesapce on /a/b/c", func() {

			ok := IsNamespaceRelatedToNamespace(ns, ns)

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeTrue)
			})
		})
	})

	Convey("Given I have an empty namespace", t, func() {

		Convey("When I call IsNamespaceChildrenOfNamespace on empty string", func() {

			ok := IsNamespaceRelatedToNamespace("", "")

			Convey("Then ok should be false", func() {
				So(ok, ShouldBeFalse)
			})
		})
	})
}
