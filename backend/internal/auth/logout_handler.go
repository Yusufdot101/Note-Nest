package auth

import (
	"fmt"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func (h *authHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey)
	msg := fmt.Sprintf("logout endpoint. userID = %v", userID)
	err := utilities.WriteJSON(w, utilities.Message{"message": msg}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
