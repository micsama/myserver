package main

import (
	"fmt"
	_ "io"
	_ "os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var testing bool = true
var database = "./data/data.db"
var myurl string
var tag = "index.html"                      //default web
var uname, message, messagetype, acc string //cookie
var wg sync.WaitGroup                       //sync
var userdb *gorm.DB                         //db

func main() {
	gin.SetMode(gin.ReleaseMode)
	go initsql() //初始化数据库
	runweb()     //启动网页
	wg.Wait()
}
func initsql() {
	wg.Add(1)
	var dberr error
	username := "myserver"         //账号
	password := "4NX2EmiFaTGcNmLw" //密码
	host := "8.131.56.216"         //数据库地址，可以是Ip或者域名
	port := 3306                   //数据库端口
	Dbname := "myserver"           //数据库名
	timeout := "10s"               //连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	userdb, dberr = gorm.Open("mysql", dsn)
	if dberr != nil {
		fmt.Println(dberr)
	} else {
		fmt.Println("mysql connected success")
	}
	wg.Done()
}
