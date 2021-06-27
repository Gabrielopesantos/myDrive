package http

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/auth"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type authHandlers struct {
	cfg    *config.Config
	logger logger.Logger
}

func NewAuthHandlers(cfg *config.Config, logger logger.Logger) auth.Handlers {
	return &authHandlers{
		cfg:    cfg,
		logger: logger,
	}
}

// TODO: Continue
func (h *authHandlers) Login() echo.HandlerFunc {
	type Login struct {
		// Incluir validate
		Email    string `json:"email" db:"email"`
		Password string `json:"password" db:"password"`
	}
	return func(ctx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(ctx, "authHandlers.Login")
		defer span.Finish()

		login := &Login{}
		if err := utils.ReadRequest(c, login); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

	}
}
