package usecase
//
//import (
//	"context"
//	"github.com/gabrielopesantos/myDrive-api/config"
//	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
//	"github.com/gabrielopesantos/myDrive-api/pkg/utl/models"
//	"github.com/opentracing/opentracing-go"
//	"testing"
//
//	"github.com/golang/mock/gomock"
//)
//
//func TestUsersUC_Register(t *testing.T) {
//	t.Parallel()
//
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	cfg := &config.Config{
//		Server: config.ServerConfig{
//			JwtSecretKey: "secret",
//		},
//		Logger: config.LoggerConfig{
//			Development: true,
//			DisableCaller: false,
//			DisableStacktrace: false,
//			Encoding: "json",
//		},
//	}
//
//	apiLogger := logger.NewApiLogger(cfg)
//	mockUsersRepo := mock.NewMockRepository(ctrl)
//	usersUC := NewUsersUseCase(cfg, mockUsersRepo, apiLogger)
//
//	user := &models.User{
//		Password: "123456",
//		Email: "email@gmail.com",
//	}
//
//	ctx := context.Background()
//	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "usersUC.UploadAvatar")
//	defer span.Finish()
//
//}