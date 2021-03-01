package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "io"
	_ "os"
	"sync"
)

var testing bool = true

var datebase = "data/date.db"
var myurl string
var tag = "index.html"                      //default web
var uname, message, messagetype, acc string //cookie
var wg sync.WaitGroup                       //sync
var userdb *gorm.DB                         //db

func main() {
	// f, err := os.OpenFile("./config", os.O_RDONLY, 0600)
	// defer f.Close()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	testing = false
	// } else {
	// 	_, _ = io.ReadAll(f)
	// 	testing = true
	// }
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
		fmt.Println("testokookkkk")
		userdb, dberr = gorm.Open("sqlite3", datebase)
	} else {
		fmt.Println("t----------111est")
		userdb, dberr = gorm.Open("mysql", "sql_www.dzmfg:XaFrbPKEh65t6JGd@(127.0.0.1:3306)/sql_www.dzmfg?charset=utf8&parseTime=True&loc=Local")
	}
	if dberr != nil {
		fmt.Println(dberr)
	} else {
		fmt.Println("success")
	}
	wg.Done()
}
