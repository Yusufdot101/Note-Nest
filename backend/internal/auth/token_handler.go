package auth

import (
	"fmt"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func (h *authHandler) NewAccessToken(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Context().Value(middleware.CtxTokenString)
	msg := fmt.Sprintf("refreshtoken endpoint. tokenString = %v", tokenString)
	err := utilities.WriteJSON(w, utilities.Message{"message": msg}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
