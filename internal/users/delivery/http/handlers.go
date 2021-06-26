package http

import (
	"log"
	"net/http"

	"github.com/gabrielopesantos/myDrive-api/internal/users"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/utl/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/models"
	"github.com/gabrielopesantos/myDrive-api/pkg/utl/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

type userHandlers struct {
	// cfg     *config.Config
	usersUC users.UseCase
}

func NewUsersHandlers(usersUC users.UseCase) users.Handlers {
	// return &authHandlers{cfg: cfg}
	return &userHandlers{usersUC: usersUC}
}

func (u *userHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "users.Register")
		defer span.Finish()

		user := &models.User{}
		if err := utils.ReadRequest(c, user); err != nil {
			// utils.LogResponseError(c, h.)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createdUser, err := u.usersUC.Register(ctx, user)
		log.Printf("ERRO: %v", err)
		if err != nil {
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
