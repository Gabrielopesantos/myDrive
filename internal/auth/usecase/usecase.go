package usecase

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/auth"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	usersRepo "github.com/gabrielopesantos/myDrive-api/internal/user/repository"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type authUC struct {
	cfg       *config.Config
	usersRepo *usersRepo.UsersRepo
	logger    logger.Logger
}

func NewAuthUC(cfg *config.Config, usersRepo *usersRepo.UsersRepo, logger logger.Logger) auth.UseCase {
	return &authUC{
		cfg:       cfg,
		usersRepo: usersRepo,
		logger:    logger,
	}
}

func (u *authUC) Login(ctx context.Context, user *models.User) (*models.UserWithToken, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.Login")
	defer span.Finish()

	foundUser, err := u.usersRepo.GetByID(ctx, user.UserID)
	if err != nil {
		return nil, err
	}

	if err = foundUser.ComparePasswords(user.Password); err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUC.Login.ComparePasswords"))
	}

	foundUser.SanitizePassword()

	token, err := utils.GenerateJWT(foundUser, u.cfg)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUc.Login.GenerateJWTToken"))
	}

	return &models.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}
