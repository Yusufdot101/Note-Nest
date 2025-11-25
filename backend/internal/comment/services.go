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
	n, err := svc.noteSvc.GetNote(userID, noteID)
	if err != nil {
		return err
	}
	// since GetNote didn't return an error(customerrors.ErrNoRecord), it means the note is either public or belongs to the current user, so the user can comment on it
	// insert comment
	return svc.repo.insert(c, n.ProjectID)
}

func (svc *commentService) getComments(userID, noteID int) ([]*comment, error) {
	return svc.repo.get(userID, noteID)
}

func (svc *commentService) updateComment(userID, commentID int, content string) error {
	return svc.repo.update(userID, commentID, content)
}

func (svc *commentService) deleteComment(userID, commentID int) error {
	return svc.repo.delete(userID, commentID)
}
