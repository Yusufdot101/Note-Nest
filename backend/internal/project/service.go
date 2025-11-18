package project

import (
	"strings"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (ps *ProjectService) newProject(v *validator.Validator, userID int, title, description, visibility, color string) error {
	cleanedTitle := strings.TrimSpace(title)
	cleanedDescription := strings.TrimSpace(description)
	cleanedVisibility := strings.ToLower(strings.TrimSpace(visibility))
	cleanedColor := strings.ToLower(strings.TrimSpace(color))
	p := &Project{
		UserID:      userID,
		Title:       cleanedTitle,
		Description: cleanedDescription,
		Visibility:  cleanedVisibility,
		Color:       cleanedColor,
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

func (ps *ProjectService) GetProject(userID, projectID int) (*Project, error) {
	project, err := ps.Repo.getOne(projectID)
	if err != nil {
		return nil, err
	}

	// only allow the owner to see private projects
	if project.UserID != userID && project.Visibility != "public" {
		return nil, custom_errors.ErrNoRecord
	}

	return project, nil
}

func (ps *ProjectService) deleteProject(userID, projectID int) error {
	project, err := ps.Repo.getOne(projectID)
	if err != nil {
		return err
	}
	// can only delete your projects
	if project.UserID != userID {
		return custom_errors.ErrNoRecord
	}

	return ps.Repo.delete(project.ID)
}

func (ps *ProjectService) updateProject(v *validator.Validator, userID, projectID int, title, description, visibility, color *string) error {
	if title != nil {
		cleanedTitle := strings.TrimSpace(*title)
		validateTitle(v, cleanedTitle)
	}

	if visibility != nil {
		cleanedVisibility := strings.TrimSpace(*visibility)
		validateVisibility(v, cleanedVisibility)
	}

	if color != nil {
		cleanedColor := strings.TrimSpace(*color)
		validateColor(v, cleanedColor)
	}

	if !v.IsValid() {
		return validator.ErrFailedValidation
	}

	return ps.Repo.update(userID, projectID, title, description, visibility, color)
}
