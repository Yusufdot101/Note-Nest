package project

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
)

func TestNewProjectHandler(t *testing.T) {
	tests := []struct {
		title            string
		userID           int
		payload          string
		wantStatusCode   int
		wantInsertCalled bool
	}{
		{
			title:  "valid",
			userID: 1,
			payload: `{
				"title": "project title",
				"description": "project description",
				"visibility": "private",
				"color": "#ffffff"
			}`,
			wantStatusCode:   http.StatusCreated,
			wantInsertCalled: true,
		},
		{
			title:  "missing title",
			userID: 1,
			payload: `{
				"title": "",
				"description": "project description",
				"visibility": "private",
				"color": "#ffffff"
			}`,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			title:  "unknown field",
			userID: 1,
			payload: `{
				"title": "project title",
				"description": "project description",
				"visibility": "private",
				"color": "#ffffff",
				"owner": "me"
			}`,
			wantStatusCode: http.StatusBadRequest,
		},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/projects", strings.NewReader(test.payload))
			if err != nil {
				t.Fatalf("unexpected error = %v", err)
			}

			u := &user.User{
				ID: test.userID,
			}

			ctx := context.WithValue(req.Context(), middleware.CtxUserKey, u)
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
