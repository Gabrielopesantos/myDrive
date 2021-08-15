package http

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	"github.com/gabrielopesantos/myDrive-api/internal/auth"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	"github.com/gabrielopesantos/myDrive-api/internal/session"
	"github.com/gabrielopesantos/myDrive-api/internal/user"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"net/http"
)

type authHandlers struct {
	cfg            *config.Config
	authService    auth.Service
	userService    user.Service
	sessionService session.Service
	logger         logger.Logger
}

func NewAuthHandlers(cfg *config.Config, authService auth.Service, userService user.Service, sessionService session.Service, logger logger.Logger) auth.Handlers {
	return &authHandlers{
		cfg:            cfg,
		authService:    authService,
		userService:    userService,
		sessionService: sessionService,
		logger:         logger,
	}
}

// Register godoc
// @Summary Register a new user
// @Description register new user, returns user and token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.UserWithToken
// @Router /auth/register [post]
func (h *authHandlers) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "users.Register")
		defer span.Finish()

		newUser := &models.User{}
		if err := utils.ReadRequest(c, newUser); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		createdUser, err := h.authService.Register(ctx, newUser)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		session, err := h.sessionService.CreateSession(
			ctx,
			&models.Session{
				UserID: createdUser.User.UserID,
			},
			h.cfg.Session.Expire)

		c.SetCookie(utils.CreateSessionCookie(h.cfg, session))

		return c.JSON(http.StatusCreated, createdUser)
	}
}

func (h *authHandlers) Login() echo.HandlerFunc {
	type Login struct {
		Email    string `json:"email" db:"email" validate:"omitempty,lte=60,email"`
		Password string `json:"password,omitempty" db:"password" validate:"required,gte=3"`
	}
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "authHandlers.Login")
		defer span.Finish()

		login := &Login{}
		if err := utils.ReadRequest(c, login); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		userWithToken, err := h.authService.Login(ctx, &models.User{
			Email:    login.Email,
			Password: login.Password,
		})
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err)) // 404 if account does not exist and 401 if fields were wrong?
		}

		if err = h.userService.UpdateLastLogin(ctx, userWithToken.User.Email); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(http.StatusInternalServerError, httpErrors.NewInternalServerError(httpErrors.InternalServerError))
		}

		sess, err := h.sessionService.CreateSession(ctx, &models.Session{UserID: userWithToken.User.UserID}, h.cfg.Session.Expire)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		c.SetCookie(utils.CreateSessionCookie(h.cfg, sess))

		return c.JSON(http.StatusOK, userWithToken)

	}
}

func (h *authHandlers) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "authHandlers.Logout")
		defer span.Finish()

		cookie, err := c.Cookie("session-id")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				utils.LogResponseError(c, h.logger, err)
				return c.JSON(http.StatusUnauthorized, httpErrors.NewUnauthorizedError(err))
			}
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(http.StatusInternalServerError, httpErrors.NewInternalServerError(err))
		}

		if err = h.sessionService.DeleteSessionByID(ctx, cookie.Value); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		utils.DeleteSessionCookie(c, h.cfg.Session.Name)

		return c.NoContent(http.StatusOK)
	}
}
