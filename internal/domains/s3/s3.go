package s3

import "mime/multipart"

type FileInput struct {
	File     multipart.File
	Filename string
}
