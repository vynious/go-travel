package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type S3Client struct {
	s3Client *s3.S3
}

func NewS3Client() (*S3Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"), // e.g., "us-west-1"
	})
	if err != nil {
		log.Fatalf("Error creating AWS session: %v", err)
	}

	// Create S3 service client
	svc := s3.New(sess)

	return &S3Client{
		s3Client: svc,
	}, nil
}

func (client *S3Client) UploadMediaToBucket() {
	//input := &s3.PutObjectInput{}
	//output, err := client.s3Client.PutObject(input)
	//if err != nil {
	//
	//}
}

func (client *S3Client) GetMediaFromBucket() {
	//input := &s3.GetObjectInput{
	//	Bucket: &bucketName,
	//	Key:    &key,
	//}
	//output, err := client.s3Client.GetObject(input)
	//if err != nil {
	//
	//}
}

func (client *S3Client) UpdateMediaFromBucket() {

}

func (client *S3Client) DeleteMediaFromBucket() {

}