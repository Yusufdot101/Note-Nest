package auth

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func (h *authHandler) NewAccessToken(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.CtxUserIDKey).(int)
	if !ok {
		custom_errors.ServerErrorResponse(w, errors.New("invalid userID format in the context"))
		return
	}
	token, err := h.svc.tokenSvc.NewToken(token.JWT, token.ACCESS, userID)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
	err = utilities.WriteJSON(w, utilities.Message{"token": token}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
