package auth

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Auth Service Interface
type Service interface {
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
}
