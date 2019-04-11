package elemental

import (
	"encoding/json"
)

// A Response contains the response from a Request.
type Response struct {
	StatusCode int             `json:"statusCode"`
	Data       json.RawMessage `json:"data,omitempty"`
	Count      int             `json:"count"`
	Total      int             `json:"total"`
	Messages   []string        `json:"messages,omitempty"`
	Redirect   string          `json:"redirect,omitempty"`
	RequestID  string          `json:"requestID"`

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
