package http

import (
	"github.com/gabrielopesantos/myDrive-api/internal/users"
	"github.com/labstack/echo/v4"
)

func MapUserRoutes(usersGroup *echo.Group, u users.Handlers) {
	usersGroup.POST("/register", u.Register())
	usersGroup.GET("/:user_id", u.GetUserByID())
}
