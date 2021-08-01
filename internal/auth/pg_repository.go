package auth

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Auth Postgres Repository Interface
type Repository interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
