// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import "fmt"

// An Error represents a computational error.
//
// They can be encoded and sent back to the clients.
type Error struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Subject     string      `json:"subject"`
	Title       string      `json:"title"`
	Data        interface{} `json:"data"`
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
	return fmt.Sprintf("error %d (%s): %s: %s", e.Code, e.Subject, e.Title, e.Description)
}

// Errors represents a list of Error.
type Errors []error

// NewErrors creates a new Errors.
func NewErrors(errors ...error) Errors {

	return append(Errors{}, errors...)
}

func (e Errors) Error() string {

	var str string

	for i, err := range e {
		str += fmt.Sprintf("error %d: %s\n", i, err.Error())
	}

	return str
}

// Code returns the code of the first error code in the Errors.
func (e Errors) Code() int {

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
		return NewError("Unknown Error", "Error doesn't exists", "elemental", -1)
	}
}
