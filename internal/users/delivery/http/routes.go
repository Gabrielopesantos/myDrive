package http

import (
	"github.com/gabrielopesantos/myDrive-api/internal/users"
	"github.com/labstack/echo/v4"
)

func MapUserRoutes(usersGroup *echo.Group, h users.Handlers) {
	usersGroup.POST("/register", h.Register())
	usersGroup.GET("/:user_id", h.GetUserByID())
}
