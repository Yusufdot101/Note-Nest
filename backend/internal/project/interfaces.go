package project

import (
	"time"

	"github.com/Yusufdot101/note-nest/internal/validator"
)

type Project struct {
	ID          int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      int
	Name        string
	Description string
	Visibility  string
}

type MockRepo struct {
	insertCalled bool
}

func (mr *MockRepo) insert(p *Project) error {
	mr.insertCalled = true
	return nil
}

type Repo interface {
	insert(p *Project) error
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
	v.CheckAddError(p.Visibility != "", "visibility", "must be given")
	allowedVisibility := []string{"private", "public"}
	v.CheckAddError(validator.ValueInList(p.Visibility, allowedVisibility...), "visibility", "not allowed")
}
