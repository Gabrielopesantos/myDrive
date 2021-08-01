package repository

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/internal/auth"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) auth.Repository {
	return &authRepo{db: db}
}

func (r *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Register")
	defer span.Finish()

	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, createUserQuery, &user.FirstName, &user.LastName, &user.Email,
		&user.Password, &user.Role, &user.About, &user.Avatar,
	).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "authRepo.Register.StructScan")
	}

	return u, nil
}

func (r *authRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.FindByEmail")
	defer span.Finish()

	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, findByEmailQuery, email).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "userRepo.FindByEmail.StructScan")
	}

	return u, nil
}
