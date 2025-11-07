package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (h *authHandler) SignupUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	refreshToken, accessToken, err := h.svc.registerUser(v, input.Name, input.Email, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, validator.ErrFailedValidation):
			custom_errors.FailedValidationErrorResponse(w, v.Errors)
		case errors.Is(err, user.ErrDuplicateEmail):
			v.AddError("email", "a user with this email already exists")
			custom_errors.FailedValidationErrorResponse(w, v.Errors)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	ttl, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRATION_TIME"))
	if err != nil {
		custom_errors.ServerErrorResponse(w, errors.New("invalid refresh token expiration time"))
		return
	}
	err = utilities.SetTokenCookie(w, "REFRESH", refreshToken, "/auth", ttl)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"token": accessToken}, http.StatusCreated)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
