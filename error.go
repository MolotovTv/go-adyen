// Copyright 2018 Matthieu CORNUT-CHAUVINC. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package adyen

import "encoding/json"

// ErrorType is the list of allowed values for the error's msg.
type ErrorType string

// Error Types
const (
	ErrorTypeAPI            ErrorType = "Default"                   // default
	ErrorTypeInvalidRequest ErrorType = "400 Bad Request"           // 400
	ErrorTypeAuthentication ErrorType = "401 Unauthorized"          // 401
	ErrorTypePermission     ErrorType = "403 Forbidden"             // 403
	ErrorTypeNotFound       ErrorType = "404 Not Found"             // 404
	ErrorTypeUnprocessable  ErrorType = "422 Unprocessable Entity"  // 422
	ErrorTypeAPIConnection  ErrorType = "500 Internal Server Error" // 500
)

// Error is the response returned when a call is unsuccessful.
type Error struct {

	// Err contains an internal error with an additional level of granularity
	// that can be used in some cases to get more detailed information about
	// what went wrong. For example, Err may hold a ChargeError that indicates
	// exactly what went wrong during a charge.
	Err            error     `json:"-"`
	HTTPStatusCode int       `json:"status"`
	Type           ErrorType `json:"type"`
	Message        string    `json:"message"`
}

// Error serializes the error object to JSON and returns it as a string.
func (e *Error) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

// APIConnectionError is a failure to connect to the Adyen API.
type APIConnectionError struct {
	adyenErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIConnectionError) Error() string {
	return e.adyenErr.Error()
}

// APIError is a catch all for any errors not covered by other types (and
// should be extremely uncommon).
type APIError struct {
	adyenErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *APIError) Error() string {
	return e.adyenErr.Error()
}

// AuthenticationError is a failure to properly authenticate during a request.
type AuthenticationError struct {
	adyenErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *AuthenticationError) Error() string {
	return e.adyenErr.Error()
}

// PermissionError results when you attempt to make an API request
// for which your API key doesn't have the right permissions.
type PermissionError struct {
	adyenErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *PermissionError) Error() string {
	return e.adyenErr.Error()
}

// UnprocessableError results when you attempt to make an API request
// for which your API key doesn't have the right permissions.
type UnprocessableError struct {
	adyenErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *UnprocessableError) Error() string {
	return e.adyenErr.Error()
}

// NotFoundError results when you attempt to make an API request
// for which your API key doesn't have the right permissions.
type NotFoundError struct {
	adyenErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *NotFoundError) Error() string {
	return e.adyenErr.Error()
}

// InvalidRequestError is an error that occurs when a request contains invalid
// parameters.
type InvalidRequestError struct {
	adyenErr *Error
}

// Error serializes the error object to JSON and returns it as a string.
func (e *InvalidRequestError) Error() string {
	return e.adyenErr.Error()
}
