// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type constant string

type Server struct {
	Annotation     map[string]string `json:"annotation"`
	AssociatedTags []string          `json:"associatedTags"`
	ParentType     string            `json:"parentType"`
	UpdatedAt      time.Time         `json:"updatedAt"`
	Number         int               `json:"number"`
	Boom           constant          `json:"boom"`
}

var ta *testing.T

func TestUnmarshalJSONWithNoError(t *testing.T) {

	Convey("Given I call the method UnmarshalJSON with a valid json created with a struct", t, func() {

		server := &Server{}

		expectedServer := &Server{}
		expectedServer.ParentType = "parent"
		expectedServer.UpdatedAt = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		expectedServer.AssociatedTags = []string{"coucou"}
		m := make(map[string]string)
		m["test"] = "ok"
		expectedServer.Annotation = m
		body, _ := json.Marshal(expectedServer)

		err := UnmarshalJSON(bytes.NewReader(body), server)

		So(err, ShouldBeNil)
		So(server.Annotation, ShouldResemble, expectedServer.Annotation)
		So(server.ParentType, ShouldResemble, expectedServer.ParentType)
		So(server.UpdatedAt, ShouldResemble, expectedServer.UpdatedAt)
		So(server.AssociatedTags, ShouldResemble, expectedServer.AssociatedTags)
	})

	Convey("Given I call the method UnmarshalJSON with a valid json created with a string", t, func() {

		server := &Server{}
		json := []byte(`{"number" : 12, "parentType" : "parentType", "associatedTags" : ["coucou"], "updatedAt" : "2009-11-10T23:00:00Z", "annotation" : {"test" : "ok"}}`)
		err := UnmarshalJSON(bytes.NewReader(json), server)

		So(err, ShouldBeNil)
		So(server.Number, ShouldResemble, 12)
		So(server.ParentType, ShouldResemble, "parentType")
		So(server.UpdatedAt.String(), ShouldResemble, time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).String())
		So(server.AssociatedTags, ShouldResemble, []string{"coucou"})
	})
}

func TestUnmarshalJSONWithInvalidJSON(t *testing.T) {

	Convey("Given I call the method UnmarshalJSON with an invalid json", t, func() {

		server := &Server{}
		err := UnmarshalJSON(bytes.NewReader([]byte(`{"parentType" : 123`)), server)

		So(err.Error(), ShouldResemble, "error 0: error 422 (elemental): Validation Error: Invalid JSON\n")
	})
}

func TestUnmarshalJSONWithInvalidKeyAndValueJSON(t *testing.T) {

	Convey("Given I call the method UnmarshalJSON with a invalid json because array instead of string", t, func() {

		server := &Server{}
		json := []byte(`{"parentType" : ["ce"], "ParentTyp" : "parent"}`)
		err := UnmarshalJSON(bytes.NewReader(json), server)

		So(err.Error(), ShouldResemble, "error 0: error 422 (elemental): Validation Error: Attribute 'ParentTyp' is invalid\nerror 1: error 422 (elemental): Validation Error: Data '[ce]' of attribute 'parentType' should be a 'string'\n")
	})

	Convey("Given I call the method UnmarshalJSON with a invalid json because string instead of int", t, func() {

		server := &Server{}
		json := []byte(`{"number" : "12"}`)
		err := UnmarshalJSON(bytes.NewReader(json), server)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldResemble, "error 0: error 422 (elemental): Validation Error: Data '12' of attribute 'number' should be a 'int'\n")
	})

	Convey("Given I call the method UnmarshalJSON with a invalid json because string instead of int", t, func() {

		server := &Server{}
		json := []byte(`{"annotation" : "12"}`)
		err := UnmarshalJSON(bytes.NewReader(json), server)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldResemble, "error 0: error 422 (elemental): Validation Error: Data '12' of attribute 'annotation' should be a 'map[string]string'\n")
	})

	Convey("Given I call the method UnmarshalJSON with a invalid json because []int instead of []string", t, func() {

		server := &Server{}
		json := []byte(`{"associatedTags" : [12]}`)
		err := UnmarshalJSON(bytes.NewReader(json), server)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldResemble, "error 0: error 422 (elemental): Validation Error: Data '[12]' of attribute 'associatedTags' should be a '[]string'\n")
	})

	Convey("Given I call the method UnmarshalJSON with an invalid key", t, func() {

		server := &Server{}
		json := []byte(`{"associateags" : [12]}`)
		err := UnmarshalJSON(bytes.NewReader(json), server)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldResemble, "error 0: error 422 (elemental): Validation Error: Attribute 'associateags' is invalid\n")
	})

	Convey("Given I call the method UnmarshalJSON with an invalid contant type", t, func() {

		server := &Server{}
		json := []byte(`{"boom" : 12}`)
		err := UnmarshalJSON(bytes.NewReader(json), server)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldResemble, "error 0: error 422 (elemental): Validation Error: Data '12' of attribute 'boom' should be a 'string'\n")
	})

	Convey("Given I call the method UnmarshalJSON with an invalid key", t, func() {

		server := &Server{}
		json := []byte(`{"updatedAt" : 12}`)
		err := UnmarshalJSON(bytes.NewReader(json), server)

		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldResemble, "error 0: error 422 (elemental): Validation Error: Data '12' of attribute 'updatedAt' should be a 'string in format YYYY-MM-DDTHH:MM:SSZ'\n")
	})
}
