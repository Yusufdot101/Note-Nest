package auth

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func (h *authHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	tokenString, ok := r.Context().Value(middleware.CtxTokenString).(string)
	if !ok {
		custom_errors.ServerErrorResponse(w, errors.New("tokenString missing from the context"))
		return
	}
	err := h.svc.tokenSvc.DeleteToken(tokenString)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.DeleteTokenCookie(w, "REFRESH", "/auth")
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "logged out successfully"}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
