package http

import (
	"github.com/gabrielopesantos/myDrive-api/internal/auth"
	"github.com/labstack/echo/v4"
)

// MapAuthRoutes REST endpoints available from Auth service
func MapAuthRoutes(authGroup *echo.Group, h auth.Handlers) {
	authGroup.POST("/register", h.Register())
	authGroup.POST("/login", h.Login())
	authGroup.POST("/logout", h.Logout())
}
