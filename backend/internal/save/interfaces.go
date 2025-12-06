package save

import (
	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/note"
	"github.com/redis/go-redis/v9"
)

type save struct {
	userID, noteID int
}

type repo interface {
	insert(s *save) error
	isSaved(userID, noteID int) (bool, error)
	getSavedNotes(currentUserID, queryUserID, projectID int, title, visibility string, f *filter.Filter) ([]*note.Note, *filter.Metadata, []int, error)
	delete(userID, noteID int) error
	getByIDs(ids []int) ([]*note.Note, error)
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

func (mr *mockRepo) getSavedNotes(currentUserID, queryUserID, projectID int, title, visibility string, f *filter.Filter) ([]*note.Note, *filter.Metadata, []int, error) {
	mr.getSavedNotesCalled = true
	return nil, &filter.Metadata{}, nil, nil
}

func (mr *mockRepo) delete(userID, noteID int) error {
	mr.deleteCalled = true
	return nil
}

func (mr *mockRepo) getByIDs(ids []int) ([]*note.Note, error) {
	return nil, nil
}

type saveService struct {
	repo repo
	RDB  *redis.Client
}

type saveHandler struct {
	svc *saveService
}

func newHandler(svc *saveService) *saveHandler {
	return &saveHandler{
		svc: svc,
	}
}
