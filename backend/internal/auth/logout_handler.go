package auth

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func (h *authHandler) logoutUser(w http.ResponseWriter, r *http.Request) {
	tokenString, ok := r.Context().Value(middleware.CtxTokenString).(string)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("tokenString missing from the context"))
		return
	}
	err := h.svc.tokenSvc.DeleteToken(tokenString)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.DeleteTokenCookie(w, "REFRESH", "/auth")
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "logged out successfully"}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
	}
}
