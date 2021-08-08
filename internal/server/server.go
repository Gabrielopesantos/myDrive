package server

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
)

const (
	maxHeaderBytes = 1 << 20
)

type Server struct {
	echo        *echo.Echo
	cfg         *config.Config
	db          *sqlx.DB
	redisClient *redis.Client
	minioClient *minio.Client
	logger      logger.Logger
}

func NewServer(cfg *config.Config, db *sqlx.DB, redisClient *redis.Client, minioClient *minio.Client, logger logger.Logger) *Server {
	return &Server{
		echo:        echo.New(),
		cfg:         cfg,
		db:          db,
		redisClient: redisClient,
		minioClient: minioClient,
		logger:      logger,
	}
}

func (s *Server) Run() error {

	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		s.logger.Infof("Server is listening on port: %s", s.cfg.Server.Port)
		if err := s.echo.StartServer(server); err != nil {
			s.logger.Fatalf("Error starting server", err)
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	s.logger.Info("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)
}
