package db

import (
	// DB driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Connect function Connecting to database
func Connect() (db *xorm.Engine, err error) {
	return xorm.NewEngine("mysql", "root:@tcp(localhost:3306)/test_go_api?charset=utf8")
}
