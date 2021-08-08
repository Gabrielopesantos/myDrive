package files

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/minio/minio-go/v7"
)

// Service files service interface
type Service interface {
	Insert(ctx context.Context, file *models.File) (*minio.UploadInfo, error)
}
