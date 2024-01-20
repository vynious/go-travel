package media

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

type MediaHandler struct {
	*MediaService
	s3Client *s3.S3
}

func (h *MediaHandler) UploadNewMedia(w http.ResponseWriter, r *http.Request) {

}

func (h *MediaHandler) ViewMediaBy(w http.ResponseWriter, r *http.Request) {

}

func (h *MediaHandler) ViewMediaByEntry(w http.ResponseWriter, r *http.Request) {

}

func (h *MediaHandler) UpdateMedia(w http.ResponseWriter, r *http.Request) {

}

func (h *MediaHandler) DeleteMedia(w http.ResponseWriter, r *http.Request) {

}
