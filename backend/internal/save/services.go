package save

import "github.com/Yusufdot101/note-nest/internal/note"

func (svc *saveService) saveNote(userID, noteID int) error {
	s := &save{
		userID: userID,
		noteID: noteID,
	}
	return svc.repo.insert(s)
}

func (svc *saveService) noteIsSaved(userID, noteID int) (bool, error) {
	return svc.repo.isSaved(userID, noteID)
}

func (svc *saveService) savedNotes(userID int) ([]*note.Note, error) {
	return svc.repo.getSavedNotes(userID)
}
