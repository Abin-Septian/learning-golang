package session

import (
	Users "learning-golang/api.example.com/pkg/types/users"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

// LoginData struct
type LoginData struct {
	Token string     `json:"token"`
	User  Users.User `json:"user"`
}

// Init function
func Init(DB *xorm.Engine) {
	db = DB
}
