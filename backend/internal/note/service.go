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

	// cannot create in other's projects
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

func (ns *NoteService) GetNote(userID, noteID int) (*Note, error) {
	// fetch the note
	note, err := ns.Repo.get(noteID)
	if err != nil {
		return nil, err
	}

	// fetch the project
	project, err := ns.ProjectSvc.GetProject(userID, note.ProjectID)
	if err != nil {
		return nil, err
	}

	// do checks:
	// cannot see other's private projects/notes
	if project.UserID != userID && (project.Visibility == "private" || note.Visibility == "private") {
		return nil, custom_errors.ErrNoRecord
	}

	// fetch and return notes
	return note, nil
}

func (ns *NoteService) getNotes(currentUserID, queryUserID, projectID *int, visibility string) ([]*Note, error) {
	notes, err := ns.Repo.getMany(currentUserID, queryUserID, projectID, visibility)
	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (ns *NoteService) deleteNote(userID, noteID int) error {
	note, err := ns.Repo.get(noteID)
	if err != nil {
		return err
	}

	// fetch the project
	project, err := ns.ProjectSvc.GetProject(userID, note.ProjectID)
	if err != nil {
		return err
	}

	// do checks:
	// cannot delete other user's notes
	if project.UserID != userID {
		return custom_errors.ErrNoRecord
	}

	// delete the note
	return ns.Repo.delete(noteID, project.ID)
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

	return ns.Repo.updateNoteTitleContent(userID, noteID, cleanedTitle, cleanedContent)
}
