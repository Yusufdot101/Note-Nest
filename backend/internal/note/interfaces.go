package note

import (
	"time"

	"github.com/Yusufdot101/note-nest/internal/project"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/redis/go-redis/v9"
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
}

type NoteService struct {
	Repo       *Repository
	ProjectSvc *project.ProjectService
	RDB        *redis.Client
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
