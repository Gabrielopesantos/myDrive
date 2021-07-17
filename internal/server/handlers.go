package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	authHttp "github.com/gabrielopesantos/myDrive-api/internal/auth/delivery/http"
	authService "github.com/gabrielopesantos/myDrive-api/internal/auth/service"
	apiMiddleware "github.com/gabrielopesantos/myDrive-api/internal/middleware"
	sessionRepository "github.com/gabrielopesantos/myDrive-api/internal/session/repository"
	sessionService "github.com/gabrielopesantos/myDrive-api/internal/session/service"
	userHttp "github.com/gabrielopesantos/myDrive-api/internal/user/delivery/http"
	usersRepository "github.com/gabrielopesantos/myDrive-api/internal/user/repository"
	usersService "github.com/gabrielopesantos/myDrive-api/internal/user/service"
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
	uRepo := usersRepository.NewUserRepository(s.db)

	// Init Redis Repo
	sRedisRepo := sessionRepository.NewSessionRedisRepo(s.redisClient, s.cfg)
	uRedisRepo := usersRepository.NewUserRedisRepo(s.redisClient)

	// Init Services
	uService := usersService.NewUserService(s.cfg, uRepo, uRedisRepo, s.logger)
	sService := sessionService.NewSessionService(sRedisRepo, s.cfg)
	aService := authService.NewAuthService(s.cfg, uRepo, s.logger)

	// Init handlers
	uHandlers := userHttp.NewUsersHandlers(s.cfg, uService, sService, s.logger)
	authHandlers := authHttp.NewAuthHandlers(s.cfg, aService, uService, sService, s.logger)

	// Init middleware
	mw := apiMiddleware.NewMiddlewareManager(sService, uService, s.cfg, s.logger)
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

	usersGroup := v1.Group("/users")
	userHttp.MapUserRoutes(usersGroup, uHandlers, mw)

	authGroup := v1.Group("/auth")
	authHttp.MapAuthRoutes(authGroup, authHandlers)

	return nil
}
