package model

import "time"

type User struct {
	Base
	FirstName string
	LastName  string
	Username  string
	Email     string

	Active bool

	ProviderID string

	LastLogin          time.Time
	LastPasswordChange time.Time
}
