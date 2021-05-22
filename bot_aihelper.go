package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func aihelper(c *gin.Context) {

	c.HTML(http.StatusOK, "bot_ai.html", gin.H{})
}
