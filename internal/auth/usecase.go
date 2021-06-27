package auth

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Auth UseCase Interface
type UseCase interface {
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
}
