package repository

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"time"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/gabrielopesantos/myDrive-api/internal/user"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
)

// Users Redis repository
type userRedisRepo struct {
	redisClient *redis.Client
}

// Users redis repo constructor
func NewUserRedisRepo(redisClient *redis.Client) user.RedisRepository {
	return &userRedisRepo{
		redisClient: redisClient,
	}
}

// Get user by ID
func (r *userRedisRepo) GetByIDCtx(ctx context.Context, key string) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserRedisRepo.GetByIDCtx")
	defer span.Finish()

	userBytes, err := r.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "userRedisRepo.GetByIDCtx.redisClient.Get")
	}
	user := &models.User{}
	if err = json.Unmarshal(userBytes, user); err != nil {
		return nil, errors.Wrap(err, "userRedisRepo.GetByIDCtx.json.Unmarshal")
	}

	return user, nil
}

// SetUserCtx: Cache user for a certain duration
func (r *userRedisRepo) SetUserCtx(ctx context.Context, key string, user *models.User, seconds int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "userRedisRepo.SetUserCtx")
	defer span.Finish()

	userBytes, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "userRedisRepo.SetUserCtx.json.Marshal")
	}

	if err = r.redisClient.Set(ctx, key, userBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "userRedisRepo.SetUserCtx.Set")
	}

	return nil
}
