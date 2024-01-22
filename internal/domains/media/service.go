package media

import (
	"context"
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
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

func (s *MediaService) CreateNewMedia(ctx context.Context, eid int64, fileData FileInput) (*v4.PresignedHTTPRequest, error) {
	url, err := s.s3Client.PutSignedMediaToBucket(ctx, fileData)
	if err != nil {
		return nil, fmt.Errorf("failed to uploaded media to s3 :%w", err)
	}

	params := db.CreateMediaParams{
		EntryID: eid,
		Url:     fileData.Filename,
	}
	_, err = s.repository.Queries.CreateMedia(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("[s] failed to create media: %w", err)
	}
	return url, nil
}

func (s *MediaService) GetMediaById(ctx context.Context, mid int64) (db.Medium, error) {
	media, err := s.repository.Queries.GetMediaById(ctx, mid)
	if err != nil {
		return db.Medium{}, fmt.Errorf("[s] failed to get media :%w", err)
	}
	return media, nil
}

func (s *MediaService) GetMediasByEntryId(ctx context.Context, eid int64) ([]*v4.PresignedHTTPRequest, error) {
	medias, err := s.repository.Queries.GetAllMediaByEntryId(ctx, eid)
	if err != nil {
		return nil, fmt.Errorf("failed to get all media :%w", err)
	}
	presignedUrls := make([]*v4.PresignedHTTPRequest, len(medias))
	errCh := make(chan error, len(medias))

	var wg sync.WaitGroup
	// for loop for getting signed urls
	for i, media := range medias {
		// goroutine to get signed urls, if fail pass into errCh if success assigned into array based on idx
		wg.Add(1)
		go func(m db.Medium, idx int) {
			defer wg.Done()
			url, err := s.s3Client.GetSignedMediaFromBucket(ctx, m.Url)
			if err != nil {
				errCh <- fmt.Errorf("failed to  generate signed url: %w", err)
				return
			}
			presignedUrls[idx] = url
			errCh <- nil
		}(media, i)
	}

	// check err channel for any errors from getting signed urls
	go func() {
		wg.Wait()
		close(errCh)
	}()

	for i := 0; i < len(medias); i++ {
		if err := <-errCh; err != nil {
			return nil, err
		}
	}

	return presignedUrls, nil
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
