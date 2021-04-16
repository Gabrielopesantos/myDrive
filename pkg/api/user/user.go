// Package user contains user application services
package user

import (
	"github.com/labstack/echo/v4"

	"githuv.com/gabrielopesantos/filesharing_api"
	"githuv.com/gabrielopesantos/filesharing_api/pkg/utl/query"
)

//Create: creates a new user account
func (u User) Create(c echo.Context, req filesharing_api.User) (filesharing_api.User, error) {
	if err := u.rbac
}