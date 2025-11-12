package project

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func TestGetProjects(t *testing.T) {
	tests := []struct {
		name           string
		userID         int
		wantStatusCode int
		wantGetCalled  bool
	}{
		{
			name:           "valid",
			userID:         1,
			wantStatusCode: http.StatusOK,
			wantGetCalled:  true,
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
			req = req.WithContext(ctx)

			repo := &MockRepo{}
			h := NewHandler(&ProjectService{
				Repo: repo,
			})

			h.GetProjects(rr, req)

			rs := rr.Result()

			if statusCode := rs.StatusCode; statusCode != test.wantStatusCode {
				t.Errorf("expected status code = %d, got status code = %d", test.wantStatusCode, statusCode)
			}

			if repo.getCalled != test.wantGetCalled {
				t.Fatalf("expected repo.getCalled = %v, got repo.getCalled = %v", test.wantGetCalled, repo.getCalled)
			}
		})
	}
}

func TestGetOneProject(t *testing.T) {
	tests := []struct {
		name             string
		userID           int
		wantStatusCode   int
		wantGetOneCalled bool
	}{
		{
			name:             "valid",
			userID:           1,
			wantStatusCode:   http.StatusOK,
			wantGetOneCalled: true,
		},
		{
			name:             "another user's",
			userID:           2,
			wantStatusCode:   http.StatusNotFound,
			wantGetOneCalled: true,
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
			h := NewHandler(&ProjectService{
				Repo: repo,
			})

			h.GetProject(rr, req)

			rs := rr.Result()

			if statusCode := rs.StatusCode; statusCode != test.wantStatusCode {
				t.Errorf("expected status code = %d, got status code = %d", test.wantStatusCode, statusCode)
			}

			if repo.getOneCalled != test.wantGetOneCalled {
				t.Fatalf("expected repo.getOneCalled = %v, got repo.getOneCalled = %v", test.wantGetOneCalled, repo.getOneCalled)
			}
		})
	}
}
