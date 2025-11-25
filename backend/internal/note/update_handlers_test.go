package note

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yusufdot101/note-nest/internal/middleware"
	"github.com/Yusufdot101/note-nest/internal/user"
	"github.com/julienschmidt/httprouter"
)

func TestUpdateNoteTitleContentHandler(t *testing.T) {
	tests := []struct {
		name             string
		payload          string
		wantStatusCode   int
		wantUpdateCalled bool
	}{
		{
			name: "valid inputs",
			payload: `{
				"title": "updated title",
				"content": "updated content"
			}`,
			wantStatusCode:   http.StatusOK,
			wantUpdateCalled: true,
		},
		{
			name: "valid: one field not provided",
			payload: `{
				"content": "updated content"
			}`,
			wantStatusCode:   http.StatusOK,
			wantUpdateCalled: true,
		},
		{
			name: "invalid: empty title",
			payload: `{
				"title": ""
			}`,
			wantStatusCode:   http.StatusBadRequest,
			wantUpdateCalled: false,
		},
		{
			name:             "invalid: both fields not provided",
			wantStatusCode:   http.StatusBadRequest,
			wantUpdateCalled: false,
		},
	}

	u := &user.User{
		ID: 1,
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPatch, "/notes/1", strings.NewReader(test.payload))

			ctx := context.WithValue(req.Context(), middleware.CtxUserKey, u)
			params := httprouter.Params{httprouter.Param{Key: "id", Value: "1"}}
			ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
			req = req.WithContext(ctx)

			repo := &mockRepo{}
			h := newHandler(&NoteService{
				Repo: repo,
			})

			h.updateNoteTitleContent(recorder, req)

			rs := recorder.Result()

			if statusCode := rs.StatusCode; statusCode != test.wantStatusCode {
				t.Errorf("expected status code = %d, got status code = %d", test.wantStatusCode, statusCode)
			}

			if repo.updateNoteTitleContentCalled != test.wantUpdateCalled {
				t.Fatalf("expected repo.updateNoteTitleContentCalled = %v, got repo.updateNoteTitleContentCalled = %v", test.wantUpdateCalled, repo.updateNoteTitleContentCalled)
			}
		})
	}
}

func TestUpdateNoteColorHandler(t *testing.T) {
	tests := []struct {
		name             string
		payload          string
		wantStatusCode   int
		wantUpdateCalled bool
	}{
		{
			name: "valid inputs",
			payload: `{
				"color": "#ffffff"
			}`,
			wantStatusCode:   http.StatusOK,
			wantUpdateCalled: true,
		},
		{
			name: "invalid: empty value",
			payload: `{
				"color": ""
			}`,
			wantStatusCode:   http.StatusBadRequest,
			wantUpdateCalled: false,
		},
		{
			name: "invalid: invalid value",
			payload: `{
				"color": "white"
			}`,
			wantStatusCode:   http.StatusBadRequest,
			wantUpdateCalled: false,
		},
	}

	u := &user.User{
		ID: 1,
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPatch, "/notes/1", strings.NewReader(test.payload))

			ctx := context.WithValue(req.Context(), middleware.CtxUserKey, u)
			params := httprouter.Params{httprouter.Param{Key: "id", Value: "1"}}
			ctx = context.WithValue(ctx, httprouter.ParamsKey, params)
			req = req.WithContext(ctx)

			repo := &mockRepo{}
			h := newHandler(&NoteService{
				Repo: repo,
			})

			h.updateNoteColor(recorder, req)

			rs := recorder.Result()

			if statusCode := rs.StatusCode; statusCode != test.wantStatusCode {
				t.Errorf("expected status code = %d, got status code = %d", test.wantStatusCode, statusCode)
			}

			if repo.updateNoteColorCalled != test.wantUpdateCalled {
				t.Fatalf("expected repo.updateNoteColorCalled = %v, got repo.updateNoteColorCalled = %v", test.wantUpdateCalled, repo.updateNoteColorCalled)
			}
		})
	}
}
