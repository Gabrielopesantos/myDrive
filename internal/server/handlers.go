package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	apiMiddleware "github.com/gabrielopesantos/myDrive-api/internal/middleware"
	userHttp "github.com/gabrielopesantos/myDrive-api/internal/users/delivery/http"
	usersRepository "github.com/gabrielopesantos/myDrive-api/internal/users/repository"
	usersUseCase "github.com/gabrielopesantos/myDrive-api/internal/users/usecase"
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
	uRepo := usersRepository.NewUsersRepository(s.db)

	// Redis Repo
	uRedisRepo := usersRepository.NewUsersRedisRepo(s.redisClient)

	// Init useCases ?
	usersUC := usersUseCase.NewUsersUseCase(s.cfg, uRepo, uRedisRepo, s.logger)

	// Init handlers
	userHandlers := userHttp.NewUsersHandlers(s.cfg, usersUC, s.logger)

	// Init middleware
	mw := apiMiddleware.NewMiddlewareManager(s.cfg, s.logger)
	e.Use(mw.RequestLoggerMiddleware)

	// ?
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
		DisablePrintStack: true,
		DisableStackAll: true,
	}))

	e.Use(middleware.RequestID()) // Adds RequestID field to echo.Context struct
	e.Use(mw.MetricsMiddleware(metrics))

	// ?
	e.Use(middleware.Secure()) // Ver
	e.Use(middleware.BodyLimit("2M")) // Change to add files

	if s.cfg.Server.Debug {
		e.Use(mw.DebugMiddleware)
	}

	v1 := e.Group("/api/v1")

	usersGroup := v1.Group("/users")
	userHttp.MapUserRoutes(usersGroup, userHandlers, mw)

	//authHandlers := authHttp.NewAuthHandlers(s.cfg, usersUC, s.logger)
	//authGroup := v1.Group("/auth")

	return nil
}
