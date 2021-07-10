package http

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	utils "github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"net/http"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/gabrielopesantos/myDrive-api/internal/user"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type userHandlers struct {
	cfg     *config.Config
	service user.Service
	logger  logger.Logger
}

func NewUsersHandlers(cfg *config.Config, service user.Service, logger logger.Logger) user.Handlers {
	return &userHandlers{
		cfg:     cfg,
		service: service,
		logger:  logger,
	}
}

func (h *userHandlers) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "userHandlers.GetUsers")
		defer span.Finish()

		pagQuery, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		users, err := h.service.GetUsers(ctx, pagQuery)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)

			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, users)
	}
}

func (h *userHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "users.Register")
		defer span.Finish()

		newUser := &models.User{}
		if err := utils.ReadRequest(c, newUser); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createdUser, err := h.service.Register(ctx, newUser)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdUser)
	}
}

func (h *userHandlers) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "users.GetUserByID")
		defer span.Finish()

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		existingUser, err := h.service.GetByID(ctx, uID)
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, existingUser)
	}
}
