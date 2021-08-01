package files

import "github.com/labstack/echo/v4"

type Handlers interface {
	Insert() echo.HandlerFunc
}
