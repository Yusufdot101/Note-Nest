package note

import (
	"strings"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (ns *NoteService) newNote(
	v *validator.Validator, userID, projectID int, title, content, visibility, color string,
) (*Note, error) {
	// fetch the project
	p, err := ns.ProjectSvc.GetProject(userID, projectID)
	if err != nil {
		return nil, err
	}

	// do checks
	cleanedTitle := strings.TrimSpace(title)
	cleanedVisibility := strings.ToLower(strings.TrimSpace(visibility))
	cleanColor := strings.ToLower(strings.TrimSpace(color))
	n := &Note{
		ProjectID:  projectID,
		Title:      cleanedTitle,
		Content:    content,
		Visibility: cleanedVisibility,
		Color:      cleanColor,
	}
	if p.UserID != userID {
		return nil, custom_errors.ErrNoRecord
	}
	if p.Visibility == "private" && n.Visibility == "public" {
		v.AddError("entry", "cannot be more public than the project")
		return nil, validator.ErrFailedValidation
	}
	if validateNote(v, n); !v.IsValid() {
		return nil, validator.ErrFailedValidation
	}

	// save to db or return err
	err = ns.Repo.insert(n)
	if err != nil {
		return nil, err
	}

	return n, nil
}
