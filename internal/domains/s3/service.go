package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"os"
)

type S3Client struct {
	s3Client *s3.Client
}

func NewS3Client() (*S3Client, error) {
	// Load the Shared AWS Configuration (~/.aws/config) ??
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to load s3 config :%w", err)
	}
	client := s3.NewFromConfig(cfg)
	return &S3Client{
		s3Client: client,
	}, nil
}

func (client *S3Client) UploadMediaToBucket(ctx context.Context, fileInput FileInput) error {

	bucketName := os.Getenv("AWS_BUCKET_TRAVEL_ENTRY")
	if bucketName == "" {
		return fmt.Errorf("s3 bucket name not configured")
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileInput.Filename),
		Body:   fileInput.File,
	}
	_, err := client.s3Client.PutObject(ctx, input)
	if err != nil {
		return fmt.Errorf("[s] failed to put object in s3: %w", err)
	}
	return nil
}

func (client *S3Client) GetMediaFromBucket(ctx context.Context, key string) error {
	bucketName := os.Getenv("AWS_BUCKET_TRAVEL_ENTRY")
	if bucketName == "" {
		return fmt.Errorf("s3 bucket name not configured")
	}

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	response, err := client.s3Client.GetObject(ctx, input)
	if err != nil {
		return fmt.Errorf("[s] failed to get object in s3: %w", err)
	}
}

func (client *S3Client) UpdateMediaFromBucket() {

}

func (client *S3Client) DeleteMediaFromBucket() {

}
