package project

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
)

func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	currentUserID, ok := r.Context().Value(middleware.CtxUserIDKey).(int)
	if !ok {
		custom_errors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	var userID int

	queryUserID := r.URL.Query().Get("user")
	queryProjectID := r.URL.Query().Get("project")
	visibility := r.URL.Query().Get("visibility")

	if queryUserID != "" {
		var err error
		userID, err = strconv.Atoi(queryUserID)
		if err != nil {
			custom_errors.BadRequestErrorResponse(w, err)
			return
		}
		visibility = "public"
	} else {
		userID = currentUserID
	}

	var result any
	var err error
	if queryProjectID != "" {
		projectID, err := strconv.Atoi(queryProjectID)
		if err != nil {
			custom_errors.BadRequestErrorResponse(w, err)
			return
		}

		result, err = h.svc.GetProject(currentUserID, projectID)
		if err != nil {
			switch {
			case errors.Is(err, custom_errors.ErrNoRecord):
				custom_errors.NotFoundErrorResponse(w, r)
			default:
				custom_errors.ServerErrorResponse(w, err)
			}
			return
		}
	} else {
		result, err = h.svc.getProjects(userID, visibility)
		if err != nil {
			custom_errors.BadRequestErrorResponse(w, err)
			return
		}
	}

	err = utilities.WriteJSON(w, utilities.Message{"result": result}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}
