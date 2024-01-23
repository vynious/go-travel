package media

import (
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	db "github.com/vynious/go-travel/internal/db/sqlc"
	"mime/multipart"
)

type FileInput struct {
	File     multipart.File
	Filename string
}

type MediaResponse struct {
	SignedUrl *v4.PresignedHTTPRequest
	Media     db.Medium
}
