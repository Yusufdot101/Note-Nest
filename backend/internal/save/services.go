package save

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/Yusufdot101/note-nest/internal/filter"
	"github.com/Yusufdot101/note-nest/internal/note"
	"github.com/Yusufdot101/note-nest/internal/utilities"
	"github.com/redis/go-redis/v9"
)

func (svc *saveService) saveNote(userID, noteID int) error {
	s := &save{
		userID: userID,
		noteID: noteID,
	}
	return svc.repo.insert(s)
}

func (svc *saveService) unsaveNote(userID, noteID int) error {
	return svc.repo.delete(userID, noteID)
}

func (svc *saveService) noteIsSaved(userID, noteID int) (bool, error) {
	return svc.repo.isSaved(userID, noteID)
}

type saveQueryKey struct {
	CurrentUserID int
	QueryUserID   int
	ProjectID     int
	Title         string
	Visibility    string
	SortCol       string
	SortDir       string
	Limit         int
	Offset        int
}

func (k saveQueryKey) RedisKey() string {
	b, _ := json.Marshal(k)
	sum := sha1.Sum(b)
	return "notes:list:" + hex.EncodeToString(sum[:])
}

func (ss *saveService) savedNotes(currentUserID, queryUserID, projectID int, title, visibility string, f *filter.Filter) ([]*note.Note, *filter.Metadata, error) {
	qk := saveQueryKey{
		CurrentUserID: currentUserID,
		QueryUserID:   queryUserID,
		ProjectID:     projectID,
		Title:         title,
		Visibility:    visibility,
		SortCol:       f.SortColumn(),
		SortDir:       f.SortDirection(),
		Limit:         f.Limit(),
		Offset:        f.Offset(),
	}

	var data struct {
		IDs   []int `json:"ids"`
		Total int   `json:"total"`
	}
	key := qk.RedisKey()
	ctx := context.Background()
	err := utilities.GetUnmarshalRedisKey(ss.RDB, ctx, key, &data)
	if err == nil {
		notes, err := ss.repo.getByIDs(data.IDs)
		if err != nil {
			return nil, nil, err
		}
		metadata := filter.GenerateMetadata(f.Page, f.PageSize, data.Total)
		return notes, metadata, nil
	}

	if err != redis.Nil {
		return nil, nil, err
	}

	notes, metadata, ids, err := ss.repo.getSavedNotes(currentUserID, queryUserID, projectID, title, visibility, f)
	if err != nil {
		return nil, nil, err
	}

	data.IDs = ids
	data.Total = metadata.TotalResources
	utilities.SetRedisKey(ss.RDB, ctx, key, data, 60*time.Second)
	return notes, metadata, nil
}
