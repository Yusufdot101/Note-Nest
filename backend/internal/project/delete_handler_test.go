package project

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func TestDeleteProject(t *testing.T) {
	tests := []struct {
		name             string
		userID           int
		wantStatusCode   int
		wantDeleteCalled bool
		wantGetOneCalled bool
	}{
		{
			name:             "valid",
			userID:           1,
			wantStatusCode:   http.StatusOK,
			wantDeleteCalled: true,
			wantGetOneCalled: true,
		},
		{
			name:             "another user's project",
			userID:           2,
			wantStatusCode:   http.StatusNotFound,
			wantGetOneCalled: true,
			wantDeleteCalled: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodDelete, "/projects", nil)
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}

			ctx := context.WithValue(req.Context(), middleware.CtxUserIDKey, test.userID)
			params := httprouter.Params{httprouter.Param{Key: "id", Value: "123"}}
			ctx = context.WithValue(ctx, httprouter.ParamsKey, params)

			req = req.WithContext(ctx)

			repo := &MockRepo{}
			h := NewHandler(&ProjectService{Repo: repo})

			h.DeleteProject(rr, req)

			rs := rr.Result()
			if statusCode := rs.StatusCode; statusCode != test.wantStatusCode {
				t.Errorf("expected status code = %d, got status code = %d", test.wantStatusCode, statusCode)
			}

			if repo.getOneCalled != test.wantGetOneCalled {
				t.Fatalf("expected repo.getOneCalled = %v, got repo.getOneCalled = %v", test.wantGetOneCalled, repo.getOneCalled)
			}

			if repo.deleteCalled != test.wantDeleteCalled {
				t.Fatalf("expected repo.deleteCalled = %v, got repo.deleteCalled = %v", test.wantDeleteCalled, repo.deleteCalled)
			}
		})
	}
}
