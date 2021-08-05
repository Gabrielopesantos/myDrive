package repository

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/minio/minio-go/v7"
)

// Users Minio S3 Compatible Repository
type userMinioRepo struct {
	client *minio.Client
}

// NewUserMinioRepo Creates ?
func NewUserMinioRepo(client *minio.Client) MinioRepository {
	return userMinioRepo{client: client}
}

//func (r *userMinioRepo) PutObject(ctx context.Context, bucket string, input *models.UploadInput) (*minio.UploadInfo, error) {
//
//	//return nil,
//}

//func (r *userMinioRepo) PutObject(ctx context.Context, bucket string, input *models.UploadInput) (*minio.UploadInfo, error) {
//
//	return nil, err
//}

