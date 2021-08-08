package repository

import (
	"context"
	"fmt"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/gabrielopesantos/myDrive-api/internal/user"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

// Users Minio S3 Compatible Repository
type userMinioRepo struct {
	client *minio.Client
}

// NewUserMinioRepo Creates ?
func NewUserMinioRepo(client *minio.Client) user.MinioRepository {
	return &userMinioRepo{client: client}
}

// PutObject Upload file to Minio Server
func (r *userMinioRepo) PutObject(ctx context.Context, input models.UploadInput) (*minio.UploadInfo, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userMinioRepo.PutObject")
	defer span.Finish()

	options := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	uploadInfo, err := r.client.PutObject(ctx, input.BucketName, r.GenerateFilename(input.Name), input.File, input.Size, options)
	if err != nil {
		return &uploadInfo, errors.Wrap(err, "userMinioRepo.PutObject.PutObject")
	}

	return &uploadInfo, nil
}

// GetObject Download file to Minio Server
func (r *userMinioRepo) GetObject(ctx context.Context, bucket string, fileName string) (*minio.Object, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userMinioRepo.PutObject")
	defer span.Finish()

	object, err := r.client.GetObject(ctx, bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "userMinioRepo.GetObject.GetObject")
	}

	return object, nil
}

// RemoveObject Remove object from Minio Server
func (r *userMinioRepo) RemoveObject(ctx context.Context, bucket string, fileName string) error {

	err := r.client.RemoveObject(ctx, bucket, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return errors.Wrap(err, "userMinioRepo.RemoveObject.RemoveObject")
	}

	return nil
}

// GenerateFilename Generate a unique filename
func (r *userMinioRepo) GenerateFilename(filename string) string {
	prefixUuid := uuid.New().String()
	return fmt.Sprintf("%s-%s", prefixUuid, filename)
}
