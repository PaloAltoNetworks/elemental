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
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

// A Request represents an abstract request on an elemental model.
type Request struct {
	RequestID            string
	Namespace            string
	Recursive            bool
	Operation            Operation
	Identity             Identity
	Order                []string
	ObjectID             string
	ParentIdentity       Identity
	ParentID             string
	Data                 []byte
	Parameters           Parameters
	Headers              http.Header
	Username             string
	Password             string
	Page                 int
	PageSize             int
	OverrideProtection   bool
	Version              int
	ExternalTrackingID   string
	ExternalTrackingType string
	ContentType          EncodingType
	Accept               EncodingType

	Metadata           map[string]interface{}
	ClientIP           string
	ClientIPSource     string
	TLSConnectionState *tls.ConnectionState

	req *http.Request
}

// NewRequest returns a new Request.
func NewRequest() *Request {

	return &Request{
		RequestID:   uuid.Must(uuid.NewV4()).String(),
		Parameters:  Parameters{},
		Headers:     http.Header{},
		Metadata:    map[string]interface{}{},
		ContentType: EncodingTypeJSON,
		Accept:      EncodingTypeJSON,
	}
}

// NewRequestFromHTTPRequest returns a new Request from the given http.Request.
func NewRequestFromHTTPRequest(req *http.Request, manager ModelManager) (*Request, error) {

	if req.URL == nil || req.URL.String() == "" {
		return nil, NewError("Bad Request", "Request must have an url", "elemental", http.StatusBadRequest)
	}

	var operation Operation
	var identity Identity
	var parentIdentity Identity
	var ID string
	var parentID string
	var username string
	var password string
	var version int
	var data []byte
	var err error

	auth := strings.Split(req.Header.Get("Authorization"), " ")
	if len(auth) == 2 {
		username = auth[0]
		password = auth[1]
	}

	components := strings.Split(req.URL.Path, "/")

	// We remove the first element as it's always empty
	components = append(components[:0], components[1:]...)

	// If the first one is "v" it means the next one has to be a int for the version number.
	if components[0] == "v" {
		version, err = strconv.Atoi(components[1])
		if err != nil {
			return nil, NewError("Bad Request", fmt.Sprintf("Invalid api version number '%s'", components[1]), "elemental", http.StatusBadRequest)
		}
		// once we've set the version, we remove it, and continue as usual.
		components = append(components[:0], components[2:]...)
	}

	switch len(components) {
	case 1:
		identity = manager.IdentityFromCategory(components[0])
	case 2:
		identity = manager.IdentityFromCategory(components[0])
		ID = components[1]
	case 3:
		parentIdentity = manager.IdentityFromCategory(components[0])
		parentID = components[1]
		identity = manager.IdentityFromCategory(components[2])
	default:
		return nil, NewError("Bad Request", fmt.Sprintf("%s is not a valid elemental request path", req.URL), "elemental", http.StatusBadRequest)
	}

	if parentIdentity.IsEmpty() {
		parentIdentity = RootIdentity
	}

	switch req.Method {
	case http.MethodDelete:
		operation = OperationDelete

	case http.MethodGet:
		if len(components) == 1 || len(components) == 3 {
			operation = OperationRetrieveMany
		} else {
			operation = OperationRetrieve
		}

	case http.MethodHead:
		operation = OperationInfo

	case http.MethodPatch:
		operation = OperationPatch
		data, err = ioutil.ReadAll(req.Body)
		defer req.Body.Close() // nolint: errcheck
		if err != nil {
			return nil, err
		}

	case http.MethodPost:
		operation = OperationCreate
		data, err = ioutil.ReadAll(req.Body)
		defer req.Body.Close() // nolint: errcheck
		if err != nil {
			return nil, err
		}

	case http.MethodPut:
		operation = OperationUpdate
		data, err = ioutil.ReadAll(req.Body)
		defer req.Body.Close() // nolint: errcheck
		if err != nil {
			return nil, err
		}
	}

	var page, pageSize int
	var recursive, override bool
	var order []string

	q := req.URL.Query()
	if v := q.Get("page"); v != "" {
		page, err = strconv.Atoi(v)
		if err != nil {
			return nil, NewError("Bad Request", "Parameter `page` must be an integer", "elemental", http.StatusBadRequest)
		}
		q.Del("page")
	}

	if v := q.Get("pagesize"); v != "" {
		pageSize, err = strconv.Atoi(v)
		if err != nil {
			return nil, NewError("Bad Request", "Parameter `pagesize` must be an integer", "elemental", http.StatusBadRequest)
		}
		q.Del("pagesize")
	}

	if v := q.Get("recursive"); v != "" {
		recursive = true
		q.Del("recursive")
	}

	if v := q.Get("override"); v != "" {
		override = true
		q.Del("override")
	}

	if v, ok := q["order"]; ok {
		for _, o := range v {
			if o == "" || o == "\u0000" {
				return nil, NewError("Bad Request", "Parameter `order` must be set when provided", "elemental", http.StatusBadRequest)
			}
		}
		order = v
		q.Del("order")
	}

	paramsMap := Parameters{}
	qKeys := map[string]struct{}{}
	for k := range q {
		qKeys[k] = struct{}{}
	}

	for _, pdef := range ParametersForOperation(manager.Relationships(), identity, parentIdentity, operation) {
		p, err := pdef.Parse(q[pdef.Name])
		if err != nil {
			return nil, err
		}
		delete(qKeys, pdef.Name)
		paramsMap[pdef.Name] = *p
	}

	if len(qKeys) > 0 {
		errs := NewErrors()
		for k := range qKeys {
			errs = errs.Append(NewError("Bad Request", fmt.Sprintf("Unknown parameter: `%s`", k), "elemental", http.StatusBadRequest))
		}
		return nil, errs
	}

	rel := RelationshipInfoForOperation(manager.Relationships(), identity, parentIdentity, operation)
	if rel != nil {
		if err := paramsMap.Validate(rel.RequiredParameters); err != nil {
			return nil, err
		}
	}

	var clientIP string
	var clientIPSource string
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		clientIP = ip
		clientIPSource = "X-Forwarded-For"
	} else if ip := req.Header.Get("X-Real-IP"); ip != "" {
		clientIP = ip
		clientIPSource = "X-Real-IP"
	} else {
		clientIP = req.RemoteAddr
		clientIPSource = "RemoteAddr"
	}

	contentType, acceptType, err := EncodingFromHeaders(req.Header)
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestID:            uuid.Must(uuid.NewV4()).String(),
		Namespace:            req.Header.Get("X-Namespace"),
		Recursive:            recursive,
		Page:                 page,
		PageSize:             pageSize,
		Operation:            operation,
		Identity:             identity,
		ObjectID:             ID,
		ParentID:             parentID,
		ParentIdentity:       parentIdentity,
		Parameters:           paramsMap,
		Username:             username,
		Password:             password,
		Data:                 data,
		TLSConnectionState:   req.TLS,
		Headers:              req.Header,
		OverrideProtection:   override,
		Metadata:             map[string]interface{}{},
		Version:              version,
		ExternalTrackingID:   req.Header.Get("X-External-Tracking-ID"),
		ExternalTrackingType: req.Header.Get("X-External-Tracking-Type"),
		Order:                order,
		ClientIP:             clientIP,
		ClientIPSource:       clientIPSource,
		ContentType:          contentType,
		Accept:               acceptType,
		req:                  req,
	}, nil
}

// Duplicate duplicates the Request.
func (r *Request) Duplicate() *Request {

	req := NewRequest()

	req.Namespace = r.Namespace
	req.Recursive = r.Recursive
	req.Page = r.Page
	req.PageSize = r.PageSize
	req.Operation = r.Operation
	req.Identity = r.Identity
	req.ObjectID = r.ObjectID
	req.ParentID = r.ParentID
	req.ParentIdentity = r.ParentIdentity
	req.Username = r.Username
	req.Password = r.Password
	req.Data = r.Data
	req.Version = r.Version
	req.OverrideProtection = r.OverrideProtection
	req.TLSConnectionState = r.TLSConnectionState
	req.ExternalTrackingID = r.ExternalTrackingID
	req.ExternalTrackingType = r.ExternalTrackingType
	req.ClientIP = r.ClientIP
	req.Order = append([]string{}, r.Order...)
	req.req = r.req
	req.ContentType = r.ContentType
	req.Accept = r.Accept

	for k, v := range r.Headers {
		req.Headers[k] = v
	}

	for k, v := range r.Parameters {
		req.Parameters[k] = v
	}

	for k, v := range r.Metadata {
		req.Metadata[k] = v
	}

	return req
}

// GetEncoding returns the encoding used to encode the body.
func (r *Request) GetEncoding() EncodingType {
	return r.ContentType
}

// Decode decodes the data into the given destination.
func (r *Request) Decode(dst interface{}) error {

	return Decode(r.GetEncoding(), r.Data, dst)
}

// HTTPRequest returns the native http.Request, if any.
func (r *Request) HTTPRequest() *http.Request {
	return r.req
}

func (r *Request) String() string {

	return fmt.Sprintf("<request id:%s operation:%s namespace:%s recursive:%v identity:%s objectid:%s parentidentity:%s parentid:%s version:%d>",
		r.RequestID,
		r.Operation,
		r.Namespace,
		r.Recursive,
		r.Identity,
		r.ObjectID,
		r.ParentIdentity,
		r.ParentID,
		r.Version,
	)
}
