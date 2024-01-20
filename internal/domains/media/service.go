package media

import (
	"context"
	"fmt"
	repo "github.com/vynious/go-travel/internal/db"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"time"
)

type MediaService struct {
	repository *repo.Repository
	timeout    time.Duration
}

func NewMediaService(repo *repo.Repository) *MediaService {
	return &MediaService{
		repository: repo,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *MediaService) CreateNewMedia(ctx context.Context, eid int64, url string) (db.Medium, error) {
	params := db.CreateMediaParams{
		EntryID: eid,
		Url:     url,
	}
	media, err := s.repository.Queries.CreateMedia(ctx, params)
	if err != nil {
		return db.Medium{}, fmt.Errorf("[s] failed to create media: %w", err)
	}
	return media, nil
}

func (s *MediaService) GetMediaById(ctx context.Context, mid int64) (db.Medium, error) {
	media, err := s.repository.Queries.GetMediaById(ctx, mid)
	if err != nil {
		return db.Medium{}, fmt.Errorf("[s] failed to get media :%w", err)
	}
	return media, nil
}

func (s *MediaService) GetMediasByEntryId(ctx context.Context, eid int64) ([]db.Medium, error) {
	medias, err := s.repository.Queries.GetAllMediaByEntryId(ctx, eid)
	if err != nil {
		return []db.Medium{}, fmt.Errorf("[s] failed to get medias :%w", err)
	}
	return medias, nil
}

func (s *MediaService) UpdateMediaById(ctx context.Context, mid int64, url string) (db.Medium, error) {
	params := db.UpdateMediaByIdParams{
		ID:  mid,
		Url: url,
	}
	media, err := s.repository.Queries.UpdateMediaById(ctx, params)
	if err != nil {
		return db.Medium{}, fmt.Errorf("[s] failed to update media :%w", err)
	}
	return media, nil
}

func (s *MediaService) DeleteMediaById(ctx context.Context, mid int64) (db.Medium, error) {
	media, err := s.repository.Queries.DeleteMediaById(ctx, mid)
	if err != nil {
		return db.Medium{}, fmt.Errorf("[s] failed to delete media :%w", err)
	}
	return media, nil
}
