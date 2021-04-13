// Package main provides ...
package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"os"
	"os/exec"
)

type Citydata struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Deathnum string `json:"deathNum"`
	Curenum  string `json:"cureNum"`
	Zerodays string `json:"zerodays"`
}
type Sinajson struct {
	Data       string
	Jwsr       string
	Wzz        string
	Gntotal    string
	Deathtotal string
	Curetotal  string
	Econnum    string
	Myvalue    string
	Mydeathnum string
	Mycurenum  string
	Myzerodays string
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary
var Body string

func getjson() {
	s := "rm fymap2020_data.d.json.1;rm fymap2020_data.d.json;wget https://interface.sina.cn/news/wap/fymap2020_data.d.json"
	cmd := exec.Command("/bin/bash", "-c", s)
	out, err := cmd.Output()
	fmt.Println(out, err)
	exec.Command("/bin/bash", "-c", "echo "+string(out[:])+"> 1.log")
}

func initjson() Sinajson {
	// url := "https://interface<0;21;8M.sina.cn/news/wap/fymap2020_data.d.json"
	// resp, _ := http.Get(url)
	file, _ := os.Open("fymap2020_data.d.json")
	Body, _ := io.ReadAll(file)
	any := jsoniter.Get(Body, "data", "jwsrNum")
	var dat Sinajson
	dat.Jwsr = any.ToString()
	any = jsoniter.Get(Body, "data", "asymptomNum")
	dat.Wzz = any.ToString()
	any = jsoniter.Get(Body, "data", "deathtotal")
	dat.Deathtotal = any.ToString()
	any = jsoniter.Get(Body, "data", "times")
	dat.Data = any.ToString()
	any = jsoniter.Get(Body, "data", "gntotal")
	dat.Gntotal = any.ToString()
	any = jsoniter.Get(Body, "data", "curetotal")
	dat.Curetotal = any.ToString()
	any = jsoniter.Get(Body, "data", "econNum")
	dat.Econnum = any.ToString()

	var p Citydata
	for i := 0; i < 36; i++ {
		City := jsoniter.Get(Body, "data", "list", i, "name")
		if City.ToString() == Thecity {
			any1 := jsoniter.Get(Body, "data", "list", i)
			any1.ToVal(&p)
			dat.Mycurenum = p.Curenum
			dat.Mydeathnum = p.Deathnum
			dat.Myvalue = p.Value
			dat.Myzerodays = p.Zerodays
		}

	}
	file.Close()
	return dat
}
