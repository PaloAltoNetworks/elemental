package elemental

import (
	"encoding/json"
	"net/http"
)

// A Response contains the response from a Request.
type Response struct {
	StatusCode int             `json:"statusCode"`
	Data       json.RawMessage `json:"data,omitempty"`
	Count      int             `json:"count"`
	Total      int             `json:"total"`
	Messages   []string        `json:"messages,omitempty"`
	Redirect   string          `json:"redirect,omitempty"`
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

	if r.StatusCode == http.StatusNoContent {
		return nil
	}

	return json.Unmarshal(r.Data, &dst)
}
