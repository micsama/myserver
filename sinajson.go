// Package main provides ...
package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type Sinajson struct {
	Data string
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func initjson() {
	url := "https://interface.sina.cn/news/wap/fymap2020_data.d.json"
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
}
