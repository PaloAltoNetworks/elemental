// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import "fmt"

// Error represents a computational error.
type Error struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Subject     string      `json:"subject"`
	Title       string      `json:"title"`
	Data        interface{} `json:"data"`
}

// NewError creates a new *Error.
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

// Errors represents a list of errors
type Errors []*Error

// NewErrors creates a new Errors.
func NewErrors(errors ...*Error) Errors {

	return append(Errors{}, errors...)
}
