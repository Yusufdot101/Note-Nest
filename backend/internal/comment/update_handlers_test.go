package comment

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

func TestUpdateCommentHandler(t *testing.T) {
	tests := []struct {
		name             string
		newContent       string
		wantStatusCode   int
		wantUpdateCalled bool
	}{
		{
			name:             "valid comment",
			newContent:       "this is valid new content",
			wantStatusCode:   http.StatusOK,
			wantUpdateCalled: true,
		},
		{
			name:             "invalid: missing content",
			newContent:       "",
			wantStatusCode:   http.StatusBadRequest,
			wantUpdateCalled: false,
		},
	}

	u := &user.User{
		ID: 1,
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			payload := fmt.Sprintf(`{
					"content": "%s"
			}`, test.newContent)

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPatch, "/notes/1/comments", strings.NewReader(payload))

			ctx := context.WithValue(req.Context(), middleware.CtxUserKey, u)
			params := httprouter.Params{httprouter.Param{Key: "id", Value: "1"}}
			ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
			req = req.WithContext(ctx)

			repo := &mockRepo{}
			h := newHandler(&commentService{
				repo: repo,
			})

			h.updateComment(recorder, req)

			rs := recorder.Result()

			if statusCode := rs.StatusCode; statusCode != test.wantStatusCode {
				t.Errorf("expected status code = %d, got status code = %d", test.wantStatusCode, statusCode)
			}

			if repo.updateCalled != test.wantUpdateCalled {
				t.Fatalf("expected repo.updateCalled = %v, got repo.updateCalled = %v", test.wantUpdateCalled, repo.updateCalled)
			}
		})
	}
}
