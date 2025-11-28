package project

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/julienschmidt/httprouter"
)

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       *string `json:"title"`
		Description *string `json:"description"`
		Visibility  *string `json:"visibility"`
		Color       *string `json:"color"`
	}

	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

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

	v := validator.NewValidator()
	err = h.svc.updateProject(v, u.ID, projectID, input.Title, input.Description, input.Visibility, input.Color)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		case errors.Is(err, validator.ErrFailedValidation):
			customerrors.FailedValidationErrorResponse(w, v.Errors)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "project updated successfully"}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}

func (h *ProjectHandler) updateProjectVisibility(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middleware.CtxUserKey).(*user.User)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	var input struct {
		Visibility string `json:"visibility"`
	}

	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	noteID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	err = h.svc.updateProjectVisibility(v, u.ID, noteID, input.Visibility)
	if err != nil {
		switch {
		case errors.Is(err, validator.ErrFailedValidation):
			customerrors.FailedValidationErrorResponse(w, v.Errors)
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		case errors.Is(err, customerrors.ErrUpdateTimeout):
			customerrors.UpdateTimeoutErrorResponse(w)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "project visibility updated successfully"}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
	}
}

func (h *ProjectHandler) updateProjectColor(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middleware.CtxUserKey).(*user.User)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("userID missing from context"))
		return
	}

	var input struct {
		Color string `json:"color"`
	}

	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	noteID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	v := validator.NewValidator()
	err = h.svc.updateProjectColor(v, u.ID, noteID, input.Color)
	if err != nil {
		switch {
		case errors.Is(err, validator.ErrFailedValidation):
			customerrors.FailedValidationErrorResponse(w, v.Errors)
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		case errors.Is(err, customerrors.ErrUpdateTimeout):
			customerrors.UpdateTimeoutErrorResponse(w)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "project color updated successfully"}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
	}
}
