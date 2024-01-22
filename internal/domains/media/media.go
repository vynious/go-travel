package media

import "mime/multipart"

type FileInput struct {
	File     multipart.File
	Filename string
}
