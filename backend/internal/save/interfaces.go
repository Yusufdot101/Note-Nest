package save

import "github.com/Yusufdot101/note-nest/internal/note"

type save struct {
	userID, noteID int
}

type repo interface {
	insert(s *save) error
	isSaved(userID, noteID int) (bool, error)
	getSavedNotes(userID int) ([]*note.Note, error)
	delete(userID, noteID int) error
}

type saveService struct {
	repo repo
}

type saveHandler struct {
	svc *saveService
}

func newHandler(svc *saveService) *saveHandler {
	return &saveHandler{
		svc: svc,
	}
}
