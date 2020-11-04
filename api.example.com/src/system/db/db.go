package db

import (
	// DB driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Connect function Connecting to database
func Connect(host string, port string, user string, pass string, database string, options string) (db *xorm.Engine, err error) {
	return xorm.NewEngine("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+database+"?charset=utf8"+options)
}

// Find function
func Find(DB *xorm.Engine, findBy interface{}, objects interface{}) error {
	return DB.Find(objects, findBy)
}

// FindBy function
func FindBy(DB *xorm.Engine, model interface{}) (err error) {
	_, err = DB.Get(model)
	return
}

// Exist function
func Exist(DB *xorm.Engine, model interface{}) (bool, error) {
	return DB.Get(model)
}

// Update function
func Update(DB *xorm.Engine, id int64, model interface{}) (err error) {
	_, err = DB.Id(id).Update(model)
	return
}

// Store function
func Store(DB *xorm.Engine, model interface{}) (err error) {
	_, err = DB.Insert(model)
	return
}

// Destroy function
func Destroy(DB *xorm.Engine, id int64, model interface{}) (err error) {
	_, err = DB.Id(id).Delete(model)
	return
}
