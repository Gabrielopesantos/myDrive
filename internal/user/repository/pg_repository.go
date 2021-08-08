package repository

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"
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

func (r *userRepo) GetUsers(ctx context.Context, pagQuery *utils.PaginationQuery) ([]*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.GetUsers")
	defer span.Finish()

	var numUsers int
	if err := r.db.GetContext(ctx, &numUsers, getNumUsersQuery); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetUsers.GetContext")
	}

	if numUsers == 0 {
		return []*models.User{}, nil
	}

	var users = make([]*models.User, 0, numUsers)
	if err := r.db.SelectContext(ctx, &users, getAllUsersQuery, pagQuery.GetOrderBy()); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetUsers.SelectContext")
	}

	return users, nil
}

func (r *userRepo) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.GetByID")
	defer span.Finish()

	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, getUserQuery, userID).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "userRepo.GetByID.StructScan")
	}

	return u, nil
}

func (r *userRepo) UpdateLastLogin(ctx context.Context, email string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.UpdateLastLogin")
	defer span.Finish()

	if _, err := r.db.ExecContext(ctx, updateUserLastLoginQuery, email); err != nil {
		return errors.Wrap(err, "userRepo.UpdateLastLogin.ExecContext")
	}

	return nil
}

func (r *userRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRepo.Update")
	defer span.Finish()

	u := &models.User{}
	if err := r.db.GetContext(ctx, u, updateUserQuery, &user.FirstName, &user.LastName, &user.Email,
		&user.Role, &user.About, &user.EmailVerified, &user.Avatar, &user.UserID); err != nil {
		return nil, errors.Wrap(err, "userRepo.Update.GetContext")
	}

	return u, nil
}
