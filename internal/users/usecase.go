package users

import (
	"context"

	"github.com/google/uuid"

	"github.com/gabrielopesantos/myDrive-api/pkg/utl/models"
)

// Auth repository interface
type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
}
