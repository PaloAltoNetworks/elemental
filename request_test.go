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
	"bytes"
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type brokenReader struct{}

// erroredNamespacer is a namespacer that returns for testing.
type erroredNamespacer struct{}

// Extract implements the corresponding method of the interface.
func (d *erroredNamespacer) Extract(r *http.Request) (string, error) {
	return "", fmt.Errorf("Test Error")
}

// Inject implements the corresponding method of the interface.
func (d *erroredNamespacer) Inject(r *http.Request, namespace string) error {
	return fmt.Errorf("Test Error Inject")
}

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

func TestRequest_NewRequestFromHTTPRequest(t *testing.T) {

	Convey("Given I have a get http request on /lists with page", t, func() {

		req, err := http.NewRequest(http.MethodGet, "http://server/v/10/lists?page=1&pagesize=2&recursive=true&override=true&propagated=true&rlgmp1=A&rlgmp2=true", nil)
		req.Header.Set("X-Namespace", "ns")
		req.Header.Set("X-Forwarded-for", "1.1.1.1, 5.5.5.5, 10.10.10.10")
		req.Header.Set("X-Real-IP", "2.2.2.2")
		req.Header.Set("Accept", "application/msgpack")
		req.Header.Set("Content-Type", "application/json")
		req.RemoteAddr = "42.42.42.42"

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(r, ShouldNotBeNil)
			})

			Convey("Then r should be correct", func() {
				So(err, ShouldBeNil)
				So(r.Operation, ShouldEqual, OperationRetrieveMany)
				So(r.Version, ShouldEqual, 10)
				So(r.Namespace, ShouldResemble, "ns")
				So(r.RequestID, ShouldNotBeEmpty)
				So(r.Identity, ShouldResemble, ListIdentity)
				So(r.ObjectID, ShouldBeEmpty)
				So(r.ParentIdentity, ShouldResemble, RootIdentity)
				So(r.ParentID, ShouldBeEmpty)
				So(r.Username, ShouldBeEmpty)
				So(r.Password, ShouldBeEmpty)
				So(r.Data, ShouldBeEmpty)
				So(r.ClientIP, ShouldEqual, "1.1.1.1")
				So(r.Page, ShouldEqual, 1)
				So(r.PageSize, ShouldEqual, 2)
				So(r.After, ShouldEqual, "")
				So(r.Limit, ShouldEqual, 0)
				So(r.Recursive, ShouldBeTrue)
				So(r.Propagated, ShouldBeTrue)
				So(r.OverrideProtection, ShouldBeTrue)
				So(r.Accept, ShouldEqual, EncodingTypeMSGPACK)
				So(r.ContentType, ShouldEqual, EncodingTypeJSON)
				So(r.HTTPRequest(), ShouldEqual, req)
			})
		})
	})

	Convey("Given I have a get http request on /lists with after", t, func() {

		req, err := http.NewRequest(http.MethodGet, "http://server/v/10/lists?after=42&limit=2&recursive=true&override=true&rlgmp1=A&rlgmp2=true", nil)
		req.Header.Set("X-Namespace", "ns")
		req.Header.Set("X-Forwarded-for", "1.1.1.1")
		req.Header.Set("X-Real-IP", "2.2.2.2")
		req.Header.Set("Accept", "application/msgpack")
		req.Header.Set("Content-Type", "application/json")
		req.RemoteAddr = "42.42.42.42"

		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationRetrieveMany)
				So(r.Version, ShouldEqual, 10)
				So(r.Namespace, ShouldResemble, "ns")
				So(r.RequestID, ShouldNotBeEmpty)
				So(r.Identity, ShouldResemble, ListIdentity)
				So(r.ObjectID, ShouldBeEmpty)
				So(r.ParentIdentity, ShouldResemble, RootIdentity)
				So(r.ParentID, ShouldBeEmpty)
				So(r.Username, ShouldBeEmpty)
				So(r.Password, ShouldBeEmpty)
				So(r.Data, ShouldBeEmpty)
				So(r.ClientIP, ShouldEqual, "1.1.1.1")
				So(r.PageSize, ShouldEqual, 0)
				So(r.After, ShouldEqual, "42")
				So(r.Limit, ShouldEqual, 2)
				So(r.Recursive, ShouldBeTrue)
				So(r.Propagated, ShouldBeFalse)
				So(r.OverrideProtection, ShouldBeTrue)
				So(r.Accept, ShouldEqual, EncodingTypeMSGPACK)
				So(r.ContentType, ShouldEqual, EncodingTypeJSON)
				So(r.HTTPRequest(), ShouldEqual, req)
			})
		})
	})

	Convey("Given I have a head http request on /lists", t, func() {

		req, _ := http.NewRequest(http.MethodHead, "http://server/lists?rlgmp1=A&rlgmp2=true", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")
		req.Header.Set("X-Namespace", "ns")
		req.Header.Set("X-Real-IP", "2.2.2.2")
		req.RemoteAddr = "42.42.42.42"

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationInfo)
				So(r.Namespace, ShouldResemble, "ns")
				So(r.RequestID, ShouldNotBeEmpty)
				So(r.Identity, ShouldResemble, ListIdentity)
				So(r.ObjectID, ShouldBeEmpty)
				So(r.ParentIdentity, ShouldResemble, RootIdentity)
				So(r.ParentID, ShouldBeEmpty)
				So(r.Username, ShouldEqual, "user")
				So(r.Password, ShouldEqual, "pass")
				So(r.Data, ShouldBeEmpty)
				So(r.ClientIP, ShouldEqual, "2.2.2.2")
			})
		})
	})

	Convey("Given I have a patch http request on /lists", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPatch, "http://server/lists?lup1=A&lup2=true", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationPatch)
				So(r.Namespace, ShouldResemble, "ns")
				So(r.RequestID, ShouldNotBeEmpty)
				So(r.Identity, ShouldResemble, ListIdentity)
				So(r.ObjectID, ShouldBeEmpty)
				So(r.ParentIdentity, ShouldResemble, RootIdentity)
				So(r.ParentID, ShouldBeEmpty)
				So(r.Username, ShouldEqual, "user")
				So(r.Password, ShouldEqual, "pass")
				So(string(r.Data), ShouldEqual, `{"name": "toto"}`)
				So(r.Order, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a patch http request on /lists using multipart/form-data", t, func() {

		RegisterSupportedContentType("multipart/form-data")
		defer func() { externalSupportedAcceptType = map[string]struct{}{} }()

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPatch, "http://server/lists?lup1=A&lup2=true", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")
		req.Header.Add("Content-Type", "multipart/form-data; boundary=x")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationPatch)
				So(string(r.Data), ShouldEqual, "") // should not be decoded
			})
		})
	})

	Convey("Given I have a post http request on /lists", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPost, "http://server/lists?order=name&order=toto&rlcp1=A&rlcp2=true", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationCreate)
				So(r.Namespace, ShouldResemble, "ns")
				So(r.RequestID, ShouldNotBeEmpty)
				So(r.Order, ShouldResemble, []string{"name", "toto"})
				So(r.Identity, ShouldResemble, ListIdentity)
				So(r.ObjectID, ShouldBeEmpty)
				So(r.ParentIdentity, ShouldResemble, RootIdentity)
				So(r.ParentID, ShouldBeEmpty)
				So(r.Username, ShouldEqual, "user")
				So(r.Password, ShouldEqual, "pass")
				So(string(r.Data), ShouldEqual, `{"name": "toto"}`)
			})
		})
	})

	Convey("Given I have a post http request on /lists and the namespacer returns an error", t, func() {

		SetNamespacer(&erroredNamespacer{})

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPost, "http://server/lists?order=name&order=toto&rlcp1=A&rlcp2=true", buffer)
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			_, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldResemble, "Test Error")
			})
		})

		SetNamespacer(&defaultNamespacer{})
	})

	Convey("Given I have a post http request on /lists and the namespacer is nil", t, func() {

		SetNamespacer(nil)

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPost, "http://server/lists?order=name&order=toto&rlcp1=A&rlcp2=true", buffer)

		Convey("When I create a new elemental Request from it", func() {

			_, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})
		})

		SetNamespacer(&defaultNamespacer{})
	})

	Convey("Given I have a post http request on /lists using multipart/form-data", t, func() {

		RegisterSupportedContentType("multipart/form-data")
		defer func() { externalSupportedAcceptType = map[string]struct{}{} }()

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPost, "http://server/lists?order=name&order=toto&rlcp1=A&rlcp2=true", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")
		req.Header.Add("Content-Type", "multipart/form-data; boundary=x")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationCreate)

				So(string(r.Data), ShouldEqual, ``)
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx?lgp1=A&lgp2=true", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationRetrieve)
				So(r.Namespace, ShouldResemble, "ns")
				So(r.RequestID, ShouldNotBeEmpty)
				So(r.Identity, ShouldResemble, ListIdentity)
				So(r.ObjectID, ShouldEqual, "xx")
				So(r.ParentIdentity, ShouldResemble, RootIdentity)
				So(r.ParentID, ShouldBeEmpty)
				So(r.Username, ShouldEqual, "user")
				So(r.Password, ShouldEqual, "pass")
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a put http request on /lists/xx", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPut, "http://server/lists/xx?lup1=A&lup2=true", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationUpdate)
				So(r.Namespace, ShouldResemble, "ns")
				So(r.RequestID, ShouldNotBeEmpty)
				So(r.Identity, ShouldResemble, ListIdentity)
				So(r.ObjectID, ShouldEqual, "xx")
				So(r.ParentIdentity, ShouldResemble, RootIdentity)
				So(r.ParentID, ShouldBeEmpty)
				So(r.Username, ShouldEqual, "user")
				So(r.Password, ShouldEqual, "pass")
				So(string(r.Data), ShouldEqual, `{"name": "toto"}`)
			})
		})
	})

	Convey("Given I have a put http request on /lists/xx using multipart/form-data", t, func() {

		RegisterSupportedContentType("multipart/form-data")
		defer func() { externalSupportedAcceptType = map[string]struct{}{} }()

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPut, "http://server/lists/xx?lup1=A&lup2=true", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")
		req.Header.Add("Content-Type", "multipart/form-data")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationUpdate)
				So(string(r.Data), ShouldEqual, ``)
			})
		})
	})

	Convey("Given I have a delete http request on /lists/xx", t, func() {

		req, _ := http.NewRequest(http.MethodDelete, "http://server/lists/xx?ldp1=A&ldp2=true", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationDelete)
				So(r.Namespace, ShouldResemble, "ns")
				So(r.RequestID, ShouldNotBeEmpty)
				So(r.Identity, ShouldResemble, ListIdentity)
				So(r.ObjectID, ShouldEqual, "xx")
				So(r.ParentIdentity, ShouldResemble, RootIdentity)
				So(r.ParentID, ShouldBeEmpty)
				So(r.Username, ShouldEqual, "user")
				So(r.Password, ShouldEqual, "pass")
				So(r.Data, ShouldBeEmpty)
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx/tasks", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/tasks?ltgp1=A&ltgp2=true", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then r should be correct", func() {
				So(r.Operation, ShouldEqual, OperationRetrieveMany)
				So(r.Namespace, ShouldResemble, "ns")
				So(r.RequestID, ShouldNotBeEmpty)
				So(r.Identity, ShouldResemble, TaskIdentity)
				So(r.ObjectID, ShouldBeEmpty)
				So(r.ParentIdentity, ShouldResemble, ListIdentity)
				So(r.ParentID, ShouldEqual, "xx")
				So(r.Username, ShouldEqual, "user")
				So(r.Password, ShouldEqual, "pass")
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

	Convey("Given I have a post http request invalid accept ", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx", nil)
		req.Header.Set("Accept", "dfsdfsd sdfsdfsdf")

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

	Convey("Given I have a post http request invalid content-type ", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/", nil)
		req.Header.Set("Content-Type", "dfsdfsd sdfsdfsdf")

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

	Convey("Given I have a http request %00 as order parameter ", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/?order=%00", nil)
		req.Header.Set("Content-Type", "application/json")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 400 (elemental): Bad Request: Parameter `order` must be set when provided")
			})
		})
	})

	Convey("Given I have a http request %00 as page parameter", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/?page=%00", nil)
		req.Header.Set("Content-Type", "application/json")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 400 (elemental): Bad Request: Parameter `page` must be an integer")
			})
		})
	})

	Convey("Given I have a http request %00 as pagesize parameter", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/?pagesize=%00", nil)
		req.Header.Set("Content-Type", "application/json")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 400 (elemental): Bad Request: Parameter `pagesize` must be an integer")
			})
		})
	})

	Convey("Given I have a http request %00 as after parameter ", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/?after=%00", nil)
		req.Header.Set("Content-Type", "application/json")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 400 (elemental): Bad Request: Parameter `after` must be set when provided")
			})
		})
	})

	Convey("Given I have a http request with after and 2 order parameters", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/?after=xxx&order=a&order=b", nil)
		req.Header.Set("Content-Type", "application/json")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 400 (elemental): Bad Request: You can only order on a single field when using 'after'")
			})
		})
	})

	Convey("Given I have a http request with after and page parameters", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/?after=xxx&page=2", nil)
		req.Header.Set("Content-Type", "application/json")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 400 (elemental): Bad Request: You cannot set 'after' and 'page' at the same time")
			})
		})
	})

	Convey("Given I have a http request with pagesize and limit parameters", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/?limit=xxx&pagesize=2", nil)
		req.Header.Set("Content-Type", "application/json")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "error 400 (elemental): Bad Request: You cannot set 'limit' and 'pagesize' at the same time")
			})
		})
	})

	Convey("Given I have a http request with a limit that is not a number", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/tasks?limit=not-int", nil)

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
		req.Limit = 2
		req.After = "after"
		req.Parameters = Parameters{
			"p1": Parameter{
				ptype:  ParameterTypeString,
				values: []any{"A"},
			},
			"p2": Parameter{
				ptype:  ParameterTypeBool,
				values: []any{true},
			},
		}
		req.ParentID = "zzz"
		req.ParentIdentity = TaskIdentity
		req.Password = "pass"
		req.Recursive = true
		req.Propagated = true
		req.Username = "user"
		req.Version = 12
		req.Order = []string{"key1", "key2"}
		req.ClientIP = "1.2.3.4"
		req.Metadata = map[string]any{"a": 1}
		req.ContentType = EncodingTypeMSGPACK
		req.Accept = EncodingTypeMSGPACK

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
				So(req2.PageSize, ShouldEqual, req.PageSize)
				So(req2.Parameters, ShouldResemble, req.Parameters)
				So(req2.ParentID, ShouldEqual, req.ParentID)
				So(req2.ParentIdentity.IsEqual(req.ParentIdentity), ShouldBeTrue)
				So(req2.Password, ShouldEqual, req.Password)
				So(req2.Recursive, ShouldEqual, req.Recursive)
				So(req2.Propagated, ShouldEqual, req.Propagated)
				So(req2.Username, ShouldEqual, req.Username)
				So(req2.RequestID, ShouldNotEqual, req.RequestID)
				So(req2.Version, ShouldEqual, req.Version)
				So(req2.Order, ShouldResemble, req.Order)
				So(req2.ClientIP, ShouldResemble, req.ClientIP)
				So(req2.Metadata, ShouldResemble, req.Metadata)
				So(req2.ContentType, ShouldEqual, req.ContentType)
				So(req2.Accept, ShouldEqual, req.Accept)
				So(req2.After, ShouldEqual, req.After)
				So(req2.Limit, ShouldEqual, req.Limit)
			})
		})
	})
}

func TestRequest_NewRequestFromHTTPRequestParameters(t *testing.T) {

	Convey("Given I have a get http request on /lists with good params", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/v/10/lists?page=1&pagesize=2&recursive=true&override=true&rlgmp1=A&rlgmp2=true", nil)
		req.Header.Set("X-Namespace", "ns")
		req.Header.Set("X-Real-IP", "2.2.2.2")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, Parameters{
					"rlgmp1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"A"},
					},
					"rlgmp2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
				})
			})
		})
	})

	Convey("Given I have a get http request on /lists with not good params", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/v/10/lists?page=1&pagesize=2&recursive=true&override=true&rlgmp1=A&rlgmp2=notbool", nil)
		req.Header.Set("X-Namespace", "ns")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Invalid Parameter: Parameter 'rlgmp2' must be a boolean, got 'notbool'`)
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a head http request on /lists with good params", t, func() {

		req, _ := http.NewRequest(http.MethodHead, "http://server/lists?rlgmp1=A&rlgmp2=true", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, Parameters{
					"rlgmp1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"A"},
					},
					"rlgmp2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
				})
			})
		})
	})

	Convey("Given I have a head http request on /lists with not good params", t, func() {

		req, _ := http.NewRequest(http.MethodHead, "http://server/lists?rlgmp1=A&rlgmp2=nottrue", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Invalid Parameter: Parameter 'rlgmp2' must be a boolean, got 'nottrue'`)
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a patch http request on /lists with good params", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPatch, "http://server/lists?lup1=A&lup2=true", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, Parameters{
					"lup1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"A"},
					},
					"lup2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
				})
			})
		})
	})

	Convey("Given I have a patch http request on /lists with not good params", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPatch, "http://server/lists?lup1=A&lup2=nottrue", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Invalid Parameter: Parameter 'lup2' must be a boolean, got 'nottrue'`)
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a post http request on /lists with good params", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPost, "http://server/lists?order=name&order=toto&rlcp1=A&rlcp2=true", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, Parameters{
					"rlcp1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"A"},
					},
					"rlcp2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
				})
			})
		})
	})

	Convey("Given I have a post http request on /lists with not good params", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPost, "http://server/lists?order=name&order=toto&rlcp1=A&rlcp2=nottrue", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Invalid Parameter: Parameter 'rlcp2' must be a boolean, got 'nottrue'`)
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx with good params", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx?lgp1=A&lgp2=true&sAp1=ok&sAp2=true&sBp1=ok&sBp2=true", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, Parameters{
					"lgp1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"A"},
					},
					"lgp2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
					"sAp1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"ok"},
					},
					"sAp2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
					"sBp1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"ok"},
					},
					"sBp2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
				})
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx with not good params", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx?lgp1=A&lgp2=nottrue", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Invalid Parameter: Parameter 'lgp2' must be a boolean, got 'nottrue'`)
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a put http request on /lists/xx with good params", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPut, "http://server/lists/xx?lup1=A&lup2=true", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, Parameters{
					"lup1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"A"},
					},
					"lup2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
				})
			})
		})
	})

	Convey("Given I have a put http request on /lists/xx with not good params", t, func() {

		buffer := bytes.NewBuffer([]byte(`{"name": "toto"}`))
		req, _ := http.NewRequest(http.MethodPut, "http://server/lists/xx?lup1=A&lup2=nottrue", buffer)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Invalid Parameter: Parameter 'lup2' must be a boolean, got 'nottrue'`)
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a delete http request on /lists/xx with good paramts", t, func() {

		req, _ := http.NewRequest(http.MethodDelete, "http://server/lists/xx?ldp1=A&ldp2=true", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, Parameters{
					"ldp1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"A"},
					},
					"ldp2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
				})
			})
		})
	})

	Convey("Given I have a delete http request on /lists/xx with not good paramts", t, func() {

		req, _ := http.NewRequest(http.MethodDelete, "http://server/lists/xx?ldp1=A&ldp2=nottrue", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Invalid Parameter: Parameter 'ldp2' must be a boolean, got 'nottrue'`)
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx/tasks with good params", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/tasks?ltgp1=A&ltgp2=true", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, Parameters{
					"ltgp1": Parameter{
						ptype:  ParameterTypeString,
						values: []any{"A"},
					},
					"ltgp2": Parameter{
						ptype:  ParameterTypeBool,
						values: []any{true},
					},
				})
			})
		})
	})

	Convey("Given I have a get http request on /lists/xx/tasks with not good params", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/lists/xx/tasks?ltgp1=A&ltgp2=nottrue", nil)
		req.Header.Add("X-Namespace", "ns")
		req.Header.Add("Authorization", "user pass")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, `error 400 (elemental): Invalid Parameter: Parameter 'ltgp2' must be a boolean, got 'nottrue'`)
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a get http request on /lists with no params", t, func() {

		req, _ := http.NewRequest(http.MethodGet, "http://server/v/10/lists", nil)
		req.Header.Set("X-Namespace", "ns")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the parameters should be correct", func() {
				So(r.Parameters, ShouldResemble, Parameters{
					"rlgmp1": Parameter{
						ptype:  ParameterTypeString,
						values: nil,
					},
					"rlgmp2": Parameter{
						ptype:  ParameterTypeBool,
						values: nil,
					},
				})
				So(r.Parameters["rlgmp1"].StringValue(), ShouldEqual, "")
				So(r.Parameters["rlgmp2"].BoolValue(), ShouldEqual, false)
			})
		})
	})
}

func TestRequest_RequiredParameters(t *testing.T) {

	Convey("Given I have a delete http request on /users with no params", t, func() {

		req, _ := http.NewRequest(http.MethodDelete, "http://server/v/10/users/id", nil)
		req.Header.Set("X-Namespace", "ns")
		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})

	Convey("Given I have a http request on /lists with unknown params", t, func() {

		req, _ := http.NewRequest(http.MethodDelete, "http://server/v/10/lists?what=1&the=0", nil)
		req.Header.Set("X-Namespace", "ns")

		Convey("When I create a new elemental Request from it", func() {

			r, err := NewRequestFromHTTPRequest(req, Manager())

			Convey("Then err should not be nil", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "error 400 (elemental): Bad Request: Unknown parameter: `what`")
				So(err.Error(), ShouldContainSubstring, "error 400 (elemental): Bad Request: Unknown parameter: `the`")
			})

			Convey("Then r should be nil", func() {
				So(r, ShouldBeNil)
			})
		})
	})
}

func TestDecode(t *testing.T) {

	Convey("Given I have a list and a request with Content-Type JSON", t, func() {

		lst := NewList()
		lst.Name = "hello"
		data, _ := Encode(EncodingTypeJSON, lst)

		req := &Request{
			ContentType: EncodingTypeJSON,
			Data:        data,
		}

		Convey("When I call Decode", func() {

			lst2 := NewList()
			err := req.Decode(lst2)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the object should be correctly decoded", func() {

				So(lst2, ShouldResemble, lst)
			})
		})
	})

	Convey("Given I have a list and a request with Content-Type MSGPACK", t, func() {

		lst := NewList()
		lst.Name = "hello"
		data, _ := Encode(EncodingTypeMSGPACK, lst)

		req := &Request{
			ContentType: EncodingTypeMSGPACK,
			Data:        data,
		}

		Convey("When I call Decode", func() {

			lst2 := NewList()
			err := req.Decode(lst2)

			Convey("Then err should be nil", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the object should be correctly decoded", func() {

				So(lst2, ShouldResemble, lst)
			})
		})
	})

}
