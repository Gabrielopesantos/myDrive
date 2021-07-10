package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	authHttp "github.com/gabrielopesantos/myDrive-api/internal/auth/delivery/http"
	apiMiddleware "github.com/gabrielopesantos/myDrive-api/internal/middleware"
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

	// Redis Repo
	uRedisRepo := usersRepository.NewUserRedisRepo(s.redisClient)

	uService := usersService.NewUserService(s.cfg, uRepo, uRedisRepo, s.logger)

	// Init handlers
	uHandlers := userHttp.NewUsersHandlers(s.cfg, uService, s.logger)
	authHandlers := authHttp.NewAuthHandlers(s.cfg, uService, s.logger)

	// Init middleware
	mw := apiMiddleware.NewMiddlewareManager(s.cfg, s.logger)
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
