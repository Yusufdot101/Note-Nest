package project

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/utilities"
<<<<<<< HEAD
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
=======
	"github.com/julienschmidt/httprouter"
)

func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	var userID int
	queryUserID := r.URL.Query().Get("user")
>>>>>>> 9946a38 (port monolithic handlers.go into refactored files)
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
<<<<<<< HEAD
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
=======
		var ok bool
		userID, ok = r.Context().Value(middleware.CtxUserIDKey).(int)
		if !ok {
			custom_errors.ServerErrorResponse(w, errors.New("userID missing from context"))
>>>>>>> 9946a38 (port monolithic handlers.go into refactored files)
			return
		}
	}

<<<<<<< HEAD
	err = utilities.WriteJSON(w, utilities.Message{"result": result}, http.StatusOK)
=======
	projects, err := h.svc.getProjects(userID, visibility)
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

	project, err := h.svc.getProject(userID, projectID)
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
>>>>>>> 9946a38 (port monolithic handlers.go into refactored files)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}
