package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `gorm:"column:username" form:"username"`
	Password string `gorm:"column:password" form:"password"`
}

func (u User) TableName() string {
	return "users"
}

var cc = "欢迎！！"
var sum [32]byte

func register(c *gin.Context) {
	var login, u User
	//从form输入绑定到login
	login.Username = c.PostForm("username")
	login.Password = c.PostForm("password")
	//绑定成功后先判断是否存在用户名
	tag = "register.html"
	userdb.Where("username = ?", login.Username).Take(&u)
	if u.Username != "" {
		cc = "该用户名已被注册！请重新注册！"
	} else {
		sum = sha256.Sum256([]byte(login.Password))
		login.Password = string(sum[:])
		if err := userdb.Create(login).Error; err != nil {
			fmt.Println("插入失败", err)
			return
		}
		cc = "注册成功！ 请直接登录吧！"
		tag = "signin.html"
	}
	c.HTML(http.StatusOK, tag, gin.H{
		"message": cc,
		"m":       login.Password,
		"c":       sum,
	})

}
func signin(c *gin.Context) {
	var login, u User
	var cc string
	//从form输入绑定到login
	if err := c.ShouldBind(&login); err == nil {
		//绑定成功后先判断是否存在用户名
		userdb.Where("username = ?", login.Username).Take(&u)
		if u.Username == "" {
			cc = "目标用户不存在，请自行注册！"
			tag = "register.html"
		} else {
			sum = sha256.Sum256([]byte(login.Password))
			if string(sum[:]) == u.Password {
				cc = "登陆成功！"
				tag = "home.html"
				uname = "欢迎！ " + login.Username + "   "
				c.SetCookie("name", login.Username, 1000, "/", myurl, false, true)
				c.SetCookie("acc", "admin", 1000, "/", myurl, false, true)
			} else {
				cc = "密码错误！"
				tag = "signin.html"
			}
		}
		c.HTML(http.StatusOK, tag, gin.H{
			"uname":   uname,
			"message": cc,
			"m":       login.Password,
			"c":       sum,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func logout(c *gin.Context) {
	c.SetCookie("name", "Shimin Li", -1, "/", myurl, false, true)
	c.SetCookie("acc", "Shimin Li", -1, "/", myurl, false, true)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": cc,
	})
}
