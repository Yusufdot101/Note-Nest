/*
Package comment provides logic and handlers for fetching and saving comments
*/
package comment

import (
	"time"

	"github.com/Yusufdot101/note-nest/internal/note"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/redis/go-redis/v9"
)

type comment struct {
	ID         int
	Edited     bool
	CreatedAt  time.Time
	UserID     int
	Username   string
	NoteID     int
	Content    string
	LikesCount int
}

type repo interface {
	insert(c *comment, projectID int) error
	get(userID, noteID int) ([]*comment, []int, error)
	update(userID, commentID int, content string) error
	delete(userID, commentID int) error
	getByIDs(ids []int) ([]*comment, error)
}

type mockRepo struct {
	getCalled    bool
	updateCalled bool
	deleteCalled bool
}

func (mr *mockRepo) insert(c *comment, projectID int) error {
	return nil
}

func (mr *mockRepo) get(userID, noteID int) ([]*comment, []int, error) {
	mr.getCalled = true
	c := &comment{
		UserID: userID,
		NoteID: noteID,
	}
	return []*comment{c}, []int{1}, nil
}

func (mr *mockRepo) update(userID, commentID int, content string) error {
	mr.updateCalled = true
	return nil
}

func (mr *mockRepo) delete(userID, commentID int) error {
	mr.deleteCalled = true
	return nil
}

func (mr *mockRepo) getByIDs(ids []int) ([]*comment, error) {
	return nil, nil
}

type commentService struct {
	repo    repo
	noteSvc *note.NoteService
	RDB     *redis.Client
}

type commentHandler struct {
	svc *commentService
}

func newHandler(svc *commentService) *commentHandler {
	return &commentHandler{
		svc: svc,
	}
}

func validateComment(v *validator.Validator, c *comment) {
	v.CheckAddError(c.Content != "", "content", "must be given")
}
