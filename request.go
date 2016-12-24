package elemental

import (
	"encoding/json"
	"net/url"

	uuid "github.com/satori/go.uuid"
)

// A Request represents an abstract request on an elemental model.
type Request struct {
	RequestID      string          `json:"rid"`
	Namespace      string          `json:"namespace"`
	Operation      Operation       `json:"operation"`
	Identity       Identity        `json:"identity"`
	ObjectID       string          `json:"objectID"`
	ParentIdentity Identity        `json:"parentIdentity"`
	ParentID       string          `json:"parentID"`
	Data           json.RawMessage `json:"data,omitempty"`
	Parameters     url.Values      `json:"parameters,omitempty"`
	Username       string          `json:"username,omitempty"`
	Password       string          `json:"password,omitempty"`
}

// NewRequest returns a new Request.
func NewRequest(ns string, op Operation, identity Identity) *Request {

	return &Request{
		RequestID: uuid.NewV4().String(),
		Namespace: ns,
		Operation: op,
		Identity:  identity,
	}
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

	return json.Unmarshal(r.Data, &dst)
}

// A Response contains the response from a Request.
type Response struct {
	Request    *Request        `json:"request"`
	StatusCode int             `json:"statusCode"`
	Data       json.RawMessage `json:"data,omitempty"`
	Count      int             `json:"count"`
	Total      int             `json:"total"`
}

// NewResponse returns a new Response
func NewResponse() *Response {

	return &Response{}
}

// Encode encodes the given identifiable into the request.
func (r *Response) Encode(obj interface{}) error {

	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	r.Data = data

	return nil
}

// Decode decodes the data into the given destination
func (r *Response) Decode(dst interface{}) error {

	return json.Unmarshal(r.Data, &dst)
}
