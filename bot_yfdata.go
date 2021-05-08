package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var sellout string = ""
var yfcount int

type yfdatasql struct {
	Id       int
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
	a.Minimum, _ = strconv.Atoi(c.PostForm("min"))
	fmt.Println(a)
	if err := userdb.Create(&a).Error; err != nil {
		fmt.Println("插入失败", err)
	} else {
		fmt.Println(err)
	}
	c.HTML(http.StatusOK, "bot_yaofang_addnew_ok.html", gin.H{
		"message": "添加成功！",
	})
}
func refreshyfdata() {
	userdb = userdb.Table("yfdatasqls")
	userdb.Find(&yfdatamap)
	userdb.Model(yfdatasql{}).Count(&yfcount)
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
func updatayf(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	num, _ := strconv.Atoi(c.Query("num"))
	var u yfdatasql
	userdb = userdb.Table("yfdatasqls")
	userdb.Where("id = ?", id).Take(&u)
	fmt.Println(id)
	fmt.Println(u)
	num = num + u.Count
	userdb.Model(&u).Update("count", num)
}
