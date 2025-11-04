package user

import (
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func (h *userHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	err := logoutUser()
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
	msg := "logout endpoint"
	err = utilities.WriteJSON(w, utilities.Message{"message": msg}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
	}
}
