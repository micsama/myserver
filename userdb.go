// Package main provides ...
package main

import (
	_ "fmt"
)

type User struct {
	Username   string `gorm:"column:username" form:"username"`
	Password   string `gorm:"column:password" form:"password"`
	createTime int64  `gorm:"column:createtime"`
}

func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "users"
}
