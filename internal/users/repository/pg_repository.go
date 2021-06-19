package repository

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/gabrielopesantos/myDrive-api/internal/users"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/models"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) users.Repository {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	log.Println("UsersRepo.Register")
	span, ctx := opentracing.StartSpanFromContext(ctx, "usersRepo.Register")
	defer span.Finish()

	log.Printf("user, %+v\n", user)
	u := &models.User{}
	log.Printf("u, %+v\n", u)
	log.Println("After creating an empty user struct")
	if err := r.db.QueryRowxContext(ctx, createUserQuery, &user.FirstName, &user.LastName, &user.Email,
		&user.Password, &user.Role, &user.About, &user.Avatar,
	).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "authUsers.Register.StructScan")
	}
	log.Printf("It would be nice to be here, %+v\n", u)

	return u, nil
}

func (r *UsersRepo) GetByID(ctx context.Context, UserID uuid.UUID) (*models.User, error) {
	log.Println("Repo")
	span, ctx := opentracing.StartSpanFromContext(ctx, "usersRepo.Register")
	defer span.Finish()

	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, getUserQuery, UserID).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "authUsers.GetByID.StructScan")
	}

	return u, nil
}
