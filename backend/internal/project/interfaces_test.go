package project

import (
	"testing"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

func TestValidateProject(t *testing.T) {
	tests := []struct {
		name        string
		projectName string
		description string
		visibility  string
		color       string
		wantValid   bool
		wantErrors  map[string]string
	}{
		{
			name:        "valid",
			projectName: "project name",
			description: "project description",
			visibility:  "public",
			color:       "#ffffff",
			wantValid:   true,
		},
		{
			name:        "missing name",
			projectName: "",
			description: "project description",
			visibility:  "private",
			color:       "#ffffff",
			wantValid:   false,
			wantErrors:  map[string]string{"name": "must be given"},
		},
		{
			name:        "invalid visibility value",
			projectName: "project name",
			description: "project description",
			visibility:  "unknown",
			color:       "#ffffff",
			wantValid:   false,
			wantErrors:  map[string]string{"visibility": "not allowed"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := &Project{
				Name:        test.projectName,
				Description: test.description,
				Visibility:  test.visibility,
				Color:       test.color,
			}

			v := validator.NewValidator()
			if validateProject(v, p); v.IsValid() != test.wantValid {
				t.Fatalf("expected v.IsValid = %v, got v.IsValid = %v", test.wantValid, v.IsValid())
			}
			if len(test.wantErrors) == 0 && len(v.Errors) != 0 {
				t.Fatalf("expected v.Errors to be empty, got v.Errors = %v", v.Errors)
			}
			for k, val := range test.wantErrors {
				if _, exists := v.Errors[k]; !exists {
					t.Errorf("expected %s to be in v.Errors", k)
					continue
				}
				if v.Errors[k] != test.wantErrors[k] {
					t.Errorf("expected v.Errors[%s] = %s, got v.Errors[%s] = %s", k, val, k, v.Errors[k])
				}
			}
		})
	}
}
