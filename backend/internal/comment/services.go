package comment

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/Yusufdot101/note-nest/internal/validator"
	"github.com/redis/go-redis/v9"
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

type commentQueryKey struct {
	UserID int
	NoteID int
}

func (k commentQueryKey) RedisKey() string {
	b, _ := json.Marshal(k)
	sum := sha1.Sum(b)
	return "comments:list:" + hex.EncodeToString(sum[:])
}

func (cs *commentService) getComments(userID, noteID int) ([]*comment, error) {
	qk := commentQueryKey{
		UserID: userID,
		NoteID: noteID,
	}

	var data struct {
		IDs []int `json:"ids"`
	}
	key := qk.RedisKey()
	ctx := context.Background()
	err := utilities.GetUnmarshalRedisKey(cs.RDB, ctx, key, &data)
	if err == nil {
		comments, err := cs.repo.getByIDs(data.IDs)
		if err != nil {
			return nil, err
		}
		return comments, nil
	}

	if err != redis.Nil {
		return nil, err
	}

	comments, ids, err := cs.repo.get(userID, noteID)
	if err != nil {
		return nil, err
	}

	data.IDs = ids
	utilities.SetRedisKey(cs.RDB, ctx, key, data, 60*time.Second)
	return comments, nil
}

func (svc *commentService) updateComment(v *validator.Validator, userID, commentID int, content string) error {
	c := &comment{
		Content: content,
	}
	if validateComment(v, c); !v.IsValid() {
		return validator.ErrFailedValidation
	}
	return svc.repo.update(userID, commentID, content)
}

func (svc *commentService) deleteComment(userID, commentID int) error {
	return svc.repo.delete(userID, commentID)
}
