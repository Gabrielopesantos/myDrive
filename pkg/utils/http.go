package utils

import (
	"context"
	"github.com/gabrielopesantos/myDrive-api/config"
	httpErrors "github.com/gabrielopesantos/myDrive-api/pkg/http_errors"
	"github.com/gabrielopesantos/myDrive-api/pkg/logger"
	"github.com/pkg/errors"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get context  with request id
func GetRequestCtx(c echo.Context) context.Context {
	return context.WithValue(c.Request().Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// Read request body and validate
func ReadRequest(ctx echo.Context, request interface{}) error {
	if err := ctx.Bind(request); err != nil {
		return err
	}
	return validate.StructCtx(ctx.Request().Context(), request)
}

// Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}

	return "./config/config-local"
}

// Create Session Cookie
func CreateSessionCookie(cfg *config.Config, session string) *http.Cookie {
	return &http.Cookie{
		Name:       cfg.Session.Name,
		Value:      session,
		Path:       "/",
		RawExpires: "",
		MaxAge:     cfg.Cookie.MaxAge,
		Secure:     cfg.Cookie.Secure,
		HttpOnly:   cfg.Cookie.HttpOnly,
		SameSite:   0,
	}
}

// DeleteSessionCookie - Deletes session cookie
func DeleteSessionCookie(c echo.Context, sessionName string) {
	c.SetCookie(&http.Cookie{
		Name:   sessionName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}

// UserCtxKey: Used for the user object in the context
type UserCtxKey struct{}

// Get user from context
//func GetUserFromCtx(ctx context.Context) (*models.User, error) {
//	user, ok := ctx.Value(UserCtxKey{}).(*models.User)
//	if !ok {
//		return nil, httpErrors.Unauthorized
//	}
//
//	return user, nil
//}

// Get user ip address
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}

// Error response with logging error for echo context
func ErrResponseWithLog(c echo.Context, logger logger.Logger, err error) error {
	logger.Errorf(
		"LogResponseError, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(c), GetIPAddress(c), err,
	)

	return c.JSON(httpErrors.ErrorResponse(err))
}

// Error response with logging error for echo context
func LogResponseError(c echo.Context, logger logger.Logger, err error) {
	logger.Errorf(
		"LogResponseError, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(c), GetIPAddress(c), err,
	)
}

// ReadImage Reads image from Request Body and checks if its type is valid
func ReadImage(c echo.Context, field string) (*multipart.FileHeader, error) {

	img, err := c.FormFile(field)
	if err != nil {
		return nil, err
	}

	// Check if content type of file is allowed
	if err = CheckImageContentType(img); err != nil {
		return nil, err
	}

	return img, nil
}

func CheckReturnImageFileContentType(fileContent []byte) (string, error) {
	contentType := http.DetectContentType(fileContent)

	extension, ok := allowedImageContentTypes[contentType]
	if !ok {
		return "", errors.New("file content type not allowed")
	}

	return  extension, nil
}

// ReadFile Same as ReadImage but without file extension restrictions
func ReadFile(c echo.Context, field string) (*multipart.FileHeader, error) {

	file, err := c.FormFile(field)
	if err != nil {
		return nil, err
	}

	return file, nil
}
