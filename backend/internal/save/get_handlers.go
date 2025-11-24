package save

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/julienschmidt/httprouter"
)

func (h *saveHandler) noteIsSaved(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.CtxUserIDKey).(int)
	if !ok {
		custom_errors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	noteID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	isSaved, err := h.svc.noteIsSaved(userID, noteID)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"state": isSaved}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}
