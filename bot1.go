// Package main provides ...
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func bot1(c *gin.Context) {
	acc, _ = c.Cookie("acc")
	if acc == "admin" {
		c.HTML(http.StatusOK, "bot1.html", gin.H{
			"message":     "欢迎！",
			"messagetype": "1",
			"uname":       uname,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":     "您的权限不足！",
			"messagetype": "4",
			"uname":       uname,
		})
	}
}
