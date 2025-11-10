package note

import (
	"strings"

	"github.com/Yusufdot101/note-nest/internal/custom_errors"
	"github.com/Yusufdot101/note-nest/internal/validator"
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
	cleanedVisibility := strings.ToLower(strings.TrimSpace(visibility))
	cleanColor := strings.ToLower(strings.TrimSpace(color))
	n := &Note{
		ProjectID:  projectID,
		Title:      cleanedTitle,
		Content:    content,
		Visibility: cleanedVisibility,
		Color:      cleanColor,
	}
	// cannot create in others projects
	if p.UserID != userID {
		return custom_errors.ErrNoRecord
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

func (ns *NoteService) getNotes(userID, projectID, noteID int, visibility string) ([]*Note, error) {
	// fetch the project
	p, err := ns.ProjectSvc.GetProject(userID, projectID)
	if err != nil {
		return nil, err
	}

	// do checks:
	// cannot access other user's private projects/notes
	if p.UserID != userID {
		if p.Visibility == "private" || visibility == "private" {
			return nil, custom_errors.ErrNoRecord
		}
		// set the visibility to public incase its not given and its "" which would fetch all notse(public + private)
		visibility = "public"
	}

	// fetch and return notes
	return ns.Repo.get(projectID, noteID, visibility)
}

func (ns *NoteService) deleteNote(userID, projectID, noteID int) error {
	// fetch the project
	p, err := ns.ProjectSvc.GetProject(userID, projectID)
	if err != nil {
		return err
	}

	// do checks:
	// cannot delete other user's notes
	if p.UserID != userID {
		return custom_errors.ErrNoRecord
	}

	// delete the note
	return ns.Repo.delete(noteID)
}
