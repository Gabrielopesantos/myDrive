package http

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
)

// Files handlers

type fileHandlers struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewFileHandlers(cfg *config.Config, logger logger.Logger) *fileHandlers {
	return &fileHandlers{
		cfg:    cfg,
		logger: logger,
	}
}

//func (h *fileHandlers) GetFiles() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		//return
//	}
//}

//func (h *fileHandlers) Insert() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "fileHandlers.Insert")
//		defer span.Finish()
//
//
//
//
//	}
//}
