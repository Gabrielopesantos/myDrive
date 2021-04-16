package main

//AccessRole represents access role type
type AccessRole int

const (
	// SuperAdminRole: has all permissions
	SuperAdminRole = 5

	// AdminRole: has all permissions
	AdminRole = 3

	// UserRole: StandardRole
	UserRole = 1
)

// Role model
type Role struct {
	ID          AccessRole `json:"id`
	AccessLevel AccessRole `json:"access_level"`
	Name        string     `json:"name"`
}
