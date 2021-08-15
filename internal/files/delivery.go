package files

import "github.com/labstack/echo/v4"

type Handlers interface {
	GetFileById() echo.HandlerFunc
	Insert() echo.HandlerFunc
	//Delete() echo.HandlerFunc
	//Update() echo.HandlerFunc
}
