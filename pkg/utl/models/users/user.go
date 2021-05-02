package users

import (
	"strings"
	"time"

	base "github.com/gabrielopesantos/filesharing_api/pkg/utl/models"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	base.Base

	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Username  string  `json:"userName"`
	Email     string  `json:"email"`
	Password  string  `json:"password"`
	About     *string `json:"about"`
	Role      *string `json:"role"`

	EmailVerified bool `json:"emailVerified"`
	Active        bool `json:"active"`

	LastLogin          time.Time `json:"lastLogin"`
	LastPasswordChange time.Time `json:"lastPasswordChange"`
}

// type UserStorage interface {
// 	ListUsers() ([]User, error)
// 	GetUser(id uuid.UUID) (User, error)
// 	AddUser(u User) (User, error)
// 	UpdateUser(u User) (User, error)
// 	DeleteUser(id uuid.UUID) error
// }

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
