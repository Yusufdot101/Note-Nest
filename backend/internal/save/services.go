package save

import (
	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/note"
)

func (svc *saveService) saveNote(userID, noteID int) error {
	s := &save{
		userID: userID,
		noteID: noteID,
	}
	return svc.repo.insert(s)
}

func (svc *saveService) unsaveNote(userID, noteID int) error {
	return svc.repo.delete(userID, noteID)
}

func (svc *saveService) noteIsSaved(userID, noteID int) (bool, error) {
	return svc.repo.isSaved(userID, noteID)
}

func (svc *saveService) savedNotes(currentUserID, queryUserID, projectID int, title, visibility string, filter *filter.Filter) ([]*note.Note, *filter.Metadata, error) {
	return svc.repo.getSavedNotes(currentUserID, queryUserID, projectID, title, visibility, filter)
}
