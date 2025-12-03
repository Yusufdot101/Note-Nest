package project

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

func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middleware.CtxUserKey).(*user.User)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("userID missing from context"))
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

	input.title = utilities.ReadStr(qs, "title", "")
	input.visibility = utilities.ReadStr(qs, "visibility", "")
	input.userID = utilities.ReadInt(qs, "user_id", -1, v)
	input.Page = utilities.ReadInt(qs, "page", 1, v)
	input.PageSize = utilities.ReadInt(qs, "page_size", 100, v)
	input.Sort = utilities.ReadStr(qs, "sort", "likes_count")
	input.Order = utilities.ReadStr(qs, "order", "descending")
	input.SafeSortList = []string{
		"id",
		"title",
		"user_id",
		"visibility",
		"created_at",
		"likes_count",
		"entries_count",
		"comments_count",
	}

	if filter.ValidateFilter(v, &input.Filter); !v.IsValid() {
		customerrors.FailedValidationErrorResponse(w, v.Errors)
		return
	}

	projects, err := h.svc.getProjects(u.ID, input.userID, input.title, input.visibility, input.Filter)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"projects": projects}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}

func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middleware.CtxUserKey).(*user.User)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	projectID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	project, err := h.svc.GetProject(u.ID, projectID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"project": project}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}
