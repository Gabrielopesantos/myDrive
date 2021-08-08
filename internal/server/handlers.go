package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	authHttp "github.com/gabrielopesantos/myDrive-api/internal/auth/delivery/http"
	authRepository "github.com/gabrielopesantos/myDrive-api/internal/auth/repository"
	authService "github.com/gabrielopesantos/myDrive-api/internal/auth/service"
	apiMiddleware "github.com/gabrielopesantos/myDrive-api/internal/middleware"
	sessionRepository "github.com/gabrielopesantos/myDrive-api/internal/session/repository"
	sessionService "github.com/gabrielopesantos/myDrive-api/internal/session/service"
	userHttp "github.com/gabrielopesantos/myDrive-api/internal/user/delivery/http"
	usersRepository "github.com/gabrielopesantos/myDrive-api/internal/user/repository"
	usersService "github.com/gabrielopesantos/myDrive-api/internal/user/service"
	//filesHttp "github.com/gabrielopesantos/myDrive-api/internal/files/delivery/http"
	"github.com/gabrielopesantos/myDrive-api/pkg/metric"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	metrics, err := metric.CreateMetrics(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	if err != nil {
		s.logger.Errorf("CreateMetrics Error: %s", err)
	}
	s.logger.Infof(
		"Metrics available URL: %s, ServiceName: %s",
		s.cfg.Metrics.URL,
		s.cfg.Metrics.ServiceName,
	)
	// Init repos
	userRepo := usersRepository.NewUserRepository(s.db)
	authRepo := authRepository.NewAuthRepository(s.db)

	// MinIO Storage
	userMinioStorage := usersRepository.NewUserMinioRepo(s.minioClient)

	// Init Redis Repo
	sessionRedisRepo := sessionRepository.NewSessionRedisRepo(s.redisClient, s.cfg)
	userRedisRepo := usersRepository.NewUserRedisRepo(s.redisClient)

	// Init Services
	userServ := usersService.NewUserService(s.cfg, userRepo, userRedisRepo, userMinioStorage, s.logger)
	sessionServ := sessionService.NewSessionService(sessionRedisRepo, s.cfg)
	authServ := authService.NewAuthService(s.cfg, authRepo, s.logger)

	// Init handlers
	userHandlers := userHttp.NewUsersHandlers(s.cfg, userServ, sessionServ, s.logger)
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authServ, userServ, sessionServ, s.logger)
	//fileHandlers := filesHttp.NewFileHandlers(s.cfg, s.logger)

	// Init middleware
	mw := apiMiddleware.NewMiddlewareManager(sessionServ, userServ, s.cfg, s.logger)
	e.Use(mw.RequestLoggerMiddleware)

	// ?
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))

	e.Use(middleware.RequestID()) // Adds RequestID field to echo.Context struct
	e.Use(mw.MetricsMiddleware(metrics))

	// ?
	e.Use(middleware.Secure())        // Ver
	e.Use(middleware.BodyLimit("2M")) // Change to add files

	if s.cfg.Server.Debug {
		e.Use(mw.DebugMiddleware)
	}

	v1 := e.Group("/api/v1")

	userGroup := v1.Group("/users")
	userHttp.MapUserRoutes(userGroup, userHandlers, mw)

	authGroup := v1.Group("/auth")
	authHttp.MapAuthRoutes(authGroup, authHandlers)

	//filesGroup := v1.Group("/files")
	//filesHttp.MapFileRoutes(filesGroup, fileHandlers, mw)
	//
	return nil
}
