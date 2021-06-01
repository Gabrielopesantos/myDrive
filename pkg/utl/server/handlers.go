package server

import (
	"github.com/labstack/echo/v4"

	userHttp "github.com/gabrielopesantos/myDrive-api/internal/users/delivery/http"
	usersRepository "github.com/gabrielopesantos/myDrive-api/internal/users/repository"
	usersUseCase "github.com/gabrielopesantos/myDrive-api/internal/users/usecase"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	uRepo := usersRepository.NewUsersRepository(s.db)
	usersUC := usersUseCase.NewUsersUseCase(uRepo)

	userHandlers := userHttp.NewUsersHandlers(usersUC)

	v1 := e.Group("/api/v1")

	usersGroup := v1.Group("/users")
	userHttp.MapUserRoutes(usersGroup, userHandlers)

	return nil
}
