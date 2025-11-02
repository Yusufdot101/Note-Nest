package user

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/jsonutil"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (h *userHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := jsonutil.ReadJSON(w, r, &input)
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	t, err := h.svc.loginUser(v, input.Email, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, validator.ErrFailedValidation):
			custom_errors.FailedValidationErrorResponse(w, v.Errors)
		case errors.Is(err, custom_errors.ErrNoRecord) || errors.Is(err, custom_errors.ErrInvalidCredentials):
			custom_errors.InvalidCredentialsErrorResponse(w)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = jsonutil.WriteJSON(w, jsonutil.Message{"token": t}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
