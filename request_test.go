package elemental

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type brokenReader struct{}

func (brokenReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("nope")
}

func TestRequest_NewRequest(t *testing.T) {

	Convey("Given I create a new request", t, func() {
		r := NewRequest()

		Convey("Then it should be correctly initialized", func() {
			So(r.RequestID, ShouldNotBeEmpty)
		})
	})
}

func TestRequest_EncodeDecode(t *testing.T) {

	Convey("Given I create a new request", t, func() {
		r := NewRequest()

		Convey("When I encode an object into the request", func() {

			o := &List{ID: "1", Name: "hello"}
			err := r.Encode(o)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then data should not be nil", func() {
				So(len(r.Data), ShouldNotBeEmpty)
			})

			Convey("When I Decode it", func() {
				o1 := &List{}

				err := r.Decode(&o1)

				Convey("Then err should be nil", func() {
					So(err, ShouldBeNil)
				})

				Convey("Then o2 should resemble to o", func() {
					So(o1, ShouldResemble, o)
				})
			})
		})

		Convey("When I encode an unmarshallable object into the request", func() {

			o := &UnmarshalableList{}
			err := r.Encode(o)

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then data should be empty", func() {
				So(len(r.Data), ShouldEqual, 0)
			})

			Convey("When I Decode it", func() {
				o1 := &List{}

				err := r.Decode(&o1)

				Convey("Then err should not be nil", func() {
					So(err, ShouldNotBeNil)
				})
			})
		})
	})
}

func TestRequest_FromHttp(t *testing.T) {

	Convey("Given I have a get http request on /lists", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/v/10/lists?p=v&page=1&pagesize=2&recursive=true&override=true", nil)
		req.Header.Set("X-Namespace", "ns")
		req.RemoteAddr = "42.42.42.42"

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationRetrieveMany", func() {
				So(r.Operation, ShouldEqual, OperationRetrieveMany)
			})

			Convey("Then the version should be 10", func() {
				So(r.Version, ShouldEqual, 10)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			// Convey("Then the parameters should be correct", func() {
			// 	So(r.Parameters, ShouldResemble, req.URL.Query())
			// })

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be empty", func() {
				So(r.Username, ShouldBeEmpty)
			})

			Convey("Then the Password should be empty", func() {
				So(r.Password, ShouldBeEmpty)
			})

			Convey("Then the Data should be empty", func() {
				So(r.Data, ShouldBeEmpty)
			})

			Convey("Then the ClientIP should be set", func() {
				So(r.ClientIP, ShouldEqual, "42.42.42.42")
			})

			Convey("Then the Page should be 1", func() {
				So(r.Page, ShouldEqual, 1)
			})

			Convey("Then the PageSize should be 2", func() {
				So(r.PageSize, ShouldEqual, 2)
			})

			Convey("Then the Recursive should be true", func() {
				So(r.Recursive, ShouldBeTrue)
			})

			Convey("Then the OverrideProtection should be true", func() {
				So(r.OverrideProtection, ShouldBeTrue)
			})

			Convey("Then I can retrieve the original request", func() {
				So(r.HTTPRequest(), ShouldEqual, req)
			})
		})
	})

	Convey("Given I have a head http request on /lists", t, func() {

		req, _ := http.NewRequest(http.MethodHead, "http://server/lists?p=v", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationInfo", func() {
				So(r.Operation, ShouldEqual, OperationInfo)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			// Convey("Then the parameters should be correct", func() {
			// 	So(r.Parameters, ShouldResemble, req.URL.Query())
			// })

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be empty", func() {
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a patch http request on /lists", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPatch, "http://server/lists?p=v", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationPatch", func() {
				So(r.Operation, ShouldEqual, OperationPatch)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			// Convey("Then the parameters should be correct", func() {
			// 	So(r.Parameters, ShouldResemble, req.URL.Query())
			// })

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(string(r.Data), ShouldEqual, `{"name": "toto"}`)
			})

			Convey("Then the order should be nil", func() {
				So(r.Order, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a post http request on /lists", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPost, "http://server/lists?p=v&order=name&order=toto", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationCreate", func() {
				So(r.Operation, ShouldEqual, OperationCreate)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			Convey("Then the order should be correct", func() {
				So(r.Order, ShouldResemble, []string{"name", "toto"})
			})

			// Convey("Then the parameters should be correct", func() {
			// 	So(r.Parameters, ShouldResemble, req.URL.Query())
			// })

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(string(r.Data), ShouldEqual, `{"name": "toto"}`)
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx?p=v", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationRetrieve", func() {
				So(r.Operation, ShouldEqual, OperationRetrieve)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			// Convey("Then the parameters should be correct", func() {
			// 	So(r.Parameters, ShouldResemble, req.URL.Query())
			// })

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be xx", func() {
				So(r.ObjectID, ShouldEqual, "xx")
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a put http request on /lists/xx", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPut, "http://server/lists/xx?p=v", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationUpdate", func() {
				So(r.Operation, ShouldEqual, OperationUpdate)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			// Convey("Then the parameters should be correct", func() {
			// 	So(r.Parameters, ShouldResemble, req.URL.Query())
			// })

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be xx", func() {
				So(r.ObjectID, ShouldEqual, "xx")
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(string(r.Data), ShouldEqual, `{"name": "toto"}`)
			})
		})
	})

	Convey("Given I have a delete http request on /lists/xx", t, func() {

		req, _ := http.NewRequest(http.MethodDelete, "http://server/lists/xx?p=v", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationDelete", func() {
				So(r.Operation, ShouldEqual, OperationDelete)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			// Convey("Then the parameters should be correct", func() {
			// 	So(r.Parameters, ShouldResemble, req.URL.Query())
			// })

			Convey("Then the identity should be ListIdentity", func() {
				So(r.Identity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ObjectID should be xx", func() {
				So(r.ObjectID, ShouldEqual, "xx")
			})

			Convey("Then the parent identity should be empty", func() {
				So(r.ParentIdentity, ShouldResemble, EmptyIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldBeEmpty)
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx/tasks", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/tasks?p=v", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should not be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the operation should be OperationRetrieveMany", func() {
				So(r.Operation, ShouldEqual, OperationRetrieveMany)
			})

			Convey("Then the Namespace should be ns", func() {
				So(r.Namespace, ShouldResemble, "ns")
			})

			Convey("Then the RequestID should not be empty", func() {
				So(r.RequestID, ShouldNotBeEmpty)
			})

			// Convey("Then the parameters should be correct", func() {
			// 	So(r.Parameters, ShouldResemble, req.URL.Query())
			// })

			Convey("Then the identity should be TaskIdentity", func() {
				So(r.Identity, ShouldResemble, TaskIdentity)
			})

			Convey("Then the ObjectID should be empty", func() {
				So(r.ObjectID, ShouldBeEmpty)
			})

			Convey("Then the parent identity should be ListIdentity", func() {
				So(r.ParentIdentity, ShouldResemble, ListIdentity)
			})

			Convey("Then the ParentID should be empty", func() {
				So(r.ParentID, ShouldEqual, "xx")
			})

			Convey("Then the Username should be user", func() {
				So(r.Username, ShouldEqual, "user")
			})

			Convey("Then the Password should be pass", func() {
				So(r.Password, ShouldEqual, "pass")
			})

			Convey("Then the Data should be correct", func() {
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a patch http request with a brokenReader ", t, func() {

		req, _ := http.NewRequest(http.MethodPatch, "http://server/lists/xx/tasks?p=v", brokenReader{})

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a post http request with a brokenReader ", t, func() {

		req, _ := http.NewRequest(http.MethodPost, "http://server/lists/xx/tasks?p=v", brokenReader{})

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a put http request with a brokenReader ", t, func() {

		req, _ := http.NewRequest(http.MethodPut, "http://server/lists/xx/tasks?p=v", brokenReader{})

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a http request with no url ", t, func() {

		req, _ := http.NewRequest(http.MethodPut, "", nil)

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a http request with a page that is not a number", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/tasks?page=not-int", nil)

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a http request with a pagesize that is not a number", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/tasks?pagesize=not-int", nil)

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a http request with a version that is not a number", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/v/A/lists/xx/tasks", nil)

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("Given I have a http request with a bad path", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/tasks/yy/what?", nil)

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestRequest_Duplicate(t *testing.T) {

	Convey("Given I have a Request", t, func() {

		req := NewRequest()
		req.Data = []byte(`{"hello": "world"}`)
		req.Headers = http.Header{"x-h1": []string{"hey"}}
		req.Identity = UserIdentity
		req.Namespace = "ns"
		req.ObjectID = "xxx"
		req.Operation = OperationPatch
		req.Page = 1
		req.PageSize = 2
		// req.Parameters = url.Values{"p1": []string{"v1"}}
		req.ParentID = "zzz"
		req.ParentIdentity = TaskIdentity
		req.Password = "pass"
		req.Recursive = true
		req.Username = "user"
		req.Version = 12
		req.Order = []string{"key1", "key2"}
		req.ClientIP = "1.2.3.4"
		req.Metadata = map[string]interface{}{"a": 1}

		Convey("When I use Duplicate()", func() {

			req2 := req.Duplicate()

			Convey("Then the duplicated request should be correct", func() {
				So(req2.Data, ShouldResemble, req.Data)
				So(req2.Headers.Get("x-h1"), ShouldEqual, req.Headers.Get("x-h1"))
				So(req2.Identity.IsEqual(req.Identity), ShouldBeTrue)
				So(req2.Namespace, ShouldEqual, req.Namespace)
				So(req2.ObjectID, ShouldEqual, req.ObjectID)
				So(req2.Operation, ShouldEqual, req.Operation)
				So(req2.Page, ShouldEqual, req.Page)
				So(req2.PageSize, ShouldEqual, req.PageSize)
				// So(req2.Parameters.Get("p1"), ShouldEqual, req.Parameters.Get("p1"))
				So(req2.ParentID, ShouldEqual, req.ParentID)
				So(req2.ParentIdentity.IsEqual(req.ParentIdentity), ShouldBeTrue)
				So(req2.Password, ShouldEqual, req.Password)
				So(req2.Recursive, ShouldEqual, req.Recursive)
				So(req2.Username, ShouldEqual, req.Username)
				So(req2.RequestID, ShouldNotEqual, req.RequestID)
				So(req2.Version, ShouldEqual, req.Version)
				So(req2.Order, ShouldResemble, req.Order)
				So(req2.ClientIP, ShouldResemble, req.ClientIP)
				So(req2.Metadata, ShouldResemble, req.Metadata)
			})
		})
	})
}
