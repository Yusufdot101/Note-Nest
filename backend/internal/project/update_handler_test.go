package project

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func TestUpdateProjectHandler(t *testing.T) {
	tests := []struct {
		name             string
		userID           any
		payload          string
		wantStatusCode   int
		wantUpdateCalled bool
	}{
		{
			name:   "valid",
			userID: 1,
			payload: `{
				"name": "project name",
				"description": "project description",
				"visibility": "private",
				"color": "#ffffff"
			}`,
			wantStatusCode:   http.StatusOK,
			wantUpdateCalled: true,
		},
		{
			name:   "valid-missing name",
			userID: 1,
			payload: `{
				"description": "project description",
				"visibility": "private",
				"color": "#ffffff"
			}`,
			wantStatusCode:   http.StatusOK,
			wantUpdateCalled: true,
		},
		{
			name:   "another user's",
			userID: 2,
			payload: `{
				"description": "project description",
				"visibility": "private",
				"color": "#ffffff"
			}`,
			wantStatusCode:   http.StatusNotFound,
			wantUpdateCalled: false,
		},
		{
			name:   "invalid inputs",
			userID: 1,
			payload: `{
				"description": "project description",
				"visibility": "private",
				"color": "ffffff"
			}`,
			wantStatusCode:   http.StatusBadRequest,
			wantUpdateCalled: false,
		},
		{
			name:   "userID missing from context",
			userID: nil,
			payload: `{
				"name": "project name",
				"description": "project description",
				"visibility": "private",
				"color": "#ffffff"
			}`,
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:   "unknown field",
			userID: 1,
			payload: `{
				"name": "project name",
				"description": "project description",
				"visibility": "private",
				"color": "#ffffff",
				"owner": "me"
			}`,
			wantStatusCode: http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPatch, "/projects", strings.NewReader(test.payload))
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

			h.UpdateProject(rr, req)

			rs := rr.Result()
			if rs.StatusCode != test.wantStatusCode {
				t.Fatalf("expected status code = %d, got status code = %d", test.wantStatusCode, rs.StatusCode)
			}

			if repo.updateCalled != test.wantUpdateCalled {
				t.Fatalf("expected repo.updateCalled = %v, got repo.updateCalled = %v", test.wantUpdateCalled, repo.updateCalled)
			}
		})
	}
}
