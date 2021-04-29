package query

import (
	"github.com/labstack/echo/v4"

	"github.com/gabrielopesantos/main"
)

//List: prepares data for list queries
func List(u main.AuthUser) (*main.ListQuery, error) {
	switch true {
	case u.Role <= main.AdminRole:
		return nil, nil
	default:
		return nil, echo.ErrForbidden
	}
}
