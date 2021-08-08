package utils

import (
	"github.com/pkg/errors"
	"mime/multipart"
)

var allowedImageContentTypes = map[string]string{
	"image/png":                "png",
	"image/jpeg":               "jpeg",
	"image/jpg":                "jpg",
}

func CheckImageContentType(fileContent *multipart.FileHeader) error {
	contentType := fileContent.Header.Get("content-type")

	_, ok := allowedImageContentTypes[contentType]
	if !ok {
		return errors.New("file content type not allowed")
	}

	return  nil
}
