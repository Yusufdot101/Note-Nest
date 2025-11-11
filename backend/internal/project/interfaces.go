package project

import (
	"time"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

type Project struct {
	ID            int
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	UserID        int
	Name          string
	Description   string
	Visibility    string
	Color         string
	EntriesCount  int
	LikesCount    int
	CommentsCount int
}

type MockRepo struct {
	insertCalled bool
	getCalled    bool
	getOneCalled bool
	deleteCalled bool
	updateCalled bool
}

func (mr *MockRepo) insert(p *Project) error {
	mr.insertCalled = true
	return nil
}

func (mr *MockRepo) get(userID int, visibility string) ([]*Project, error) {
	mr.getCalled = true
	return nil, nil
}

func (mr *MockRepo) getOne(ID int) (*Project, error) {
	mr.getOneCalled = true
	return nil, nil
}

func (mr *MockRepo) delete(projectID int) error {
	mr.deleteCalled = true
	return nil
}

func (mr *MockRepo) update(userID, projectID int, name, description, visibility, color string) error {
	mr.updateCalled = true
	return nil
}

type Repo interface {
	insert(p *Project) error
	get(userID int, visibility string) ([]*Project, error)
	getOne(ID int) (*Project, error)
	delete(projectID int) error
	update(userID, projectID int, name, description, visibility, color string) error
}

type ProjectService struct {
	Repo Repo
}

type ProjectHandler struct {
	svc *ProjectService
}

func NewHandler(svc *ProjectService) *ProjectHandler {
	return &ProjectHandler{
		svc: svc,
	}
}

func validateProject(v *validator.Validator, p *Project) {
	v.CheckAddError(p.Name != "", "name", "must be given")
	validateVisibility(v, p.Visibility)
	validateColor(v, p.Color)
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
