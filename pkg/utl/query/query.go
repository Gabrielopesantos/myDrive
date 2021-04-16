package query

import (
	"github.com/labstack/echo/v4"

	"github.com/gabrielopesantos/filesharing_api"
)

//List: prepares data for list queries
func List(u filesharing_api.AuthUser) (*filesharing_api.ListQuery, error) {
	switch true {
	case u.Role <= filesharing_api.AdminRole:
		return nil, nil
	default:
		return nil, echo.ErrForbidden
	}
}
