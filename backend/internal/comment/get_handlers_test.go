package comment

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

func TestGetCommentsHandler(t *testing.T) {
	u := &user.User{
		ID: 1,
	}
	wantStatusCode := http.StatusOK
	wantGetCalled := true

	t.Run("test get comments", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/notes/1/comments", nil)

		ctx := context.WithValue(req.Context(), middleware.CtxUserKey, u)
		params := httprouter.Params{httprouter.Param{Key: "id", Value: "1"}}
		ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
		req = req.WithContext(ctx)

		repo := &mockRepo{}
		h := newHandler(&commentService{
			repo: repo,
		})

		h.getComments(recorder, req)

		rs := recorder.Result()

		if statusCode := rs.StatusCode; statusCode != wantStatusCode {
			t.Errorf("expected status code = %d, got status code = %d", wantStatusCode, statusCode)
		}

		if repo.getCalled != wantGetCalled {
			t.Fatalf("expected repo.getCalled = %v, got repo.getCalled = %v", wantGetCalled, repo.getCalled)
		}
	})
}
