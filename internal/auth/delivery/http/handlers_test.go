package http

import (
	"github.com/gabrielopesantos/myDrive-api/config"
	mockAuth "github.com/gabrielopesantos/myDrive-api/internal/auth/mock"
	"github.com/gabrielopesantos/myDrive-api/internal/models"
	mockSess "github.com/gabrielopesantos/myDrive-api/internal/session/mock"
	mockUser "github.com/gabrielopesantos/myDrive-api/internal/user/mock"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/gabrielopesantos/myDrive-api/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthHandlers_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthService := mockAuth.NewMockService(ctrl)
	mockUserService := mockUser.NewMockService(ctrl)
	mockSessService := mockSess.NewMockService(ctrl)

	cfg := &config.Config{
		Session: config.Session{
			Expire: 10,
		},
		Logger: config.LoggerConfig{
			Development: true,
		},
	}

	t.Run("Successful login", func(t *testing.T) {

		apiLogger := logger.NewApiLogger(cfg)
		authHandlers := NewAuthHandlers(cfg, mockAuthService, mockUserService, mockSessService, apiLogger)

		user := &models.User{
			FirstName: "First",
			LastName:  "Last",
			Email:     "email@email.com",
			Password:  "123456",
		}

		buf, err := utils.AnyToBytesBuffer(user)

		require.NoError(t, err, "Err should be nil")
		require.NotNil(t, buf, "Bytes buffer should not be nil")

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", strings.NewReader(buf.String()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		ctx := utils.GetRequestCtx(c)
		span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "auth.Register")
		defer span.Finish()

		handlerFunc := authHandlers.Register()

		userID := uuid.New()
		userWithToken := &models.UserWithToken{
			User: &models.User{
				UserID: userID,
			},
		}
		sess := &models.Session{
			UserID: userID,
		}

		session := "session"

		mockAuthService.EXPECT().Register(ctxWithTrace, gomock.Eq(user)).Return(userWithToken, nil)
		mockSessService.EXPECT().CreateSession(ctxWithTrace, gomock.Eq(sess), 10).Return(session, nil)

		err = handlerFunc(c)
		require.NoError(t, err)
		require.Nil(t, err)
	})
}
