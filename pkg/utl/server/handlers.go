package server

import (
	"github.com/labstack/echo/v4"

	userHttp "github.com/gabrielopesantos/myDrive-api/interal/users/delivery/http"
	usersRepository "github.com/gabrielopesantos/myDrive-api/internal/users/repository"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	uRepo := usersRepository.NewUsersRepository(s.db)
	userHandlers := userHttp.NewAuthHandlers()

	v1 := e.Group("/api/v1")

	usersGroup := v1.Group("/users")
	userHttp.MapUserRoutes(usersGroup, userHandlers)
}
