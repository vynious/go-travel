package media

import (
	"github.com/vynious/go-travel/internal/domains/s3"
	"net/http"
)

type MediaHandler struct {
	*MediaService
	s3Client s3.S3Client
}

func (h *MediaHandler) UploadNewMedia(w http.ResponseWriter, r *http.Request) {
	h.s3Client.UploadMediaToBucket()
}

func (h *MediaHandler) ViewMedia(w http.ResponseWriter, r *http.Request) {
	h.s3Client.GetMediaFromBucket()
}

func (h *MediaHandler) ViewMediaByEntry(w http.ResponseWriter, r *http.Request) {
	// loop
	h.s3Client.GetMediaFromBucket()

}

func (h *MediaHandler) UpdateMedia(w http.ResponseWriter, r *http.Request) {
	h.s3Client.UpdateMediaFromBucket()
}

func (h *MediaHandler) DeleteMedia(w http.ResponseWriter, r *http.Request) {
	h.s3Client.DeleteMediaFromBucket()
}
