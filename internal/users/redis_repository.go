package users

import (
	"context"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Users Redis repository interface
type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.User, error)
	SetUserCtx(ctx context.Context, key string, user *models.User, seconds int) error
}
