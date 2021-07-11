package middleware

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/session"
	"github.com/gabrielopesantos/myDrive-api/internal/user"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	sessionService session.Service
	userService    user.Service
	cfg            *config.Config
	logger         logger.Logger
}

// MiddlewareManager constructor
func NewMiddlewareManager(sessionService session.Service, userService user.Service, cfg *config.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{
		sessionService: sessionService,
		userService:    userService,
		cfg:            cfg,
		logger:         logger,
	}
}
