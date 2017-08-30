// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import (
	"fmt"
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
	RequestID   string      `json:"requestID"`
}

// NewError returns a new Error.
func NewError(title, description, subject string, code int) Error {

	return Error{
		Code:        code,
		Description: description,
		Subject:     subject,
		Title:       title,
	}
}

func (e Error) Error() string {

	if e.RequestID != "" {
		return fmt.Sprintf("error %d (%s): %s: %s [requestID: %s]", e.Code, e.Subject, e.Title, e.Description, e.RequestID)
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

	var strs []string

	for _, err := range e {
		strs = append(strs, err.Error())
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
