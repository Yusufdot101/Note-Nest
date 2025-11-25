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

type mockRepo struct {
	insertCalled        bool
	isSavedCalled       bool
	getSavedNotesCalled bool
	deleteCalled        bool
}

func (mr *mockRepo) insert(s *save) error {
	mr.insertCalled = true
	return nil
}

func (mr *mockRepo) isSaved(userID, noteID int) (bool, error) {
	mr.isSavedCalled = true
	return true, nil
}

func (mr *mockRepo) getSavedNotes(userID int) ([]*note.Note, error) {
	mr.getSavedNotesCalled = true
	return nil, nil
}

func (mr *mockRepo) delete(userID, noteID int) error {
	mr.deleteCalled = true
	return nil
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
