package auth

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/token"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func (h *authHandler) newAccessToken(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.CtxUserIDKey).(int)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("invalid userID format in the context"))
		return
	}
	token, err := h.svc.tokenSvc.NewToken(token.JWT, token.ACCESS, userID)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
	err = utilities.WriteJSON(w, utilities.Message{"access_token": token}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
	}
}
