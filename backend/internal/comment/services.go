package comment

import (
	"github.com/Yusufdot101/note-nest/internal/validator"
)

func (svc *commentService) newComment(v *validator.Validator, userID, noteID int, content string) error {
	c := &comment{
		UserID:  userID,
		NoteID:  noteID,
		Content: content,
	}
	if validateComment(v, c); !v.IsValid() {
		return validator.ErrFailedValidation
	}

	// fetch note
	_, err := svc.noteSvc.GetNote(userID, noteID)
	if err != nil {
		return err
	}
	// since GetNote didn't returned an error(custom_errors.ErrNoRecord), it means the note is either public or belongs to the current user, so the user can comment on it
	// insert comment
	return svc.repo.insert(c)
}
