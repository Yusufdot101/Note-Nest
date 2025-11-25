package save

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/julienschmidt/httprouter"
)

func (h *saveHandler) noteIsSaved(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middleware.CtxUserKey).(*user.User)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	noteID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	isSaved, err := h.svc.noteIsSaved(u.ID, noteID)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"state": isSaved}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}

func (h *saveHandler) savedNotes(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middleware.CtxUserKey).(*user.User)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	notes, err := h.svc.savedNotes(u.ID)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"notes": notes}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}
