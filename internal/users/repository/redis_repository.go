package repository

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"time"

	"github.com/gabrielopesantos/myDrive-api/internal/users"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/models"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
)

// Users Redis repository
type usersRedisRepo struct {
	redisClient *redis.Client
}

// Users redis repo constructor
func NewUsersRedisRepo(redisClient *redis.Client) users.RedisRepository {
	return &usersRedisRepo{
		redisClient: redisClient,
	}
}

// Get user by ID
func (u *usersRedisRepo) GetByIDCtx(ctx context.Context, key string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsersRedisRepo.GetByIDCtx")
	defer span.Finish()

	userBytes, err := u.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "usersRedisRepo.GetByIDCtx.redisClient.Get")
	}
	user := &models.User{}
	if err = json.Unmarshal(userBytes, user); err != nil {
		return nil, errors.Wrap(err, "usersRedisRepo.GetByIDCtx.json.Unmarshal")
	}

	return user, nil
}

// SetUserCtx: Cache user for a certain duration
func (u *usersRedisRepo) SetUserCtx(ctx context.Context, key string, user *models.User, seconds int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UsersRedisRepo.SetUserCtx")
	defer span.Finish()

	userBytes, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err,"usersRedisRepo.SetUserCtx.json.Marshal")
	}

	if err = u.redisClient.Set(ctx, key, userBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "usersRedisRepo.SetUserCtx.Set")
	}

	return nil
}
