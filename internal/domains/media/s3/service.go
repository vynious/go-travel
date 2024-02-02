package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"os"
	"strconv"
)

type S3Client struct {
	s3Client *s3.Client
	psClient *s3.PresignClient
}

func NewS3Client() (*S3Client, error) {
	// Load the Shared AWS Configuration (~/.aws/config)

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY"),
			os.Getenv("AWS_SECRET_KEY"),
			"")), // empty session token
	)
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

func (client *S3Client) PutSignedMediaToBucket(ctx context.Context, media *db.Medium) (*v4.PresignedHTTPRequest, error) {

	bucketName := os.Getenv("AWS_BUCKET_TRAVEL_ENTRY")
	if bucketName == "" {
		return nil, fmt.Errorf("s3 bucket name not configured")
	}
	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(generateS3Key(media.EntryID, media.Key)),
		ACL:    types.ObjectCannedACL("public-read"),
	}

	url, err := client.psClient.PresignPutObject(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get signed url: %w", err)
	}
	return url, nil
}

func (client *S3Client) GetSignedMediaFromBucket(ctx context.Context, media *db.Medium) (*v4.PresignedHTTPRequest, error) {
	bucketName := os.Getenv("AWS_BUCKET_TRAVEL_ENTRY")
	if bucketName == "" {
		return nil, fmt.Errorf("s3 bucket name not configured")
	}

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(generateS3Key(media.EntryID, media.Key)),
	}

	url, err := client.psClient.PresignGetObject(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get signed url: %w", err)
	}

	return url, nil
}

func (client *S3Client) DeleteMediaFromBucket(ctx context.Context, media *db.Medium) (*v4.PresignedHTTPRequest, error) {
	bucketName := os.Getenv("AWS_BUCKET_TRAVEL_ENTRY")

	if bucketName == "" {
		return nil, fmt.Errorf("s3 bucket name not configured")
	}
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(generateS3Key(media.EntryID, media.Key)),
	}
	url, err := client.psClient.PresignDeleteObject(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get signed url: %w", err)
	}
	return url, nil
}

func generateS3Key(eid int64, key string) string {
	strEID := strconv.FormatInt(eid, 10)
	res := "travel_entry_" + strEID + "_" + key
	return res
}
