package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
	FirstName     string    `json:"first_name" db:"first_name"`
	LastName      string    `json:"last_name" db:"last_name"`
	Email         string    `json:"email" db:"email"`
	Password      string    `json:"password,omitempty" db:"password"`
	Role          *string   `json:"role" db:"role"`
	About         *string   `json:"about,omitempty" db:"about"`
	Avatar        *string   `json:"avatar,omitempty" db:"avatar"`
	EmailVerified bool      `json:"is_email_verified" db:"is_email_verified"`
	LastLogin     time.Time `json:"last_login,omitempty" db:"last_login" validate:"omitempty"`
	Base
}

func (u *User) HashPassword() error {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPwd)
	return nil
}

func (u *User) ComparePasswords(pwd string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(u.Password)); err != nil {
		return err
	}
	return nil
}

func (u *User) PrepareCreate() error {

	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}
	return nil
}

// Sanitze user password
func (u *User) SanitizePassword() {
	u.Password = ""
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
