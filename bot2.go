// Package main provides ...
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func bot2(c *gin.Context) {
	acc, _ = c.Cookie("acc")
	if acc == "admin" {
		c.HTML(http.StatusOK, "bot2.html", gin.H{
			"uname": uname,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "您的权限不足！",
			"uname":   uname,
		})
	}
}
