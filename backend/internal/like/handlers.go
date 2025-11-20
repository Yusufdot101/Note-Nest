package like

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/julienschmidt/httprouter"
)

func (h *likeHandler) addLike(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int `json:"user_id"`
	}

	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	noteID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	err = h.svc.addLike(input.UserID, noteID)
	if err != nil {
		switch {
		case errors.Is(err, custom_errors.ErrNoRecord):
			custom_errors.NotFoundErrorResponse(w, r)
		case errors.Is(err, ErrNoteAlreadyLike):
			custom_errors.BadRequestErrorResponse(w, err)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "liked note"}, http.StatusCreated)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}

func (h *likeHandler) removeLike(w http.ResponseWriter, r *http.Request) {
	var input struct {
		UserID int `json:"user_id"`
	}

	err := utilities.ReadJSON(w, r, &input)
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	noteID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		custom_errors.BadRequestErrorResponse(w, err)
		return
	}

	err = h.svc.removeLike(input.UserID, noteID)
	if err != nil {
		switch {
		case errors.Is(err, custom_errors.ErrNoRecord):
			custom_errors.NotFoundErrorResponse(w, r)
		default:
			custom_errors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "unliked note"}, http.StatusOK)
	if err != nil {
		custom_errors.ServerErrorResponse(w, err)
		return
	}
}
