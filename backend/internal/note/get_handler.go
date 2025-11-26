package note

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/julienschmidt/httprouter"
)

func (h *NoteHandler) getNote(w http.ResponseWriter, r *http.Request) {
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

	note, err := h.svc.GetNote(u.ID, noteID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"note": note}, http.StatusCreated)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}

func (h *NoteHandler) getNotes(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middleware.CtxUserKey).(*user.User)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	var input struct {
		title      string
		projectID  int
		userID     int
		visibility string
		filter.Filter
	}

	qs := r.URL.Query()
	v := validator.NewValidator()

	input.title = utilities.ReadStr(qs, "title", "")
	input.visibility = utilities.ReadStr(qs, "visibility", "")
	input.userID = utilities.ReadInt(qs, "user_id", -1, v)
	input.projectID = utilities.ReadInt(qs, "project_id", -1, v)
	input.Page = utilities.ReadInt(qs, "page", 1, v)
	input.PageSize = utilities.ReadInt(qs, "page_size", 100, v)
	input.Sort = utilities.ReadStr(qs, "sort", "created_at")
	input.SafeSortList = []string{
		"id", "-id",
		"name", "-name",
		"user_id", "-user_id",
		"visibility", "-visibility",
		"created_at", "-created_at",
		"likes_count", "-likes_count",
	}

	if filter.ValidateFilter(v, &input.Filter); !v.IsValid() {
		customerrors.FailedValidationErrorResponse(w, v.Errors)
		return
	}

	notes, err := h.svc.getNotes(u.ID, input.userID, input.projectID, input.title, input.visibility, &input.Filter)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"notes": notes}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}

func (h *NoteHandler) getNoteOwner(w http.ResponseWriter, r *http.Request) {
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

	username, err := h.svc.getNoteOwner(u.ID, noteID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"username": username}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}
