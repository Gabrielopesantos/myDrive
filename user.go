package main

import "time"

// User domain model
type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	// Password  string `json:"password"`

	// PhoneNumber string `json:"phone_number,omitempty"`
	// Address     string `json:"address,omitempty"`

	Active bool `json:"active"`

	LastLogin          time.Time `json:"last_login,omitempty"`
	LastPasswordChange time.Time `json:"last_password_change,omitempty"`
	// Token string `json:"-"`
	Role *Role `json:role"`
	// RoleID AcessRole `json:"-"`
}

//AuthUser: represents data stored in JWT token for user
type AuthUser struct {
	ID       int
	Username string
	Email    string
	Role     AccessRole
}

//ChangePassword: updates user's password related fields
func (u *User) ChangePassword(hash string) {
	u.Password = hash
	u.LastPasswordChange = time.Now()
}

//UpdateLastLogin: updates last login field
func (u *User) UpdateLastLogin(token string) {
	u.Token = token
	u.LastLogin = time.Now()
}
