package user

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"

	"github.com/google/uuid"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Service  Users service interface
type Service interface {
	GetUsers(ctx context.Context, pagQuery *utils.PaginationQuery) ([]*models.User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	UpdateLastLogin(ctx context.Context, email string) error
	UploadAvatar(ctx context.Context, userID uuid.UUID, upload models.UploadInput) (*models.User, error)
}
