//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package session

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Service Redis Repository interface
type RedisRepository interface {
	CreateSession(ctx context.Context, session *models.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error)
	DeleteSessionByID(ctx context.Context, sessionID string) error
}
