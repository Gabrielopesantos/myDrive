package models

import "io"

// MinIO upload file
type UploadInput struct {
	File        io.Reader `json:"file_content,omitempty"`
	Name        string    `json:"filename,omitempty" db:"filename"`
	Size        int64     `json:"size,omitempty" db:"size"`
	ContentType string    `json:"extension,omitempty" db:"extension"`
	BucketName  string    `json:"bucket,omitempty"`
}
