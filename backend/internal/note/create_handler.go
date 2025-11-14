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
