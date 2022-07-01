package errors

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Error is an error with a message
type Error struct {
	Status      int    `json:"status,omitempty"`
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%v: %v", e.Code, e.Description)
}

// New message
func New(message string) error {
	return Error{Description: message}
}

// Wrap message
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

var (
	// OK error message
	OK = &Error{Description: "OK", Status: 200}
	// Created error message
	Created = &Error{Description: "Created", Status: 201}
	// NoContent error message
	NoContent = &Error{Description: "No Content", Status: 204}
	// BadRequest error message
	BadRequest = &Error{Description: "Bad Request", Status: 400}
	// BadParameters error message
	BadParameters = &Error{Description: "Bad Parameters", Status: 400}
	// Database error message
	Database = &Error{Description: "Database Access error", Status: 500}
	// InternalServerError message
	InternalServerError = &Error{Description: "Internal Server error", Status: 500}
	// NotImplemented error message
	NotImplemented = &Error{Description: "Method not implemented", Status: 501}
	// TokenRequired error message
	TokenRequired = &Error{Description: "This endpoint requires a Bearer token", Status: 401}
	// Unauthrorized error message
	Unauthrorized = &Error{Description: "Invalid username, password or token", Status: 401}
	// TokenInvalid error message
	TokenInvalid = &Error{Description: "Invalid token", Status: 401}
	// UnprocessableEntity error message
	UnprocessableEntity = &Error{Description: "Unprocessable Entity", Status: 422}
	// NotFound 404 error message
	NotFound = &Error{Description: "Not Found", Status: 404}
)

// SendError helper function to write an http error.
func SendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	er, ok := err.(*Error)
	if ok {
		w.WriteHeader(er.Status)
		encoder := json.NewEncoder(w)
		encoder.Encode(er)
		return
	}
	w.WriteHeader(400)
	encoder := json.NewEncoder(w)
	encoder.Encode(err)
}
