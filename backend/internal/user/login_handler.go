package user

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (h *userHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	refreshToken, accessToken, err := h.svc.loginUser(v, input.Email, input.Password)
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

	err = utilities.SetJWTCookie(w, "REFRESH", refreshToken)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"token": accessToken}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
