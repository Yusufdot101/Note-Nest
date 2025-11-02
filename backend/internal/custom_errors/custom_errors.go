package custom_errors

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/jsonutil"
)

var (
	ErrNoRecord           = errors.New("record not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func ServerErrorResponse(w http.ResponseWriter, err error) {
	log.Println(err)
	msg := "the server encountered an error and could not proceed with your request"
	errorResponse(w, msg, http.StatusInternalServerError)
}

func errorResponse(w http.ResponseWriter, errMsg any, statusCode int) {
	err := jsonutil.WriteJSON(w, jsonutil.Message{"error": errMsg}, statusCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func BadRequestErrorResponse(w http.ResponseWriter, err error) {
	errorResponse(w, err.Error(), http.StatusBadRequest)
}

func NotFoundErrorResponse(w http.ResponseWriter, r *http.Request) {
	msg := "the resource you requested for could not be found"
	errorResponse(w, msg, http.StatusNotFound)
}

func MethodNotAllowedErrorResponse(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("the %s method is not allowed for this resource", r.Method)
	errorResponse(w, msg, http.StatusMethodNotAllowed)
}

func FailedValidationErrorResponse(w http.ResponseWriter, errors map[string]string) {
	errorResponse(w, errors, http.StatusBadRequest)
}

func InvalidCredentialsErrorResponse(w http.ResponseWriter) {
	msg := "invalid credentials"
	errorResponse(w, msg, http.StatusBadRequest)
}
