package http

import (
	"github.com/gabrielopesantos/myDrive-api/internal/middleware"
	"github.com/gabrielopesantos/myDrive-api/internal/user"
	"github.com/labstack/echo/v4"
)

func MapUserRoutes(userGroup *echo.Group, h user.Handlers, mw *middleware.MiddlewareManager) {
	userGroup.GET("/", h.GetUsers())
	userGroup.POST("/register", h.Register())
	userGroup.GET("/:user_id", h.GetUserByID())
}
