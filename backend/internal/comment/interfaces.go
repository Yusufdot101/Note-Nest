/*
Package comment provides logic and handlers for fetching and saving comments
*/
package comment

import (
	"time"

	"github.com/Yusufdot101/note-nest/internal/note"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

type comment struct {
	ID         int
	Edited     bool
	CreatedAt  time.Time
	UserID     int
	NoteID     int
	Content    string
	LikesCount int
}

type repo interface {
	insert(c *comment, projectID int) error
	get(userID, noteID int) ([]*comment, error)
	update(userID, commentID int, content string) error
	delete(userID, commentID int) error
}

type commentService struct {
	repo    repo
	noteSvc *note.NoteService
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
