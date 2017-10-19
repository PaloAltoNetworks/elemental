package elemental

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"

	uuid "github.com/satori/go.uuid"
)

var snipSlice = []string{"[snip]"}

// A Request represents an abstract request on an elemental model.
type Request struct {
	RequestID            string                     `json:"rid"`
	Namespace            string                     `json:"namespace"`
	Recursive            bool                       `json:"recursive"`
	Operation            Operation                  `json:"operation"`
	Identity             Identity                   `json:"identity"`
	Order                []string                   `json:"order"`
	ObjectID             string                     `json:"objectID"`
	ParentIdentity       Identity                   `json:"parentIdentity"`
	ParentID             string                     `json:"parentID"`
	Data                 json.RawMessage            `json:"data,omitempty"`
	Parameters           url.Values                 `json:"parameters,omitempty"`
	Headers              http.Header                `json:"headers,omitempty"`
	Username             string                     `json:"username,omitempty"`
	Password             string                     `json:"password,omitempty"`
	Page                 int                        `json:"page,omitempty"`
	PageSize             int                        `json:"pageSize,omitempty"`
	OverrideProtection   bool                       `json:"overrideProtection,omitempty"`
	Version              int                        `json:"version,omitempty"`
	TrackingData         opentracing.TextMapCarrier `json:"trackingData,omitempty"`
	ExternalTrackingID   string                     `json:"externalTrackingID,omitempty"`
	ExternalTrackingType string                     `json:"externalTrackingType,omitempty"`

	Metadata map[string]interface{}

	ClientIP           string
	TLSConnectionState *tls.ConnectionState

	span        opentracing.Span
	wireContext opentracing.SpanContext
}

// NewRequest returns a new Request.
func NewRequest() *Request {

	return &Request{
		RequestID:    uuid.NewV4().String(),
		Parameters:   url.Values{},
		Headers:      http.Header{},
		Metadata:     map[string]interface{}{},
		TrackingData: opentracing.TextMapCarrier{},
	}
}

// NewRequestFromHTTPRequest returns a new Request from the given http.Request.
func NewRequestFromHTTPRequest(req *http.Request) (*Request, error) {

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
		identity = IdentityFromCategory(components[0])
	case 2:
		identity = IdentityFromCategory(components[0])
		ID = components[1]
	case 3:
		parentIdentity = IdentityFromCategory(components[0])
		parentID = components[1]
		identity = IdentityFromCategory(components[2])
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

	if v := req.URL.Query().Get("page"); v != "" {
		page, err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}

	if v := req.URL.Query().Get("pagesize"); v != "" {
		pageSize, err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}

	if v := req.URL.Query().Get("recursive"); v != "" {
		recursive = true
	}

	if v := req.URL.Query().Get("override"); v != "" {
		override = true
	}

	tracer := opentracing.GlobalTracer()
	var wireContext opentracing.SpanContext
	if tracer != nil {
		wireContext, _ = tracer.Extract(opentracing.TextMap, opentracing.HTTPHeadersCarrier(req.Header))
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
		Parameters:           req.URL.Query(),
		Username:             username,
		Password:             password,
		Data:                 data,
		TLSConnectionState:   req.TLS,
		Headers:              req.Header,
		OverrideProtection:   override,
		Metadata:             map[string]interface{}{},
		Version:              version,
		TrackingData:         opentracing.TextMapCarrier{},
		ExternalTrackingID:   req.Header.Get("X-External-Tracking-ID"),
		ExternalTrackingType: req.Header.Get("X-External-Tracking-Type"),
		wireContext:          wireContext,
		Order:                req.URL.Query()["order"],
		ClientIP:             req.RemoteAddr,
	}, nil
}

// StartTracing starts tracing the request.
func (r *Request) StartTracing() {

	tracer := opentracing.GlobalTracer()
	if tracer == nil {
		return
	}

	if r.wireContext == nil {
		r.wireContext, _ = tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(r.TrackingData))
	}

	r.span = opentracing.StartSpan(r.tracingName(), ext.RPCServerOption(r.wireContext))

	// Remove sensitive information from parameters.
	safeParameters := url.Values{}
	for k, v := range r.Parameters {
		lk := strings.ToLower(k)
		if lk == "token" || lk == "password" {
			safeParameters[k] = snipSlice
			continue
		}
		safeParameters[k] = v
	}

	// Remove sensitive information from headers.
	safeHeaders := http.Header{}
	for k, v := range r.Headers {
		lk := strings.ToLower(k)
		if lk == "authorization" {
			safeHeaders[k] = snipSlice
			continue
		}
		safeHeaders[k] = v
	}

	r.span.SetTag("elemental.request.api_version", r.Version)
	r.span.SetTag("elemental.request.external_tracking_id", r.ExternalTrackingID)
	r.span.SetTag("elemental.request.external_tracking_type", r.ExternalTrackingType)
	r.span.SetTag("elemental.request.headers", safeHeaders)
	r.span.SetTag("elemental.request.claims", r.extractClaims())
	r.span.SetTag("elemental.request.id", r.RequestID)
	r.span.SetTag("elemental.request.identity", r.Identity.Name)
	r.span.SetTag("elemental.request.namespace", r.Namespace)
	r.span.SetTag("elemental.request.object.id", r.ObjectID)
	r.span.SetTag("elemental.request.operation", r.Operation)
	r.span.SetTag("elemental.request.order_by", r.Order)
	r.span.SetTag("elemental.request.override_protection", r.OverrideProtection)
	r.span.SetTag("elemental.request.page.number", r.Page)
	r.span.SetTag("elemental.request.page.size", r.PageSize)
	r.span.SetTag("elemental.request.parameters", safeParameters)
	r.span.SetTag("elemental.request.parent.id", r.ParentID)
	r.span.SetTag("elemental.request.parent.identity", r.ParentIdentity.Name)
	r.span.SetTag("elemental.request.recursive", r.Recursive)
	r.span.SetTag("elemental.request.client_ip", r.ClientIP)

	if r.Data != nil {
		r.span.LogFields(log.Object("payload", string(r.Data)))
	}
}

// FinishTracing will finish the request tracing.
func (r *Request) FinishTracing() {

	if r.span == nil {
		return
	}

	r.span.Finish()
}

// Span returns the current span.
func (r *Request) Span() opentracing.Span {

	return r.span
}

// NewChildSpan return a new child tracing span.
func (r *Request) NewChildSpan(name string) opentracing.Span {

	if r.span == nil {
		return opentracing.StartSpan(name)
	}

	return opentracing.StartSpan(name, opentracing.ChildOf(r.span.Context()))
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
	req.span = r.span
	req.wireContext = r.wireContext
	req.ExternalTrackingID = r.ExternalTrackingID
	req.ExternalTrackingType = r.ExternalTrackingType
	req.ClientIP = r.ClientIP
	req.Order = append([]string{}, r.Order...)

	for k, v := range r.Headers {
		req.Headers[k] = v
	}

	for k, v := range r.Parameters {
		req.Parameters[k] = v
	}

	for k, v := range r.Metadata {
		req.Metadata[k] = v
	}

	for k, v := range r.TrackingData {
		req.TrackingData[k] = v
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

func (r *Request) tracingName() string {

	switch r.Operation {

	case OperationCreate:
		return fmt.Sprintf("elemental.request.create.%s", r.Identity.Category)

	case OperationRetrieveMany:
		return fmt.Sprintf("elemental.request.retrieve_many.%s", r.Identity.Category)

	case OperationInfo:
		return fmt.Sprintf("elemental.request.info.%s", r.Identity.Category)

	case OperationUpdate:
		return fmt.Sprintf("elemental.request.update.%s", r.Identity.Category)

	case OperationDelete:
		return fmt.Sprintf("elemental.request.delete.%s", r.Identity.Category)

	case OperationRetrieve:
		return fmt.Sprintf("elemental.request.retrieve.%s", r.Identity.Category)

	case OperationPatch:
		return fmt.Sprintf("elemental.request.patch.%s", r.Identity.Category)
	}

	return fmt.Sprintf("Unknown operation: %s", r.Operation)
}

func (r *Request) extractClaims() string {

	tokenParts := strings.SplitN(r.Password, ".", 3)
	if len(tokenParts) != 3 {
		return "{}"
	}

	identity, err := base64.RawStdEncoding.DecodeString(tokenParts[1])
	if err != nil {
		return "{}"
	}

	return string(identity)
}
