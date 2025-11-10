package note

import (
	"time"

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
}

type MockRepo struct {
	insertCalled bool
	getCalled    bool
	deleteCalled bool
}

func (mr *MockRepo) insert(n *Note) error {
	mr.insertCalled = true
	return nil
}

func (mr *MockRepo) get(projectID, noteID int, visibility string) ([]*Note, error) {
	mr.getCalled = true
	return nil, nil
}

func (mr *MockRepo) delete(noteID int) error {
	mr.deleteCalled = true
	return nil
}

type Repo interface {
	insert(n *Note) error
	get(projectID, noteID int, visibility string) ([]*Note, error)
	delete(noteID int) error
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
	v.CheckAddError(n.Title != "", "name", "must be given")
	v.CheckAddError(n.Visibility != "", "visibility", "must be given")
	allowedVisibility := []string{"private", "public"}
	v.CheckAddError(validator.ValueInList(n.Visibility, allowedVisibility...), "visibility", "not allowed")
	validateColor(v, n.Color)
}

func validateColor(v *validator.Validator, color string) {
	v.CheckAddError(color != "", "color", "must be provided")
	v.CheckAddError(len(color) == 7 && color[0] == '#', "color", "must be a valid hex color (e.g., #ffffff)")
	// Additional regex check if needed
}
