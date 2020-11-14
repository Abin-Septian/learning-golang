package users

import (
	"learning-golang/api.example.com/pkg/db"
)

// Users struct
type Users []Users

// User struct
type User db.Users

// TableName function
func (u *User) TableName() string {
	return "users"
}
