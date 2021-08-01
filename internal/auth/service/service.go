package service

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/auth"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type authService struct {
	cfg      *config.Config
	authRepo auth.Repository
	logger   logger.Logger
}

func NewAuthService(cfg *config.Config, authRepo auth.Repository, logger logger.Logger) auth.Service {
	return &authService{
		cfg:      cfg,
		authRepo: authRepo,
		logger:   logger,
	}
}

func (s *authService) Register(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userService.Register")
	defer span.Finish()

	if err := user.PrepareCreate(); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.Wrap(err, "userService.Register.PrepareCreate"))
	}

	createdUser, err := s.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	createdUser.SanitizePassword()

	token, err := utils.GenerateJWT(createdUser, s.cfg)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "userService.Register.GenerateJWT"))
	}

	return &models.UserWithToken{
		User:  createdUser,
		Token: token,
	}, nil
}

func (s *authService) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authService.Login")
	defer span.Finish()

	foundUser, err := s.authRepo.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if err = foundUser.ComparePasswords(user.Password); err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authService.Login.ComparePasswords"))
	}

	foundUser.SanitizePassword()

	token, err := utils.GenerateJWT(foundUser, s.cfg)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authService.Login.GenerateJWTToken"))
	}

	return &models.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}
