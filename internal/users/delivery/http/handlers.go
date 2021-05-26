package http

import ("github.com/gabrielopesantos/myDrive-api/config",
		"github.com/gabrielopesantos/MyDrive-api/internal/users")

type userHandlers struct {
	cfg *config.Config
}

func NewAuthHandlers() user.Handlers {
	// return &authHandlers{cfg: cfg}
	return &authHandlers{}
}
