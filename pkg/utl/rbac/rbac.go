package rbac

import (
	"github.com/labstack/echo/v4"

	"github.com/gabrielopesantos/filesharing_api"
)

// Service is an RBAC application service
type Service struct{}

func checkBool(b bool) error {
	if b {
		return nil
	}

	return echo.ErrForbidden
}

//User: returns user data stored in jwt token
func (s Service) User(c echo.Context) filesharing_api.AuthUser {

}
