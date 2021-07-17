package session

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Session Service Interface
type Service interface {
	CreateSession(ctx context.Context, session *models.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error)
	DeleteSessionByID(ctx context.Context, sessionID string) error
}
