package project

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
)

func TestNewProjectHandler(t *testing.T) {
	tests := []struct {
		name             string
		userID           any
		payload          string
		wantStatusCode   int
		wantInsertCalled bool
	}{
		{
			name:   "valid",
			userID: 1,
			payload: `{
				"name": "project name",
				"description": "project description",
				"visibility": "private"
			}`,
			wantStatusCode:   http.StatusCreated,
			wantInsertCalled: true,
		},
		{
			name:   "missing name",
			userID: 1,
			payload: `{
				"name": "",
				"description": "project description",
				"visibility": "private"
			}`,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:   "userID missing from context",
			userID: nil,
			payload: `{
				"name": "project name",
				"description": "project description",
				"visibility": "private"
			}`,
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name:   "unknown field",
			userID: 1,
			payload: `{
				"name": "project name",
				"description": "project description",
				"visibility": "private"
				"owner": "me"
			}`,
			wantStatusCode: http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/projects", strings.NewReader(test.payload))
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}
			ctx := context.WithValue(req.Context(), middleware.CtxUserIDKey, test.userID)
			req = req.WithContext(ctx)

			repo := &MockRepo{}
			h := NewHandler(&ProjectService{
				Repo: repo,
			})

			h.NewProject(rr, req)

			rs := rr.Result()
			if rs.StatusCode != test.wantStatusCode {
				t.Fatalf("expected status code = %d, got status code = %d", test.wantStatusCode, rs.StatusCode)
			}

			if repo.insertCalled != test.wantInsertCalled {
				t.Fatalf("expected repo.insertCalled = %v, got repo.insertCalled = %v", test.wantInsertCalled, repo.insertCalled)
			}
		})
	}
}
