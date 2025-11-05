package token

import (
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func (h *TokenHandler) NewAccessToken(w http.ResponseWriter, r *http.Request) {
	// pull userID from context, which is currently float64 because of MapClaims, then convert it to int.
	userID := int(r.Context().Value(middleware.UserIDKey).(float64))
	accessToken, err := CreateJWT("ACCESS", userID)
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}
	err = utilities.WriteJSON(w, utilities.Message{"token": accessToken}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
