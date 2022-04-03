//go:generate mockgen -source service.go -destination mock/service_mock.go -package mock
package files

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

// Service files service interface
type Service interface {
	Insert(ctx context.Context, file *models.File) (*models.File, error)
	GetFileById(ctx context.Context, fileID uuid.UUID) (*models.File, error)
	RetrieveObjectFromBucket(ctx context.Context, file *models.File) (*minio.Object, error)
}
