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
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// IsErrorWithCode returns true if the given error is an elemental.Error
// or elemental.Errors with the status set to the given code.
func IsErrorWithCode(err error, code int) bool {

	var c int
	switch e := err.(type) {
	case Error:
		c = e.Code
	case Errors:
		c = e.Code()
	}

	return c == code
}

// An Error represents a computational error.
//
// They can be encoded and sent back to the clients.
type Error struct {
	Code        int         `msgpack:"code" json:"code,omitempty"`
	Description string      `msgpack:"description" json:"description"`
	Subject     string      `msgpack:"subject" json:"subject"`
	Title       string      `msgpack:"title" json:"title"`
	Data        interface{} `msgpack:"data" json:"data,omitempty"`
	Trace       string      `msgpack:"trace" json:"trace,omitempty"`
}

// NewError returns a new Error.
func NewError(title, description, subject string, code int) Error {

	return NewErrorWithData(title, description, subject, code, nil)
}

// NewErrorWithData returns a new Error with the given opaque data.
func NewErrorWithData(title, description, subject string, code int, data interface{}) Error {

	return Error{
		Code:        code,
		Description: description,
		Subject:     subject,
		Title:       title,
		Data:        data,
	}
}

func (e Error) Error() string {

	if e.Trace != "" {
		return fmt.Sprintf("error %d (%s): %s: %s [trace: %s]", e.Code, e.Subject, e.Title, e.Description, e.Trace)
	}

	return fmt.Sprintf("error %d (%s): %s: %s", e.Code, e.Subject, e.Title, e.Description)
}

// Errors represents a list of Error.
type Errors []Error

// NewErrors creates a new Errors.
func NewErrors(errors ...error) Errors {

	out := Errors{}
	if len(errors) == 0 {
		return out
	}

	return out.Append(errors...)
}

func (e Errors) Error() string {

	strs := make([]string, len(e))

	for i := range e {
		strs[i] = e[i].Error()
	}

	return strings.Join(strs, ", ")
}

// Code returns the code of the first error code in the Errors.
func (e Errors) Code() int {

	if len(e) == 0 {
		return -1
	}

	return e[0].Code
}

// Append returns returns a copy of the receiver containing
// also the given errors.
func (e Errors) Append(errs ...error) Errors {

	out := append(Errors{}, e...)

	for _, err := range errs {
		switch er := err.(type) {
		case Error:
			out = append(out, er)
		case Errors:
			out = append(out, er...)
		default:
			out = append(out, NewError("Internal Server Error", err.Error(), "elemental", http.StatusInternalServerError))
		}
	}

	return out
}

// Trace returns Errors with all inside Error marked with the
// given trace ID.
func (e Errors) Trace(id string) Errors {

	out := Errors{}

	for _, err := range e {
		err.Trace = id
		out = append(out, err)
	}

	return out
}

// DecodeErrors decodes the given bytes into a en elemental.Errors.
func DecodeErrors(data []byte) (Errors, error) {

	es := []Error{}
	if err := json.Unmarshal(data, &es); err != nil {
		return nil, err
	}

	e := NewErrors()
	for _, err := range es {
		e = append(e, err)
	}

	return e, nil
}

// IsValidationError returns true if the given error is a validation error
// with the given title for the given attribute.
func IsValidationError(err error, title string, attribute string) bool {

	var elementalError Error
	switch e := err.(type) {

	case Errors:
		if e.Code() != http.StatusUnprocessableEntity {
			return false
		}
		if len(e) != 1 {
			return false
		}
		elementalError = e[0]

	case Error:
		if e.Code != http.StatusUnprocessableEntity {
			return false
		}
		elementalError = e

	default:
		return false
	}

	if elementalError.Title != title {
		return false
	}

	if elementalError.Data == nil {
		return false
	}

	m, ok := elementalError.Data.(map[string]interface{})
	if !ok {
		return false
	}

	return m["attribute"].(string) == attribute
}
