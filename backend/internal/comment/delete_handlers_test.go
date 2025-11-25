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

func TestDeleteCommentHandler(t *testing.T) {
	u := &user.User{
		ID: 1,
	}
	wantStatusCode := http.StatusOK
	wantDeleteCalled := true

	t.Run("test delete comment", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodDelete, "/comments/1", nil)

		ctx := context.WithValue(req.Context(), middleware.CtxUserKey, u)
		params := httprouter.Params{httprouter.Param{Key: "id", Value: "1"}}
		ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
		req = req.WithContext(ctx)

		repo := &mockRepo{}
		h := newHandler(&commentService{
			repo: repo,
		})

		h.deleteComment(recorder, req)

		rs := recorder.Result()

		if statusCode := rs.StatusCode; statusCode != wantStatusCode {
			t.Errorf("expected status code = %d, got status code = %d", wantStatusCode, statusCode)
		}

		if repo.deleteCalled != wantDeleteCalled {
			t.Fatalf("expected repo.deleteCalled = %v, got repo.deleteCalled = %v", wantDeleteCalled, repo.deleteCalled)
		}
	})
}
