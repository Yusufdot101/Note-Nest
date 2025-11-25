package project

import (
	"time"

	"github.com/Yusufdot101/note-nest/internal/customerrors"
	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/validator"
)

type Project struct {
	ID            int
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	UserID        int
	Title         string
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

func (mr *MockRepo) get(currentUserID, userID int, title, visibility string, filter filter.Filter) ([]*Project, error) {
	mr.getCalled = true
	projects := []*Project{}
	for i := range 10 {
		project := &Project{
			UserID: i,
		}
		projects = append(projects, project)
	}
	return projects, nil
}

func (mr *MockRepo) getOne(ID int) (*Project, error) {
	mr.getOneCalled = true
	project := &Project{
		UserID: 1,
	}

	return project, nil
}

func (mr *MockRepo) delete(projectID int) error {
	mr.deleteCalled = true
	return nil
}

func (mr *MockRepo) update(userID, projectID int, name, description, visibility, color *string) error {
	gotPproject := &Project{
		UserID: 1,
	}

	if gotPproject.UserID != userID {
		return customerrors.ErrNoRecord
	}

	mr.updateCalled = true
	return nil
}

type Repo interface {
	insert(p *Project) error
	get(currentUserID, userID int, title, visibility string, filter filter.Filter) ([]*Project, error)
	getOne(ID int) (*Project, error)
	delete(projectID int) error
	update(userID, projectID int, name, description, visibility, color *string) error
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
	validateTitle(v, p.Title)
	validateVisibility(v, p.Visibility)
	validateColor(v, p.Color)
}

func validateTitle(v *validator.Validator, title string) {
	v.CheckAddError(title != "", "title", "must be given")
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
