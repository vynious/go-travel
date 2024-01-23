package media

import (
	"context"
	"fmt"
	repo "github.com/vynious/go-travel/internal/db"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"github.com/vynious/go-travel/internal/domains/media/s3"
	"sync"
	"time"
)

type MediaService struct {
	repository *repo.Repository
	s3Client   *s3.S3Client
	timeout    time.Duration
}

func NewMediaService(repo *repo.Repository, s *s3.S3Client) *MediaService {
	return &MediaService{
		repository: repo,
		s3Client:   s,
		timeout:    time.Duration(2) * time.Second,
	}
}

func (s *MediaService) CreateNewMedia(ctx context.Context, eid int64, key string) (*MediaResponse, error) {
	params := db.CreateMediaParams{
		EntryID: eid,
		Key:     key,
	}
	media, err := s.repository.Queries.CreateMedia(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("[s] failed to create media: %w", err)
	}

	signedUrl, err := s.s3Client.PutSignedMediaToBucket(ctx, &media)
	if err != nil {
		return nil, fmt.Errorf("[s] failed to generate signed url: %w", err)
	}

	return &MediaResponse{
		SignedUrl: signedUrl, // for put
		Media:     media,
	}, nil
}

func (s *MediaService) GetMediaById(ctx context.Context, eid int64, key string) (*MediaResponse, error) {
	params := db.GetMediaByKeyParams{
		Key:     key,
		EntryID: eid,
	}
	media, err := s.repository.Queries.GetMediaByKey(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("[s] failed to get media :%w", err)
	}
	url, err := s.s3Client.GetSignedMediaFromBucket(ctx, &media)
	if err != nil {
		return nil, fmt.Errorf("[s] failed to get signed url :%w", err)
	}
	return &MediaResponse{
		Media:     media,
		SignedUrl: url, // for get
	}, nil
}

func (s *MediaService) GetMediasByEntryId(ctx context.Context, eid int64) ([]*MediaResponse, error) {
	/*
		get all medias under entry_id, use for-loop & concurrency to get all pre-signed urls
	*/
	medias, err := s.repository.Queries.GetAllMediaByEntryId(ctx, eid)
	if err != nil {
		return nil, fmt.Errorf("failed to get all media :%w", err)
	}
	result := make([]*MediaResponse, len(medias))
	errCh := make(chan error, len(medias))
	var wg sync.WaitGroup

	for i, media := range medias {
		wg.Add(1)
		go func(m db.Medium, idx int) {
			defer wg.Done()
			url, err := s.s3Client.GetSignedMediaFromBucket(ctx, &m)
			if err != nil {
				errCh <- fmt.Errorf("failed to  generate signed url: %w", err)
				return
			}
			result[idx] = &MediaResponse{
				Media:     m,
				SignedUrl: url, // for get
			}
		}(media, i)
	}
	go func() {
		wg.Wait()
		close(errCh)
	}()

	for i := 0; i < len(medias); i++ {
		if err := <-errCh; err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *MediaService) DeleteMediaById(ctx context.Context, eid int64, key string) (*MediaResponse, error) {
	params := db.DeleteMediaByKeyParams{
		EntryID: eid,
		Key:     key,
	}

	media, err := s.repository.Queries.DeleteMediaByKey(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("[s] failed to delete media :%w", err)
	}

	url, err := s.s3Client.DeleteMediaFromBucket(ctx, &media)
	if err != nil {
		return nil, fmt.Errorf("[s] failed to get signed url :%w", err)
	}
	return &MediaResponse{
		Media:     media,
		SignedUrl: url, //  for delete
	}, nil
}

func (s *MediaService) DeleteMediaByEntryId(ctx context.Context, eid int64) (*MediaResponse, error) {

}
