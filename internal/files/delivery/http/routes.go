package http

import (
	"github.com/gabrielopesantos/myDrive-api/internal/files"
	"github.com/labstack/echo/v4"
)

// Map file routes
func MapFileRoutes(group *echo.Group, h files.Handlers) {
	//group.GET("/", handlers.GetFiles())
	group.POST("/", h.Insert())
}
