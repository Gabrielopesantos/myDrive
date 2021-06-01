package usecase

import (
	"context"

	"github.com/gabrielopesantos/myDrive-api/internal/users"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/utl/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/models"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

// const (
// 	basePrefix    = "api-auth:"
// 	cacheDuration = 3600
// )

type usersUC struct {
	// cfg       *config.Config
	usersRepo users.Repository
}

func NewUsersUseCase(usersRepo users.Repository) *usersUC {
	return &usersUC{usersRepo: usersRepo}
}

func (u *usersUC) Register(ctx context.Context, user *models.User) (*models.UserWithoutToken, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usersUC.Register")
	defer span.Finish()

	if err := user.PrepareCreate(); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.Wrap(err, "usersUC.Register.PrepareCreate"))
	}

	createdUser, err := u.usersRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.UserWithoutToken{
		User: createdUser,
	}, nil
}

func (u *usersUC) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "usersUC.GetByID")
	defer span.Finish()

	user, err := u.usersRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// user.SanitizePassword()

	return user, nil
}
