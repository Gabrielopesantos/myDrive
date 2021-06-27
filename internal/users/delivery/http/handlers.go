package http

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	utils "github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"net/http"

	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/gabrielopesantos/myDrive-api/internal/users"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type userHandlers struct {
	cfg     *config.Config
	usersUC users.UseCase
	logger  logger.Logger
}

func NewUsersHandlers(cfg *config.Config, usersUC users.UseCase, logger logger.Logger) users.Handlers {
	return &userHandlers{
		cfg:     cfg,
		usersUC: usersUC,
		logger:  logger,
	}
}

func (u *userHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "users.Register")
		defer span.Finish()

		user := &models.User{}
		if err := utils.ReadRequest(c, user); err != nil {
			utils.LogResponseError(c, u.logger, err)
			//u.logger.Errorf("usersHandlers.Register.ReadRequest", err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createdUser, err := u.usersUC.Register(ctx, user)
		if err != nil {
			utils.LogResponseError(c, u.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdUser)
	}
}

func (u *userHandlers) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "users.GetUserByID")
		defer span.Finish()

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		user, err := u.usersUC.GetByID(ctx, uID)
		if err != nil {
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, user)
	}
}
