package users

import (
	"context"

	"github.com/gabrielopesantos/myDrive-api/pkg/utl/models"
	"github.com/google/uuid"
)

type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	GetByID(ctx context.Context, UserID uuid.UUID) (*models.User, error)
}
