// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// An Error represents a computational error.
//
// They can be encoded and sent back to the clients.
type Error struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Subject     string      `json:"subject"`
	Title       string      `json:"title"`
	Data        interface{} `json:"data"`
	Trace       string      `json:"trace"`
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
type Errors []error

// NewErrors creates a new Errors.
func NewErrors(errors ...error) Errors {

	if len(errors) == 0 {
		return Errors{}
	}

	return append(Errors{}, errors...)
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

	switch e0 := e[0].(type) {
	case Error:
		return e0.Code
	default:
		return -1
	}
}

// At returns the Error at the given index.
// If the error at the given index is not an Error or doesn't exists
// it returns an Unknown Error.
func (e Errors) At(i int) Error {

	switch ei := e[i].(type) {
	case Error:
		return ei
	default:
		return NewError("Standard error", ei.Error(), "elemental", -1)
	}
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
		elementalError = e[0].(Error)

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
