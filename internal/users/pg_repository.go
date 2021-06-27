package users

import (
	"context"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/google/uuid"
)

// Users Repository Interface
type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	GetByID(ctx context.Context, UserID uuid.UUID) (*models.User, error)
}
