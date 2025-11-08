package project

import (
	"strings"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (ps *ProjectService) newProject(v *validator.Validator, userID int, name, description, visibility string) error {
	cleanedName := strings.TrimSpace(name)
	cleanedDesription := strings.TrimSpace(description)
	cleanedVisibilyt := strings.ToLower(strings.TrimSpace(visibility))
	p := &Project{
		UserID:      userID,
		Name:        cleanedName,
		Description: cleanedDesription,
		Visibility:  cleanedVisibilyt,
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
