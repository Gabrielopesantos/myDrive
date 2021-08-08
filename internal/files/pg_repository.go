package files

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/minio/minio-go/v7"
)

type Repository interface {
	RecordFileInsertion(ctx context.Context, file *models.File) (*minio.UploadInfo, error)
}
