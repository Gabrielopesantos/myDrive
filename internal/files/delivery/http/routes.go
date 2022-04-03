package http

import (
	"github.com/gabrielopesantos/myDrive-api/internal/files"
	"github.com/gabrielopesantos/myDrive-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

// MapFileRoutes Maps handlers to group
func MapFileRoutes(group *echo.Group, h files.Handlers, mw *middleware.MiddlewareManager) {
	group.Use(mw.AuthSessionMiddleware)
	group.GET("", h.GetUserFiles())
	group.GET("/:file_id", h.GetFileById())
	group.POST("", h.Insert())
}
