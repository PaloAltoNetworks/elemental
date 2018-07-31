package elemental

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// A Request represents an abstract request on an elemental model.
type Request struct {
	RequestID            string               `json:"rid"`
	Namespace            string               `json:"namespace"`
	Recursive            bool                 `json:"recursive"`
	Operation            Operation            `json:"operation"`
	Identity             Identity             `json:"identity"`
	Order                []string             `json:"order"`
	ObjectID             string               `json:"objectID"`
	ParentIdentity       Identity             `json:"parentIdentity"`
	ParentID             string               `json:"parentID"`
	Data                 json.RawMessage      `json:"data,omitempty"`
	Parameters           map[string]Parameter `json:"parameters,omitempty"`
	Headers              http.Header          `json:"headers,omitempty"`
	Username             string               `json:"username,omitempty"`
	Password             string               `json:"password,omitempty"`
	Page                 int                  `json:"page,omitempty"`
	PageSize             int                  `json:"pageSize,omitempty"`
	OverrideProtection   bool                 `json:"overrideProtection,omitempty"`
	Version              int                  `json:"version,omitempty"`
	ExternalTrackingID   string               `json:"externalTrackingID,omitempty"`
	ExternalTrackingType string               `json:"externalTrackingType,omitempty"`

	Metadata           map[string]interface{} `json:"-"`
	ClientIP           string                 `json:"-"`
	TLSConnectionState *tls.ConnectionState   `json:"-"`

	req *http.Request
}

// NewRequest returns a new Request.
func NewRequest() *Request {

	return &Request{
		RequestID:  uuid.NewV4().String(),
		Parameters: map[string]Parameter{},
		Headers:    http.Header{},
		Metadata:   map[string]interface{}{},
	}
}

// NewRequestFromHTTPRequest returns a new Request from the given http.Request.
func NewRequestFromHTTPRequest(req *http.Request, manager ModelManager) (*Request, error) {

	if req.URL == nil || req.URL.String() == "" {
		return nil, fmt.Errorf("request must have an url")
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
			return nil, fmt.Errorf("Invalid api version number '%s'", components[1])
		}
		// once we've set the version, we remove it, and continue as usual.
		components = append(components[:0], components[2:]...)
	}

	switch len(components) {
	case 1:
		parentIdentity = RootIdentity
		identity = manager.IdentityFromCategory(components[0])
	case 2:
		parentIdentity = RootIdentity
		identity = manager.IdentityFromCategory(components[0])
		ID = components[1]
	case 3:
		parentIdentity = manager.IdentityFromCategory(components[0])
		parentID = components[1]
		identity = manager.IdentityFromCategory(components[2])
	default:
		return nil, fmt.Errorf("%s is not a valid elemental request path", req.URL)
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

	q := req.URL.Query()
	if v := q.Get("page"); v != "" {
		page, err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		q.Del("page")
	}

	if v := q.Get("pagesize"); v != "" {
		pageSize, err = strconv.Atoi(v)
		if err != nil {
			return nil, err
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

	paramsMap := map[string]Parameter{}
	for _, pdef := range ParametersForOperation(manager.Relationships(), identity, parentIdentity, operation) {
		p, err := pdef.Parse(q[pdef.Name])
		if err != nil {
			return nil, err
		}
		paramsMap[pdef.Name] = *p
	}

	return &Request{
		RequestID:            uuid.NewV4().String(),
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
		Order:                req.URL.Query()["order"],
		ClientIP:             req.RemoteAddr,
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

// Encode encodes the given identifiable into the request.
func (r *Request) Encode(entity Identifiable) error {

	data, err := json.Marshal(entity)
	if err != nil {
		return err
	}

	r.Data = data

	return nil
}

// Decode decodes the data into the given destination
func (r *Request) Decode(dst interface{}) error {

	return UnmarshalJSON(r.Data, &dst)
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
