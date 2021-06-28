package http

import "github.com/labstack/echo/v4"

// Map file routes
func MapFileRoutes(group *echo.Group, handlers fileHandlers) {
	//group.GET("/", handlers.GetFiles())
	group.POST("/", handlers.Insert())
}
