package note

import (
	"context"
	"fmt"
	"strings"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/project"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/redis/go-redis/v9"
)

func (ns *NoteService) newNote(
	v *validator.Validator, userID, projectID int, title, content, visibility, color string,
) error {
	// fetch the project
	p, err := ns.ProjectSvc.GetProject(userID, projectID)
	if err != nil {
		return err
	}

	// do checks
	cleanedTitle := strings.TrimSpace(title)
	cleanedContent := strings.TrimSpace(content)
	cleanedVisibility := strings.ToLower(strings.TrimSpace(visibility))
	cleanColor := strings.ToLower(strings.TrimSpace(color))

	n := &Note{
		ProjectID:  projectID,
		Title:      cleanedTitle,
		Content:    cleanedContent,
		Visibility: cleanedVisibility,
		Color:      cleanColor,
	}

	// cannot create in other's projects
	if p.UserID != userID {
		return customerrors.ErrNoRecord
	}
	if p.Visibility == "private" && n.Visibility == "public" {
		v.AddError("entry", "cannot be more public than the project")
		return validator.ErrFailedValidation
	}
	if validateNote(v, n); !v.IsValid() {
		return validator.ErrFailedValidation
	}

	// save to db or return err
	err = ns.Repo.insert(n)
	if err != nil {
		return err
	}

	return nil
}

func (ns *NoteService) GetNote(userID, noteID int) (*Note, error) {
	// fetch the note
	note := &Note{}
	p := &project.Project{}
	ctx := context.Background()

	err := utilities.GetUnmarshalRedisKey(ns.RDB, ctx, fmt.Sprintf("note:%d", noteID), note)
	if err != nil {
		if err == redis.Nil {
			note, err = ns.Repo.get(noteID)
			if err != nil {
				return nil, err
			}

			utilities.SetRedisKey(ns.RDB, ctx, fmt.Sprintf("note:%d", noteID), note, 0)
		} else {
			return nil, err
		}
	}

	// fetch the project
	p, err = ns.ProjectSvc.GetProject(userID, note.ProjectID)
	if err != nil {
		return nil, err
	}

	// do checks:
	// cannot see other's private projects/notes
	if p.UserID != userID && (p.Visibility == "private" || note.Visibility == "private") {
		return nil, customerrors.ErrNoRecord
	}

	// return note
	return note, nil
}

func (ns *NoteService) getNotes(currentUserID, queryUserID, projectID int, title, visibility string, filter *filter.Filter) ([]*Note, *filter.Metadata, error) {
	return ns.Repo.getMany(currentUserID, queryUserID, projectID, title, visibility, filter)
}

func (ns *NoteService) deleteNote(userID, noteID int) error {
	note := &Note{}
	p := &project.Project{}
	ctx := context.Background()

	err := utilities.GetUnmarshalRedisKey(ns.RDB, ctx, fmt.Sprintf("note:%d", noteID), note)
	if err != nil {
		if err == redis.Nil {
			note, err = ns.Repo.get(noteID)
			if err != nil {
				return err
			}

		} else {
			return err
		}
	}

	// fetch the project
	err = utilities.GetUnmarshalRedisKey(ns.RDB, ctx, fmt.Sprintf("project:%d", note.ProjectID), p)
	if err != nil {
		if err == redis.Nil {
			p, err = ns.ProjectSvc.GetProject(userID, note.ProjectID)
			if err != nil {
				return err
			}

			utilities.SetRedisKey(ns.RDB, ctx, fmt.Sprintf("project:%d", note.ProjectID), p, 0)
		} else {
			return err
		}
	}

	// do checks:
	// cannot delete other user's notes
	if p.UserID != userID {
		return customerrors.ErrNoRecord
	}

	// remove the stale note from the cache
	err = utilities.DeleteRedisKey(ns.RDB, ctx, fmt.Sprintf("note:%d", noteID))
	if err != nil {
		return err
	}

	// delete the note
	return ns.Repo.delete(noteID, p.ID)
}

func (ns *NoteService) updateNoteTitleContent(
	v *validator.Validator, userID, noteID int, title, content *string,
) error {
	var cleanedTitle, cleanedContent *string

	if title == nil && content == nil {
		v.AddError("input", "at least one field (title or content) must be provided")
		return validator.ErrFailedValidation
	}

	if title != nil {
		trimmed := strings.TrimSpace(*title)
		cleanedTitle = &trimmed
		validateTitle(v, *cleanedTitle)
	}

	if content != nil {
		trimmed := strings.TrimSpace(*content)
		cleanedContent = &trimmed
		validateContent(v, *cleanedContent)
	}

	if !v.IsValid() {
		return validator.ErrFailedValidation
	}

	// remove the stale note from the cache
	ctx := context.Background()
	err := utilities.DeleteRedisKey(ns.RDB, ctx, fmt.Sprintf("note:%d", noteID))
	if err != nil {
		return err
	}

	return ns.Repo.updateNoteTitleContent(userID, noteID, cleanedTitle, cleanedContent)
}

func (ns *NoteService) updateNoteVisibility(
	v *validator.Validator, userID, noteID int, visibility string,
) error {
	cleanedVisibility := strings.TrimSpace(visibility)
	if validateVisibility(v, cleanedVisibility); !v.IsValid() {
		return validator.ErrFailedValidation
	}
	// Fetch note and project to validate visibility constraints

	note := &Note{}
	p := &project.Project{}
	ctx := context.Background()

	err := utilities.GetUnmarshalRedisKey(ns.RDB, ctx, fmt.Sprintf("note:%d", noteID), note)
	if err != nil {
		if err == redis.Nil {
			note, err = ns.Repo.get(noteID)
			if err != nil {
				return err
			}

			utilities.SetRedisKey(ns.RDB, ctx, fmt.Sprintf("note:%d", noteID), note, 0)
		} else {
			return err
		}
	}

	// fetch the project
	err = utilities.GetUnmarshalRedisKey(ns.RDB, ctx, fmt.Sprintf("project:%d", note.ProjectID), p)
	if err != nil {
		if err == redis.Nil {
			p, err = ns.ProjectSvc.GetProject(userID, note.ProjectID)
			if err != nil {
				return err
			}

			utilities.SetRedisKey(ns.RDB, ctx, fmt.Sprintf("project:%d", note.ProjectID), p, 0)
		} else {
			return err
		}
	}

	// Ensure note can't be more public than its project
	if p.Visibility == "private" && cleanedVisibility == "public" {
		v.AddError("visibility", "cannot be more public than the project")
		return validator.ErrFailedValidation
	}

	// cannot update other's notes
	if p.UserID != userID {
		return customerrors.ErrNoRecord
	}

	// remove the stale note from the cache
	err = utilities.DeleteRedisKey(ns.RDB, ctx, fmt.Sprintf("note:%d", noteID))
	if err != nil {
		return err
	}

	return ns.Repo.updateNoteVisibility(userID, noteID, cleanedVisibility)
}

func (ns *NoteService) updateNoteColor(
	v *validator.Validator, userID, noteID int, color string,
) error {
	cleanedColor := strings.TrimSpace(color)
	if validateColor(v, cleanedColor); !v.IsValid() {
		return validator.ErrFailedValidation
	}

	// remove the stale note from the cache
	ctx := context.Background()
	err := utilities.DeleteRedisKey(ns.RDB, ctx, fmt.Sprintf("note:%d", noteID))
	if err != nil {
		return err
	}

	return ns.Repo.updateNoteColor(userID, noteID, cleanedColor)
}

func (ns *NoteService) getNoteOwner(userID, noteID int) (string, error) {
	return ns.Repo.getOwner(userID, noteID)
}
