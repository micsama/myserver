package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
)

const testing = true

var datebase = "data/date.db"
var myurl string
var tag = "index.html" //default web
var uname, acc string  //cookie
var wg sync.WaitGroup  //sync
var userdb *gorm.DB    //db

func main() {
	if testing {
		myurl = "localhost"
	} else {
		myurl = "www.dzmfg.icu"
	}
	go initlisten()
	initsql()
	runweb()
	wg.Wait()
}
func initsql() {
	wg.Add(1)
	var dberr error
	if testing {
		userdb, dberr = gorm.Open("sqlite3", datebase)
	} else {
		userdb, dberr = gorm.Open("mysql", "sql_www.dzmfg:XaFrbPKEh65t6JGd@(127.0.0.1:3306)/sql_www.dzmfg?charset=utf8&parseTime=True&loc=Local")
	}
	if dberr != nil {
		fmt.Println(dberr)
	} else {
		fmt.Println("success")
	}
	wg.Done()
}
