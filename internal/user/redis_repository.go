//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package user

import (
	"context"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Users Redis store interface
type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.User, error)
	SetUserCtx(ctx context.Context, key string, user *models.User, seconds int) error
}
