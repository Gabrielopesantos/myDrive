//go:generate mockgen -source minio_repository.go -destination mock/minio_repository_mock.go -package mock
package files

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/minio/minio-go/v7"
)

type MinioRepository interface {
	PutObject(ctx context.Context, file models.File) (*minio.UploadInfo, error)
	GetObject(ctx context.Context, fileURL string) (*minio.Object, error)
	//RemoveObject(ctx context.Context, bucket string, fileName string) error
}
