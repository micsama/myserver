package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var sellout string = ""
var yfcount int

type yfdatasql struct {
	Name     string
	Count    int
	Otherinf string
	Minimum  int
}

var yfdatamap []yfdatasql

func addnew(c *gin.Context) {
	var a yfdatasql
	a.Name = c.PostForm("name")
	a.Count, _ = strconv.Atoi(c.PostForm("num"))
	a.Otherinf = c.PostForm("inf")
	fmt.Println(a)
	userdb.Table("yfdatasql")
	if err := userdb.Create(a).Error; err != nil {
		fmt.Println("插入失败", err)
		userdb.Table("users")
	}
	c.HTML(http.StatusOK, "bot_yaofang_addnew_ok.html", gin.H{})

}
func refreshyfdata() {
	userdb.Table("yfdatasql")
	userdb.Find(&yfdatamap)
	userdb.Model(yfdatasql{}).Count(&yfcount)
	userdb.Table("users")
	sellout = ""
	for _, p := range yfdatamap {
		if p.Count <= p.Minimum {
			sellout = sellout + p.Name + " 余量为 " + strconv.Itoa(p.Count) + " !\n"
		}
	}
	if sellout != "" {
		sellout = "以下药品短缺！\n" + sellout
	} else {
		sellout = "目前药房存量充足！"
	}
	fmt.Println(sellout)

}
