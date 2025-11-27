package like

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

func TestNoteIsLikedHandler(t *testing.T) {
	u := &user.User{
		ID: 1,
	}
	noteID := "1"

	wantStatusCode := http.StatusOK

	req := httptest.NewRequest(http.MethodPost, "/notes/:id/like", nil)
	ctx := context.WithValue(req.Context(), middleware.CtxUserKey, u)
	params := httprouter.Params{httprouter.Param{Key: "id", Value: noteID}}
	ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
	req = req.WithContext(ctx)

	recorder := httptest.NewRecorder()

	repo := &mockRepo{}
	h := newHandler(&likeService{
		repo: repo,
	})

	h.noteIsLiked(recorder, req)
	res := recorder.Result()

	if statusCode := res.StatusCode; statusCode != wantStatusCode {
		t.Errorf("expected status code = %d, got status code = %d", wantStatusCode, statusCode)
	}

	if !repo.isLikedCalled {
		t.Fatalf("expected repo.isLiked to be called")
	}
}

func TestCommentIsLikedHandler(t *testing.T) {
	u := &user.User{
		ID: 1,
	}
	noteID := "1"

	wantStatusCode := http.StatusOK

	req := httptest.NewRequest(http.MethodPost, "/comments/:id/like", nil)
	ctx := context.WithValue(req.Context(), middleware.CtxUserKey, u)
	params := httprouter.Params{httprouter.Param{Key: "id", Value: noteID}}
	ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
	req = req.WithContext(ctx)

	recorder := httptest.NewRecorder()

	repo := &mockRepo{}
	h := newHandler(&likeService{
		repo: repo,
	})

	h.commentIsLiked(recorder, req)
	res := recorder.Result()

	if statusCode := res.StatusCode; statusCode != wantStatusCode {
		t.Errorf("expected status code = %d, got status code = %d", wantStatusCode, statusCode)
	}

	if !repo.isLikedCalled {
		t.Fatalf("expected repo.isLiked to be called")
	}
}
