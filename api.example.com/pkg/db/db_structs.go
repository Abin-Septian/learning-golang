package main

import (
	"time"
)

type Migrations struct {
	Id    int64     `xorm:"'id' pk autoincr" json:"id" schema:"id"`
	Name  string    `xorm:"name" json:"name" schema:"name"`
	RunOn time.Time `xorm:"run_on" json:"run_on" schema:"run_on"`
}

func (t Migrations) TableName() string {
	return "migrations"
}

func (t Migrations) SetId(id int64) {
	t.Id = id
}

func (t Migrations) GetId() int64 {
	return t.Id
}

type Users struct {
	Id        int64  `xorm:"'id' pk autoincr" json:"id" schema:"id"`
	FirstName string `xorm:"first_name" json:"first_name" schema:"first_name"`
	LastName  string `xorm:"last_name" json:"last_name" schema:"last_name"`
	Email     string `xorm:"email" json:"email" schema:"email"`
	Password  string `xorm:"password" json:"password" schema:"password"`
}
