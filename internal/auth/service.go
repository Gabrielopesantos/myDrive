//go:generate mockgen -source service.go -destination mock/service_mock.go -package mock
package auth

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
)

// Auth Service Interface
type Service interface {
	Register(ctx context.Context, user *models.User) (*models.UserWithToken, error)
	Login(ctx context.Context, user *models.User) (*models.UserWithToken, error)
}
