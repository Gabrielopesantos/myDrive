package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/gabrielopesantos/myDrive-api/internal/user"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) user.Repository {
	return &userRepo{db: db}
}

func (r *userRepo) GetUsers(ctx context.Context) ([]models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.GetUsers")
	defer span.Finish()

	var numUsers int
	if err := r.db.GetContext(ctx, &numUsers, getNumUsers); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetUsers.GetContext")
	}

	if numUsers == 0 {
		return []models.User{}, nil
	}

	var usersList = make([]models.User, 0, numUsers)
	if err := r.db.SelectContext(
		ctx,
		&usersList,
		getAllUsersQuery,
	); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetUsers.SelectContext")
	}

	return usersList, nil
}

func (r *userRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.Register")
	defer span.Finish()

	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, createUserQuery, &user.FirstName, &user.LastName, &user.Email,
		&user.Password, &user.Role, &user.About, &user.Avatar,
	).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "userRepo.Register.StructScan")
	}

	return u, nil
}

func (r *userRepo) GetByID(ctx context.Context, UserID uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.GetByID")
	defer span.Finish()

	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, getUserQuery, UserID).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetByID.StructScan")
	}

	return u, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.FindByEmail")
	defer span.Finish()

	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, findByEmailQuery, email).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "userRepo.FindByEmail.StructScan")
	}

	return u, nil
}
