package save

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func TestUnsaveNoteHandler(t *testing.T) {
	userID := 1
	noteID := "1"

	wantStatusCode := http.StatusOK

	req := httptest.NewRequest(http.MethodPost, "/notes/:id/like", nil)
	ctx := context.WithValue(req.Context(), middleware.CtxUserIDKey, userID)
	params := httprouter.Params{httprouter.Param{Key: "id", Value: noteID}}
	ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
	req = req.WithContext(ctx)

	recorder := httptest.NewRecorder()

	repo := &mockRepo{}
	h := newHandler(&saveService{
		repo: repo,
	})

	h.unsaveNote(recorder, req)
	res := recorder.Result()

	if statusCode := res.StatusCode; statusCode != wantStatusCode {
		t.Errorf("expected status code = %d, got status code = %d", wantStatusCode, statusCode)
	}

	if !repo.deleteCalled {
		t.Fatalf("expected repo.deletCalled to be called")
	}
}
