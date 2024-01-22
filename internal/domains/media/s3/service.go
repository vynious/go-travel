package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/vynious/go-travel/internal/domains/media"
	"os"
)

type S3Client struct {
	s3Client *s3.Client
	psClient *s3.PresignClient
}

func NewS3Client() (*S3Client, error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to load s3 config :%w", err)
	}
	client := s3.NewFromConfig(cfg)
	psClient := s3.NewPresignClient(client)
	return &S3Client{
		s3Client: client,
		psClient: psClient,
	}, nil
}

func (client *S3Client) PutSignedMediaToBucket(ctx context.Context, fileInput media.FileInput) (*v4.PresignedHTTPRequest, error) {

	bucketName := os.Getenv("AWS_BUCKET_TRAVEL_ENTRY")
	if bucketName == "" {
		return nil, fmt.Errorf("s3 bucket name not configured")
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileInput.Filename),
	}
	url, err := client.psClient.PresignPutObject(ctx, input, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get signed url: %w", err)
	}
	return url, nil
}

func (client *S3Client) GetSignedMediaFromBucket(ctx context.Context, key string) (*v4.PresignedHTTPRequest, error) {
	bucketName := os.Getenv("AWS_BUCKET_TRAVEL_ENTRY")
	if bucketName == "" {
		return nil, fmt.Errorf("s3 bucket name not configured")
	}

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	url, err := client.psClient.PresignGetObject(ctx, input, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get signed url: %w", err)
	}

	return url, nil
}

func (client *S3Client) DeleteMediaFromBucket() {

}
