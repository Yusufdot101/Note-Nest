package project

import (
	"strings"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (ps *ProjectService) newProject(v *validator.Validator, userID int, name, description, visibility, color string) error {
	cleanedName := strings.TrimSpace(name)
	cleanedDescription := strings.TrimSpace(description)
	cleanedVisibility := strings.ToLower(strings.TrimSpace(visibility))
	cleanColor := strings.ToLower(strings.TrimSpace(color))
	p := &Project{
		UserID:      userID,
		Name:        cleanedName,
		Description: cleanedDescription,
		Visibility:  cleanedVisibility,
		Color:       cleanColor,
	}
	if validateProject(v, p); !v.IsValid() {
		return validator.ErrFailedValidation
	}

	err := ps.Repo.insert(p)
	if err != nil {
		return err
	}

	return nil
}

func (ps *ProjectService) getProjects(userID int, visibility string) ([]*Project, error) {
	return ps.Repo.get(userID, visibility)
}
