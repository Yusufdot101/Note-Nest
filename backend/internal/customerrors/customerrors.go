/*
Package customerrors provides different custom error types and error responses for consistent responses
*/
package customerrors

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/utilities"
)

var (
	// ErrNoRecord is a custom error type used when a resource is not found
	ErrNoRecord = errors.New("record not found")
	// ErrInvalidCredentials is a custom error type used when invalid credentials are provided
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrUpdateTimeout is a custom error type used when note update is attempted past the update timeout
	ErrUpdateTimeout = errors.New("cannot update")
)

// ServerErrorResponse is custom response for server related errors
func ServerErrorResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	msg := map[string]string{"message": "the server encountered an error and could not proceed with your request"}
	errorResponse(w, msg, http.StatusInternalServerError)
}

func errorResponse(w http.ResponseWriter, errMsg any, statusCode int) {
	err := utilities.WriteJSON(w, utilities.Message{"error": errMsg}, statusCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// BadRequestErrorResponse is a response used when a client provides bad/invalid inputs
func BadRequestErrorResponse(w http.ResponseWriter, err error) {
	errorResponse(w, err.Error(), http.StatusBadRequest)
}

// NotFoundErrorResponse is a response used when a resource is not found
func NotFoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": "resource could not be found"}
	errorResponse(w, msg, http.StatusNotFound)
}

// MethodNotAllowedErrorResponse is a response used when a method is not allowed for a certain endpoint
func MethodNotAllowedErrorResponse(w http.ResponseWriter, r *http.Request) {
	msg := map[string]string{"message": fmt.Sprintf("%s method not allowed for this resource", r.Method)}
	errorResponse(w, msg, http.StatusMethodNotAllowed)
}

// FailedValidationErrorResponse is a response used when validation errors occur
func FailedValidationErrorResponse(w http.ResponseWriter, errors map[string]string) {
	errorResponse(w, errors, http.StatusBadRequest)
}

// InvalidCredentialsErrorResponse is a response used when invalid credentials are provided
func InvalidCredentialsErrorResponse(w http.ResponseWriter) {
	msg := map[string]string{"message": "invalid credentials"}
	errorResponse(w, msg, http.StatusBadRequest)
}

// RequireAuthenticationErrorResponse is a response used when an endpoint requires authenticated user
func RequireAuthenticationErrorResponse(w http.ResponseWriter) {
	msg := map[string]string{"message": "you must be logged in to access this resource"}
	errorResponse(w, msg, http.StatusUnauthorized)
}

// InvalidAuthenticationTokenErrorResponse is a response used when an invalid authorization token is provided
func InvalidAuthenticationTokenErrorResponse(w http.ResponseWriter) {
	msg := map[string]string{"message": "invalid or expired token"}
	errorResponse(w, msg, http.StatusUnauthorized)
}

// UpdateTimeoutErrorResponse is a response used when an note update is attempted past its update timeout
func UpdateTimeoutErrorResponse(w http.ResponseWriter) {
	msg := map[string]string{"message": "you can no longer update the content of this note"}
	errorResponse(w, msg, http.StatusForbidden)
}

func RateLimitExceededErrorResponse(w http.ResponseWriter) {
	msg := map[string]string{"message": "too many requests, please try again later"}
	errorResponse(w, msg, http.StatusTooManyRequests)
}
