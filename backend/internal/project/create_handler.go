package project

import (
	"errors"
	"net/http"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (h *ProjectHandler) NewProject(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Visibility  string `json:"visibility"`
		Color       string `json:"color"`
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

	v := validator.NewValidator()
	err = h.svc.newProject(v, u.ID, input.Title, input.Description, input.Visibility, input.Color)
	if err != nil {
		switch {
		case errors.Is(err, validator.ErrFailedValidation):
			customerrors.FailedValidationErrorResponse(w, v.Errors)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "project created successfully"}, http.StatusCreated)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}
