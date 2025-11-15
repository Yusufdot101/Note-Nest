package note

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/julienschmidt/httprouter"
)

func (h *NoteHandler) getNote(w http.ResponseWriter, r *http.Request) {
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

	note, err := h.svc.getNote(userID, noteID)
	if err != nil {
		switch {
		case errors.Is(err, custom_errors.ErrNoRecord):
			custom_errors.NotFoundErrorResponse(w, r)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"note": note}, http.StatusCreated)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}

func (h *NoteHandler) getNotes(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.CtxUserIDKey).(int)
	if !ok {
		custom_errors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	query := r.URL.Query()

	var queryUserID, projectID *int
	var visibility string

	projectIDStr := query.Get("projectid")
	queryUserIDStr := query.Get("userid")
	visibility = query.Get("visibility")

	if projectIDStr != "" {
		i, err := strconv.Atoi(projectIDStr)
		if err != nil {
			custom_errors.BadRequestErrorResponse(w, err)
			return
		}
		projectID = &i
	}

	if queryUserIDStr != "" {
		i, err := strconv.Atoi(queryUserIDStr)
		if err != nil {
			custom_errors.BadRequestErrorResponse(w, err)
			return
		}
		queryUserID = &i
	}

	notes, err := h.svc.getNotes(&userID, queryUserID, projectID, visibility)
	if err != nil {
		switch {
		case errors.Is(err, custom_errors.ErrNoRecord):
			custom_errors.NotFoundErrorResponse(w, r)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"notes": notes}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}
