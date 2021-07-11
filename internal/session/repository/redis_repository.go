package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/gabrielopesantos/myDrive-api/internal/session"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"time"
)

const (
	basePrefix = "api-session"
)

// Session Redis Repo
type sessionRedisRepo struct {
	redisClient *redis.Client
	basePrefix  string
	cfg         *config.Config
}

// New Session Redis Repo Constructor
func NewSessionRedisRepo(redisClient *redis.Client, cfg *config.Config) session.RedisRepository {
	return &sessionRedisRepo{
		redisClient: redisClient,
		basePrefix:  basePrefix,
		cfg:         cfg,
	}
}

// Create session in Redis
func (r *sessionRedisRepo) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRedisRepo.CreateSession")
	defer span.Finish()

	session.SessionID = uuid.New().String()
	sessionKey := r.createKey(session.SessionID)

	sessionBytes, err := json.Marshal(&session)
	if err != nil {
		return "", errors.WithMessage(err, "sessionRedisRepo.CreateSession.json.Marshal")
	}

	if err = r.redisClient.Set(ctx, sessionKey, sessionBytes, time.Second*time.Duration(expire)).Err(); err != nil {
		return "", errors.Wrap(err, "sessionRedisRepo.CreateSession.redisClient.Set")
	}

	return sessionKey, nil
}

// GetSessionByID
func (r *sessionRedisRepo) GetSessionByID(ctx context.Context, sessionID string) (*models.Session, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRedisRepo.CreateSession")
	defer span.Finish()

	sessionKey := r.createKey(sessionID)

	sessionBytes, err := r.redisClient.Get(ctx, sessionKey).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "sessionRedisRepo.GetSessionByID.redisClient.Get")
	}
	sess := &models.Session{}
	if err = json.Unmarshal(sessionBytes, &sess); err != nil {
		return nil, errors.Wrap(err, "sessionRedisRepo.GetSessionByID.json.Unmarshal")
	}

	return sess, nil
}

func (r *sessionRedisRepo) createKey(sessionID string) string {
	return fmt.Sprintf("%s: %s", r.basePrefix, sessionID)
}
