package http

import (
	"github.com/labstack/echo/v4"
)

func MapUserRoutes(authGroup *echo.Group, u users.Handlers) {
	authGroup.POST("/register", u.Register())
	authGroup.GET("/:user_id", u.GetUserByID())
}
