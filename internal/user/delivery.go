package user

import "github.com/labstack/echo/v4"

// Users HTTP Handlers Interface
type Handlers interface {
	GetUsers() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
	GetMe() echo.HandlerFunc
	UploadAvatar() echo.HandlerFunc
}
