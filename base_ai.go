package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
)

func aiask(s string) string {
	str := "{\"reqType\":0,\"perception\":{\"inputText\":{\"text\":\"" + s + "\"},\"inputImage\":{\"url\":\"imageUrl\"},\"selfInfo\":{\"location\":{\"city\":\"南京\",\"province\":\"南京\"}}},\"userInfo\":{\"apiKey\":\"e5239674321641abb414259c406aaa5d\",\"userId\":\"1\"}}"
	postbuff := []byte(str)
	req, _ := http.Post("http://openapi.tuling123.com/openapi/api/v2", "JSON", bytes.NewBuffer(postbuff))
	body, _ := ioutil.ReadAll(req.Body)
	return string(body)
}
func aiaskroute(c *gin.Context) {
	str := c.PostForm("ask")
	ans := aiask(str)
	any := jsoniter.Get([]byte(ans), "results", 0, "values", "text").ToString()
	c.JSON(200, gin.H{
		"s": any,
	})
}
