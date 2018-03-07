package elemental

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

// A Response contains the response from a Request.
type Response struct {
	StatusCode int                 `json:"statusCode"`
	Data       jsoniter.RawMessage `json:"data,omitempty"`
	Count      int                 `json:"count"`
	Total      int                 `json:"total"`
	Messages   []string            `json:"messages,omitempty"`
	Redirect   string              `json:"redirect,omitempty"`
	RequestID  string              `json:"requestID"`

	// TODO: this is kept for backward compat
	Request *Request `json:"request,omitempty"`
}

// NewResponse returns a new Response
func NewResponse(req *Request) *Response {

	return &Response{
		RequestID: req.RequestID,
		Request:   req,
	}
}

// Encode encodes the given identifiable into the request.
func (r *Response) Encode(obj interface{}) error {

	data, err := jsoniter.Marshal(obj)
	if err != nil {
		return err
	}

	r.Data = data

	return nil
}

// Decode decodes the data into the given destination
func (r *Response) Decode(dst interface{}) error {

	if r.StatusCode == http.StatusNoContent {
		return nil
	}

	return jsoniter.Unmarshal(r.Data, &dst)
}
