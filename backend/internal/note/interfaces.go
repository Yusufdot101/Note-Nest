package note

import (
	"time"

	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/project"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

type Note struct {
	ID            int
	ProjectID     int
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	Title         string
	Content       string
	Color         string
	Visibility    string
	LikesCount    uint16
	CommentsCount uint16
	SavesCount    uint16
	SharesCount   uint16
}

type Repo interface {
	insert(n *Note) error
	get(noteID int) (*Note, error)
	getMany(currentUserID, queryUserID, projectID int, title, visibility string, filter *filter.Filter) ([]*Note, error)
	delete(noteID, projectID int) error
	updateNoteTitleContent(userID, noteID int, title, content *string) error
	updateNoteVisibility(userID, noteID int, visibility string) error
	updateNoteColor(userID, noteID int, color string) error
}

type mockRepo struct {
	updateNoteTitleContentCalled bool
	updateNoteColorCalled        bool
}

func (mr *mockRepo) insert(n *Note) error {
	return nil
}

func (mr *mockRepo) get(noteID int) (*Note, error) {
	n := &Note{
		ID: noteID,
	}
	return n, nil
}

func (mr *mockRepo) getMany(currentUserID, queryUserID, projectID int, title, visibility string, filter *filter.Filter) ([]*Note, error) {
	return nil, nil
}

func (mr *mockRepo) delete(noteID, projectID int) error {
	return nil
}

func (mr *mockRepo) updateNoteTitleContent(userID, noteID int, title, content *string) error {
	mr.updateNoteTitleContentCalled = true
	return nil
}

func (mr *mockRepo) updateNoteVisibility(userID, noteID int, visibility string) error {
	return nil
}

func (mr *mockRepo) updateNoteColor(userID, noteID int, color string) error {
	mr.updateNoteColorCalled = true
	return nil
}

type NoteService struct {
	Repo       Repo
	ProjectSvc *project.ProjectService
}

type NoteHandler struct {
	svc *NoteService
}

func newHandler(svc *NoteService) *NoteHandler {
	return &NoteHandler{
		svc: svc,
	}
}

func validateNote(v *validator.Validator, n *Note) {
	validateTitle(v, n.Title)
	validateContent(v, n.Content)
	validateVisibility(v, n.Visibility)
	validateColor(v, n.Color)
}

func validateTitle(v *validator.Validator, title string) {
	v.CheckAddError(title != "", "title", "must be given")
}

func validateContent(v *validator.Validator, content string) {
	v.CheckAddError(content != "", "content", "must be given")
}

func validateVisibility(v *validator.Validator, visibility string) {
	v.CheckAddError(visibility != "", "visibility", "must be given")
	allowedVisibility := []string{"private", "public"}
	v.CheckAddError(validator.ValueInList(visibility, allowedVisibility...), "visibility", "value not allowed")
}

func validateColor(v *validator.Validator, color string) {
	v.CheckAddError(color != "", "color", "must be provided")
	v.CheckAddError(len(color) == 7 && color[0] == '#', "color", "must be a valid hex color (e.g., #ffffff)")
	// Additional regex check if needed
}
