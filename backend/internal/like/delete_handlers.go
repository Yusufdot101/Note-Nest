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

func (h *likeHandler) removeNoteLike(w http.ResponseWriter, r *http.Request) {
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

	err = h.svc.removeLike(u.ID, noteResource, noteID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "unliked note"}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}

func (h *likeHandler) removeCommentLike(w http.ResponseWriter, r *http.Request) {
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

	err = h.svc.removeLike(u.ID, commentResource, commentID)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"message": "unliked comment"}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}
