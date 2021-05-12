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
	Phone    string
}

var sum [32]byte

func register(c *gin.Context) {
	var login, u User
	//从form输入绑定到login
	login.Username = c.PostForm("username")
	login.Password = c.PostForm("password")
	login.Phone = c.PostForm("phone")
	//绑定成功后先判断是否存在用户名
	tag = "user_register.html"
	vcode, _ := c.Cookie("vcode")
	vcode2 := c.PostForm("vcode")
	// vcode3 := sha256.Sum256([]byte(vcode2))
	// vcode4 := (*string)(unsafe.Pointer(&vcode3))
	// if vcode != *vcode4 {
	if vcode != vcode2 {
		message = "验证码错误！"
		messagetype = "2"
	} else {
		userdb = userdb.Table("users")
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
			tag = "user_signin.html"
		}
	}
	c.HTML(http.StatusOK, tag, gin.H{
		"message":     message,
		"messagetype": messagetype,
		"m":           login.Password,
		"c":           sum,
	})

}
func signin(c *gin.Context) {
	userdb = userdb.Table("users")
	var login, u User
	//从form输入绑定到login
	if err := c.ShouldBind(&login); err == nil {
		//绑定成功后先判断是否存在用户名
		userdb.Where("username = ?", login.Username).Take(&u)
		if u.Username == "" {
			message = "目标用户不存在，请自行注册！"
			messagetype = "2"
			tag = "user_register.html"
		} else {
			sum = sha256.Sum256([]byte(login.Password))
			if string(sum[:]) == u.Password {
				message = "登陆成功！"
				refreshyfdata()
				initstatus()
				messagetype = "1"
				tag = "user_home.html"
				uname = " " + login.Username + " "
				accsha := sha256.Sum256([]byte(login.Username + "admin"))
				c.SetCookie("name", login.Username, 7200, "/", "", false, true)
				c.SetCookie("acc", string(accsha[:]), 7200, "/", "", false, true)
			} else {
				message = "密码错误！"
				messagetype = "3"
				tag = "user_signin.html"
			}
		}
		c.HTML(http.StatusOK, tag, gin.H{
			"uname":       uname,
			"message":     message,
			"messagenum":  messagenum,
			"messagetype": messagetype,
			"m":           login.Password,
			"c":           sum,
			"sellout":     sellout,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func logout(c *gin.Context) {
	c.SetCookie("name", "Shimin Li", -1, "/", "", false, true)
	c.SetCookie("acc", "Shimin Li", -1, "/", "", false, true)
	c.HTML(http.StatusOK, "base_index.html", gin.H{
		"message":     "注销成功！",
		"messagetype": "1",
	})
}
