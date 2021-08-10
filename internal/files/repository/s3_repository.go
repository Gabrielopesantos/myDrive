package repository

import (
	"context"
	"fmt"
	"github.com/gabrielopesantos/myDrive-api/internal/files"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/minio/minio-go/v7"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

// Files Minio S3 Compatible Repository
type fileMinioRepo struct {
	client *minio.Client
}

// NewFileMinioRepo Creates ?
func NewFileMinioRepo(client *minio.Client) files.MinioRepository {
	return &fileMinioRepo{client: client}
}

// PutObject Upload file to Minio Server
func (r *fileMinioRepo) PutObject(ctx context.Context, file models.File) (*minio.UploadInfo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userMinioRepo.PutObject")
	defer span.Finish()

	options := minio.PutObjectOptions{
		ContentType:  file.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	uploadInfo, err := r.client.PutObject(ctx, file.BucketName, r.generateFilename(&file), file.File, file.Size, options)
	if err != nil {
		return &uploadInfo, errors.Wrap(err, "fileMinioRepo.PutObject.PutObject")
	}

	return &uploadInfo, nil
}

//// GetObject Download file to Minio Server
//func (r *userMinioRepo) GetObject(ctx context.Context, bucket string, fileName string) (*minio.Object, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "userMinioRepo.PutObject")
//	defer span.Finish()
//
//	object, err := r.client.GetObject(ctx, bucket, fileName, minio.GetObjectOptions{})
//	if err != nil {
//		return nil, errors.Wrap(err, "userMinioRepo.GetObject.GetObject")
//	}
//
//	return object, nil
//}
//
//// RemoveObject Remove object from Minio Server
//func (r *userMinioRepo) RemoveObject(ctx context.Context, bucket string, fileName string) error {
//
//	err := r.client.RemoveObject(ctx, bucket, fileName, minio.RemoveObjectOptions{})
//	if err != nil {
//		return errors.Wrap(err, "userMinioRepo.RemoveObject.RemoveObject")
//	}
//
//	return nil
//}
//
func (r *fileMinioRepo) generateFilename(file *models.File) string {
	return fmt.Sprintf("%s-%s-%s", file.FileOwnerId, file.FileId, file.Name)
}
