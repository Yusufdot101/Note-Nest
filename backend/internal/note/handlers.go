package note

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/julienschmidt/httprouter"
)

func (h *NoteHandler) newNote(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.CtxUserIDKey).(int)
	if !ok {
		custom_errors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}
	params := httprouter.ParamsFromContext(r.Context())
	projectID, err := strconv.Atoi(params.ByName("projectid"))
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}
	var input struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		Visibility string `json:"visibility"`
		Color      string `json:"color"`
	}

	err = utilities.ReadJSON(w, r, &input)
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	err = h.svc.newNote(v, userID, projectID, input.Title, input.Content, input.Visibility, input.Color)
	if err != nil {
		switch {
		case errors.Is(err, validator.ErrFailedValidation):
			custom_errors.FailedValidationErrorResponse(w, v.Errors)
		case errors.Is(err, custom_errors.ErrNoRecord):
			custom_errors.NotFoundErrorResponse(w, r)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "note created successfully"}, http.StatusCreated)
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

	params := httprouter.ParamsFromContext(r.Context())
	projectID, err := strconv.Atoi(params.ByName("projectid"))
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	query := r.URL.Query()
	queryNoteID := query.Get("noteid")
	var noteID int
	if queryNoteID != "" {
		var err error
		noteID, err = strconv.Atoi(query.Get("noteid"))
		if err != nil {
			custom_errors.BadRequestErrorResponse(w, err)
			return
		}
	}
	visibility := query.Get("visibility")

	notes, err := h.svc.getNotes(userID, projectID, noteID, visibility)
	if err != nil {
		switch {
		case errors.Is(err, custom_errors.ErrNoRecord):
			custom_errors.NotFoundErrorResponse(w, r)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"notes": notes}, http.StatusCreated)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}

func (h *NoteHandler) deleteNote(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.CtxUserIDKey).(int)
	if !ok {
		custom_errors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	projectID, err := strconv.Atoi(params.ByName("projectid"))
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	query := r.URL.Query()
	queryNoteID := query.Get("noteid")
	if queryNoteID == "" {
		custom_errors.BadRequestErrorResponse(w, errors.New("noteid must be given"))
		return
	}

	var noteID int
	noteID, err = strconv.Atoi(query.Get("noteid"))
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	err = h.svc.deleteNote(userID, projectID, noteID)
	if err != nil {
		switch {
		case errors.Is(err, custom_errors.ErrNoRecord):
			custom_errors.NotFoundErrorResponse(w, r)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "note deleted successfully"}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}
