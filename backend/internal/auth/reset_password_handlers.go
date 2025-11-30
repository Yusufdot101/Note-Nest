package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (h authHandler) forgotPassword(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	err = h.svc.forgotPassword(v, input.Email)
	if err != nil {
		if errors.Is(err, validator.ErrFailedValidation) {
			customerrors.FailedValidationErrorResponse(w, v.Errors)
			return
		}
		log.Printf("forgotPassword error: %v\n", err)
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "If an account exists with that email, a reset link has been sent."}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
	}
}

func (h authHandler) resetPassword(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Password string `json:"password"`
		Token    string `json:"token"`
	}

	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	err = h.svc.resetPassword(v, input.Token, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, validator.ErrFailedValidation):
			customerrors.FailedValidationErrorResponse(w, v.Errors)
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.InvalidTokenErrorResponse(w)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "password reset successfully"}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
	}
}
