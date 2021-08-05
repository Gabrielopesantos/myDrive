package models

import "io"

// MinIO upload file
type UploadInput struct {
	File io.Reader
	Name string
	Size int64
	ContentType string
	BucketName string
}
