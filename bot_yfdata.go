package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var yfcount int

type yfdatasql struct {
	Name     string
	Count    int
	Otherinf string
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
	fmt.Println(yfdatamap)
	userdb.Model(yfdatasql{}).Count(&yfcount)
	userdb.Table("users")

}
