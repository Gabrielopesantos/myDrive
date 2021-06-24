package server

import (
	"github.com/labstack/echo/v4"

	userHttp "github.com/gabrielopesantos/myDrive-api/internal/users/delivery/http"
	usersRepository "github.com/gabrielopesantos/myDrive-api/internal/users/repository"
	usersUseCase "github.com/gabrielopesantos/myDrive-api/internal/users/usecase"
	//apiMiddleware "github.com/gabrielopesantos/myDrive-api/internal/middleware"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	// Init repos
	uRepo := usersRepository.NewUsersRepository(s.db)

	// Init useCases ?
	usersUC := usersUseCase.NewUsersUseCase(uRepo)

	// Init handlers
	userHandlers := userHttp.NewUsersHandlers(usersUC)

	// Init middleware
	//mw := apiMiddleware.NewMiddlewareManager(s.cfg, s.logger)

	v1 := e.Group("/api/v1")

	usersGroup := v1.Group("/users")
	userHttp.MapUserRoutes(usersGroup, userHandlers)

	return nil
}
