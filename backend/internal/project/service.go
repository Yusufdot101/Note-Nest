package project

import (
	"context"
	"fmt"
	"strings"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/redis/go-redis/v9"
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

func (ps *ProjectService) getProjects(currentUserID, userID int, title, visibility string, f *filter.Filter) ([]*Project, *filter.Metadata, error) {
	return ps.Repo.get(currentUserID, userID, title, visibility, f)
}

func (ps *ProjectService) GetProject(userID, projectID int) (*Project, error) {
	project := &Project{}
	ctx := context.Background()
	// fetch the project
	err := utilities.GetUnmarshalRedisKey(ps.RDB, ctx, fmt.Sprintf("project:%d", projectID), project)
	if err != nil {
		if err == redis.Nil {
			project, err = ps.Repo.getOne(projectID)
			if err != nil {
				return nil, err
			}

			utilities.SetRedisKey(ps.RDB, ctx, fmt.Sprintf("project:%d", projectID), project, 0)
		} else {
			return nil, err
		}
	}

	// only allow the owner to see private projects
	if project.UserID != userID && project.Visibility != "public" {
		return nil, customerrors.ErrNoRecord
	}

	return project, nil
}

func (ps *ProjectService) deleteProject(userID, projectID int) error {
	project := &Project{}
	ctx := context.Background()
	// fetch the project
	err := utilities.GetUnmarshalRedisKey(ps.RDB, ctx, fmt.Sprintf("project:%d", projectID), project)
	if err != nil {
		if err == redis.Nil {
			project, err = ps.Repo.getOne(projectID)
			if err != nil {
				return err
			}

		} else {
			return err
		}
	}
	// can only delete your projects
	if project.UserID != userID {
		return customerrors.ErrNoRecord
	}

	// remove the stale project from the cache
	err = utilities.DeleteRedisKey(ps.RDB, ctx, fmt.Sprintf("project:%d", projectID))
	if err != nil {
		return err
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

	// remove the stale project from the cache
	ctx := context.Background()
	err := utilities.DeleteRedisKey(ps.RDB, ctx, fmt.Sprintf("project:%d", projectID))
	if err != nil {
		return err
	}

	return ps.Repo.update(userID, projectID, title, description, visibility, color)
}

func (ps *ProjectService) updateProjectVisibility(
	v *validator.Validator, userID, projectID int, visibility string,
) error {
	cleanedVisibility := strings.TrimSpace(visibility)
	if validateVisibility(v, cleanedVisibility); !v.IsValid() {
		return validator.ErrFailedValidation
	}

	// remove the stale project from the cache
	ctx := context.Background()
	err := utilities.DeleteRedisKey(ps.RDB, ctx, fmt.Sprintf("project:%d", projectID))
	if err != nil {
		return err
	}

	return ps.Repo.updateProjectVisibility(userID, projectID, cleanedVisibility)
}

func (ps *ProjectService) updateProjectColor(
	v *validator.Validator, userID, projectID int, color string,
) error {
	cleanedColor := strings.TrimSpace(color)
	if validateColor(v, cleanedColor); !v.IsValid() {
		return validator.ErrFailedValidation
	}

	// remove the stale project from the cache
	ctx := context.Background()
	err := utilities.DeleteRedisKey(ps.RDB, ctx, fmt.Sprintf("project:%d", projectID))
	if err != nil {
		return err
	}

	return ps.Repo.updateProjectColor(userID, projectID, cleanedColor)
}
