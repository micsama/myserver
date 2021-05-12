// Package main provides ...
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	jsoniter "github.com/json-iterator/go"
)

const tqurl = "http://www.weather.com.cn/data/cityinfo/101190101.html"

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
func tianqi() string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(tqurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	tqbody := result.String()
	tq := "南京 气温在" + jsoniter.Get([]byte(tqbody), "weatherinfo", "temp1").ToString() + "~" + jsoniter.Get([]byte(tqbody), "weatherinfo", "temp2").ToString() + " ,天气：" + jsoniter.Get([]byte(tqbody), "weatherinfo", "weather").ToString()
	return tq
}
