package like

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/julienschmidt/httprouter"
)

func (h *likeHandler) addNoteLike(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middleware.CtxUserKey).(*user.User)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("user missing from context"))
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	noteID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	err = h.svc.addLike(u.ID, noteResource, noteID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		case errors.Is(err, ErrResourceAlreadyLike):
			customerrors.BadRequestErrorResponse(w, err)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "liked note"}, http.StatusCreated)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}

func (h *likeHandler) addCommentLike(w http.ResponseWriter, r *http.Request) {
	u, ok := r.Context().Value(middleware.CtxUserKey).(*user.User)
	if !ok {
		customerrors.ServerErrorResponse(w, errors.New("user missing from context"))
		return
	}

	params := httprouter.ParamsFromContext(r.Context())
	commentID, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		customerrors.BadRequestErrorResponse(w, err)
		return
	}

	err = h.svc.addLike(u.ID, commentResource, commentID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		case errors.Is(err, ErrResourceAlreadyLike):
			customerrors.BadRequestErrorResponse(w, err)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "liked comment"}, http.StatusCreated)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}
