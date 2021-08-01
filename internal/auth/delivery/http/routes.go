package http

import (
	"github.com/gabrielopesantos/myDrive-api/internal/auth"
	"github.com/labstack/echo/v4"
)

// Map Auth routes
func MapAuthRoutes(authGroup *echo.Group, h auth.Handlers) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
	authGroup.POST("/logout", h.Logout())
}
