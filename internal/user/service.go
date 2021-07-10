package user

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"

	"github.com/google/uuid"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Users Service interface
type Service interface {
	GetUsers(ctx context.Context, pagQuery *utils.PaginationQuery) ([]*models.User, error)
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error)
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
}