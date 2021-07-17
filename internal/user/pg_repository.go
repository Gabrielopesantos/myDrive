package user

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/google/uuid"
)

// Users Postgres Store Interface
type Repository interface {
	GetUsers(ctx context.Context, pagQuery *utils.PaginationQuery) ([]*models.User, error)
	Register(ctx context.Context, user *models.User) (*models.User, error)
	GetByID(ctx context.Context, UserID uuid.UUID) (*models.User, error)
	FindByEmail(ctx context.Context, UserID string) (*models.User, error)
	UpdateLastLogin(ctx context.Context, email string) error
}
