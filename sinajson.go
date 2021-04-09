// Package main provides ...
package main

import (
	jsoniter "github.com/json-iterator/go"
	"io"
	"os"
)

type Sinajson struct {
	Data       string
	Gntotal    string
	Deathtotal string
	Curetotal  string
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var Body string

func getjson() {
	s := "rm fymap2020_data.d.jso;wget https://interface.sina.cn/news/wap/fymap2020_data.d.json"
	cmd := exec.Command("/bin/bash", "-c", s)
}

func initjson() Sinajson {
	// url := "https://interface<0;21;8M.sina.cn/news/wap/fymap2020_data.d.json"
	// resp, _ := http.Get(url)
	file, _ := os.Open("fymap2020_data.d.json")
	Body, _ := io.ReadAll(file)
	any := jsoniter.Get(Body, "data", "times")
	var dat Sinajson
	dat.Data = any.ToString()
	any = jsoniter.Get(Body, "data", "gntotal")
	dat.Gntotal = any.ToString()
	any = jsoniter.Get(Body, "data", "deathtotal")
	dat.Deathtotal = any.ToString()
	any = jsoniter.Get(Body, "data", "curetotal")
	dat.Curetotal = any.ToString()
	file.Close()
	return dat
}
