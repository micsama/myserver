package main

import (
	"fmt"
	_ "net/http"
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
	// c.HTML(http.StatusOK, "yfmanager.html", gin.H{
	// 	"messag":      "添加成功",
	// 	"messagetype": "1",
	// })

}
func refreshyfdata() {
	userdb.Table("yfdatasql")
	userdb.Find(&yfdatamap)
	fmt.Println(yfdatamap)
	userdb.Model(yfdatasql{}).Count(&yfcount)
	userdb.Table("users")

}
