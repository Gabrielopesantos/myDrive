package middleware

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	cfg *config.Config
	logger logger.Logger
}

// MiddlewareManager constructor
func NewMiddlewareManager(cfg *config.Config, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{
		cfg: cfg,
		logger: logger,
	}
}