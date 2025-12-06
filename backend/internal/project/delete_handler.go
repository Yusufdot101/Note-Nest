package project

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/julienschmidt/httprouter"
)

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
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

	err = h.svc.deleteProject(u.ID, projectID)
	log.Println("here: ", err)
	if err != nil {
		switch {
		case errors.Is(err, customerrors.ErrNoRecord):
			customerrors.NotFoundErrorResponse(w, r)
		default:
			customerrors.ServerErrorResponse(w, err)
		}
		return
	}

	err = utilities.WriteJSON(w, utilities.Message{"project": "deleted successfully"}, http.StatusOK)
	if err != nil {
		customerrors.ServerErrorResponse(w, err)
		return
	}
}
