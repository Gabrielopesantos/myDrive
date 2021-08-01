package service

import (
	"context"
	"fmt"
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/gabrielopesantos/myDrive-api/internal/user"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	utils "github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

const (
	basePrefix    = "api-user"
	cacheDuration = 3600
)

type userService struct {
	cfg       *config.Config
	userRepo  user.Repository
	redisRepo user.RedisRepository
	logger    logger.Logger
}

func NewUserService(cfg *config.Config, userRepo user.Repository, redisRepo user.RedisRepository, logger logger.Logger) user.Service {
	return &userService{
		cfg:       cfg,
		userRepo:  userRepo,
		redisRepo: redisRepo,
		logger:    logger,
	}
}

// Hm?
func (s *userService) GetUsers(ctx context.Context, pagQuery *utils.PaginationQuery) ([]*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.GetUsers")
	defer span.Finish()

	return s.userRepo.GetUsers(ctx, pagQuery)
}

func (s *userService) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.GetByID")
	defer span.Finish()

	cachedUser, err := s.redisRepo.GetByIDCtx(ctx, s.generateUserKey(userID.String()))
	if err != nil {
		s.logger.Errorf("userService.GetByID.RedisRepo.GetByIDCtx: %v", err)
	}
	if cachedUser != nil {
		return cachedUser, nil
	}

	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err = s.redisRepo.SetUserCtx(ctx, s.generateUserKey(userID.String()), user, cacheDuration); err != nil {
		s.logger.Errorf("userService.GetByID.RedisRepo.SetUserCtx: %v", err)
	}

	user.SanitizePassword()

	return user, nil
}

func (s *userService) generateUserKey(userID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, userID)
}

func (s *userService) UpdateLastLogin(ctx context.Context, email string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.GetUsers")
	defer span.Finish()

	return s.userRepo.UpdateLastLogin(ctx, email)
}
