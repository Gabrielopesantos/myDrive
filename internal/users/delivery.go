package users

import "github.com/labstack/echo/v4"

// Users HTTP Handlers Interface
type Handlers interface {
	Register() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
}
