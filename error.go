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
func NewError(title, description, subject string, code int) *Error {

	return &Error{
		Code:        code,
		Description: description,
		Subject:     subject,
		Title:       title,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error %d (%s): %s: %s", e.Code, e.Subject, e.Title, e.Description)
}

// Errors represents a list of Error.
type Errors []*Error

// NewErrors creates a new Errors.
func NewErrors(errors ...*Error) Errors {

	return append(Errors{}, errors...)
}

func (e *Errors) Error() string {

	var errorString string

	for i, err := range *e {
		errorString += fmt.Sprintf("error %d: %s\n", i, err.Error())
	}

	return errorString
}

// Code returns the code of the first error code in the Errors.
func (e *Errors) Code() int {

	return (*e)[0].Code
}
