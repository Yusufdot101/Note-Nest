package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (h *authHandler) loginUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	refreshToken, accessToken, err := h.svc.loginUser(v, input.Email, input.Password)
	if err != nil {
		switch {
		case errors.Is(err, validator.ErrFailedValidation):
			customerrors.FailedValidationErrorResponse(w, v.Errors)
		case errors.Is(err, customerrors.ErrNoRecord) || errors.Is(err, customerrors.ErrInvalidCredentials):
			customerrors.InvalidCredentialsErrorResponse(w)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	ttl, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRATION_TIME"))
	if err != nil {
		customerrors.ServerErrorResponse(w, errors.New("invalid refresh token expiration time"))
		return
	}
	err = utilities.SetTokenCookie(w, "REFRESH", refreshToken, "/auth", ttl)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"access_token": accessToken}, http.StatusCreated)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
	}
}
