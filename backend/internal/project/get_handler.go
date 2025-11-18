package project

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/julienschmidt/httprouter"
)

func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.CtxUserIDKey).(int)
	if !ok {
		custom_errors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	var input struct {
		title      string
		userID     int
		visibility string
		filter.Filter
	}

	qs := r.URL.Query()
	v := validator.NewValidator()

	input.title = utilities.ReadStr(qs, "name", "")
	input.visibility = utilities.ReadStr(qs, "visibility", "")
	input.userID = utilities.ReadInt(qs, "user_id", -1, v)
	input.Page = utilities.ReadInt(qs, "page", 1, v)
	input.PageSize = utilities.ReadInt(qs, "page_size", 100, v)
	input.Sort = utilities.ReadStr(qs, "sort", "created_at")
	input.SafeSortList = []string{
		"id", "-id",
		"name", "-name",
		"user_id", "-user_id",
		"visibility", "-visibility",
		"created_at", "-created_at",
	}

	if filter.ValidateFilter(v, &input.Filter); !v.IsValid() {
		custom_errors.FailedValidationErrorResponse(w, v.Errors)
		return
	}

	projects, err := h.svc.getProjects(userID, input.userID, input.title, input.visibility, input.Filter)
	if err != nil {
		switch {
		case errors.Is(err, custom_errors.ErrNoRecord):
			custom_errors.NotFoundErrorResponse(w, r)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"projects": projects}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}

func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.CtxUserIDKey).(int)
	if !ok {
		custom_errors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	projectID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	project, err := h.svc.GetProject(userID, projectID)
	if err != nil {
		switch {
		case errors.Is(err, custom_errors.ErrNoRecord):
			custom_errors.NotFoundErrorResponse(w, r)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"project": project}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}
