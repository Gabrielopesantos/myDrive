package service

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/gabrielopesantos/myDrive-api/internal/session"
	"github.com/opentracing/opentracing-go"
)

// Session Service
type SessionService struct {
	redisRepo session.RedisRepository
	cfg       *config.Config
}

// New Session Service Constructor
func NewSessionService(redisRepo session.RedisRepository, cfg *config.Config) session.Service {
	return &SessionService{
		redisRepo: redisRepo,
		cfg:       cfg,
	}
}

// Create session
func (s *SessionService) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SessionService.CreateSession")
	defer span.Finish()

	return s.redisRepo.CreateSession(ctx, session, expire)
}

// Get session by ID
func (s *SessionService) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SessionService.GetSessionByID")
	defer span.Finish()

	return s.redisRepo.GetSessionByID(ctx, sessionID)
}

// Delete session by ID
func (s *SessionService) DeleteSessionByID(ctx context.Context, sessionID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "SessionService.DeleteSessionByID")
	defer span.Finish()

	return s.redisRepo.DeleteSessionByID(ctx, sessionID)
}
