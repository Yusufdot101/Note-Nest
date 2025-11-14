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
	// insertCalled bool
	// getCalled    bool
	// deleteCalled bool
}

// func (mr *MockRepo) insert(n *Note) error {
// 	mr.insertCalled = true
// 	return nil
// }
//
// func (mr *MockRepo) get(noteID int) (*Note, error) {
// 	mr.getCalled = true
// 	return nil, nil
// }
//
// func (mr *MockRepo) getMany(currentUserID, queryUserID, projectID *int, visibility string) (*Note, error) {
// 	mr.getCalled = true
// 	return nil, nil
// }
//
// func (mr *MockRepo) delete(noteID int) error {
// 	mr.deleteCalled = true
// 	return nil
// }

type Repo interface {
	insert(n *Note) error
	get(noteID int) (*Note, error)
	getMany(currentUserID, queryUserID, projectID *int, visibility string) ([]*Note, error)
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
	validateTitle(v, n.Title)
	validateVisibility(v, n.Visibility)
	validateColor(v, n.Color)
}

func validateTitle(v *validator.Validator, name string) {
	v.CheckAddError(name != "", "name", "must be given")
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
