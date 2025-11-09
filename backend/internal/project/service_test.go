package project

import (
	"testing"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

func TestNewProject(t *testing.T) {
	tests := []struct {
		name             string
		userID           int
		projectName      string
		description      string
		visibility       string
		color            string
		wantErr          bool
		wantInsertCalled bool
	}{
		{
			name:             "valid",
			userID:           1,
			projectName:      "project name",
			description:      "project description",
			visibility:       "private",
			color:            "#ffffff",
			wantInsertCalled: true,
		},
		{
			name:        "invalid - empty name",
			userID:      1,
			projectName: "",
			description: "project description",
			color:       "#ffffff",
			visibility:  "private",
			wantErr:     true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			repo := &MockRepo{}
			svc := ProjectService{
				Repo: repo,
			}
			v := validator.NewValidator()
			err := svc.newProject(v, test.userID, test.projectName, test.description, test.visibility, test.color)
			if (err != nil) != test.wantErr {
				t.Fatalf("expected err = %v, got err = %v", test.wantErr, err)
			}

			if repo.insertCalled != test.wantInsertCalled {
				t.Fatalf("expected repo.insertCalled = %v, got repo.insertCalled = %v", test.wantInsertCalled, repo.insertCalled)
			}
		})
	}
}
