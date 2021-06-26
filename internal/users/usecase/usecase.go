package usecase

import (
	"context"
	"fmt"
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/users"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/utl/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/models"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/utils"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

const (
	basePrefix    = "api-users"
	cacheDuration = 3600
)

type usersUC struct {
	cfg       *config.Config
	usersRepo users.Repository
	redisRepo users.RedisRepository
	logger    logger.Logger
}

func NewUsersUseCase(cfg *config.Config, usersRepo users.Repository, redisRepo users.RedisRepository, logger logger.Logger) *usersUC {
	return &usersUC{
		cfg:       cfg,
		usersRepo: usersRepo,
		redisRepo: redisRepo,
		logger:    logger,
	}
}

func (u *usersUC) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usersUC.Register")
	defer span.Finish()

	if err := user.PrepareCreate(); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.Wrap(err, "usersUC.Register.PrepareCreate"))
	}

	createdUser, err := u.usersRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.SanitizePassword()

	token, err := utils.GenerateJWT(createdUser, u.cfg)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "usersUC.Register.GenerateJWT"))
	}

	return &models.UserWithToken{
		User:  createdUser,
		Token: token,
	}, nil
}

func (u *usersUC) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usersUC.GetByID")
	defer span.Finish()

	cachedUser, err := u.redisRepo.GetByIDCtx(ctx, u.generateUserKey(userID.String()))
	if err != nil {
		u.logger.Errorf("usersUC.GetByID.RedisRepo.GetByIDCtx: %v", err)
	}
	if cachedUser != nil {
		return cachedUser, nil
	}

	user, err := u.usersRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err = u.redisRepo.SetUserCtx(ctx, u.generateUserKey(userID.String()), user, cacheDuration); err != nil {
		u.logger.Errorf("usersUC.GetByID.RedisRepo.SetUserCtx: %v", err)
	}

	user.SanitizePassword()

	return user, nil
}

func (u *usersUC) generateUserKey(userID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, userID)
}
