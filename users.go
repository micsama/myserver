package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string ` form:"username"`
	Password string ` form:"password"`
}

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
		message = "该用户名已被注册！请重新注册！"
		messagetype = "2"
	} else {
		sum = sha256.Sum256([]byte(login.Password))
		login.Password = string(sum[:])
		if err := userdb.Create(login).Error; err != nil {
			fmt.Println("插入失败", err)
			return
		}
		message = "注册成功！ 请直接登录吧！"
		messagetype = "1"
		tag = "signin.html"
	}
	c.HTML(http.StatusOK, tag, gin.H{
		"message":     message,
		"messagetype": messagetype,
		"m":           login.Password,
		"c":           sum,
	})

}
func signin(c *gin.Context) {
	var login, u User
	//从form输入绑定到login
	if err := c.ShouldBind(&login); err == nil {
		//绑定成功后先判断是否存在用户名
		userdb.Where("username = ?", login.Username).Take(&u)
		if u.Username == "" {
			message = "目标用户不存在，请自行注册！"
			messagetype = "2"
			tag = "register.html"
		} else {
			sum = sha256.Sum256([]byte(login.Password))
			if string(sum[:]) == u.Password {
				message = "登陆成功！"
				messagetype = "1"
				tag = "home.html"
				uname = " " + login.Username + " "
				c.SetCookie("name", login.Username, 7200, "/", "", false, true)
				c.SetCookie("acc", "admin", 7200, "/", "", false, true)
			} else {
				message = "密码错误！"
				messagetype = "3"
				tag = "signin.html"
			}
		}
		c.HTML(http.StatusOK, tag, gin.H{
			"uname":       uname,
			"message":     message,
			"messagetype": messagetype,
			"m":           login.Password,
			"c":           sum,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func logout(c *gin.Context) {
	c.SetCookie("name", "Shimin Li", -1, "/", "", false, true)
	c.SetCookie("acc", "Shimin Li", -1, "/", "", false, true)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"message":     "注销成功！",
		"messagetype": "1",
	})
}
