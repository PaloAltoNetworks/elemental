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

import "net/http"

// A Response contains the response from a Request.
type Response struct {
	StatusCode int
	Data       []byte
	Count      int
	Total      int
	Next       string
	Messages   []string
	Redirect   string
	RequestID  string
	Request    *Request
	Cookies    []*http.Cookie
}

// NewResponse returns a new Response
func NewResponse(req *Request) *Response {

	return &Response{
		RequestID: req.RequestID,
		Request:   req,
	}
}

// GetEncoding returns the encoding used to encode the entity.
func (r *Response) GetEncoding() EncodingType {
	return r.Request.Accept
}

// Encode encodes the given oject into the response.
func (r *Response) Encode(obj interface{}) (err error) {

	r.Data, err = Encode(r.GetEncoding(), obj)
	return err
}
