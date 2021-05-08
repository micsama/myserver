// Package main provides ...
package main

import (
	_ "fmt"
)

var statusmap [4]map[string]string

func status(i int) string {
	var s string
	for i, c := range statusmap[i] {
		s = s + i + c + "\n\n"
	}
	return s
}
func initstatus() {
	statusmap[0] = make(map[string]string)
	statusmap[1] = make(map[string]string)
	statusmap[2] = make(map[string]string)
	statusmap[3] = make(map[string]string)
	statusmap[0]["充电状态 \t:"] = "80%"
	statusmap[0]["位置   \t:"] = "（233：123）"
	statusmap[0]["病房环境 \t:"] = "良好"
	statusmap[0]["病人状况 \t:"] = "良好"
	statusmap[1]["充电状态 \t:"] = "68%"
	statusmap[1]["位置   \t:"] = "（333：13）"
	statusmap[1]["飞机状态 \t:"] = "良好"
	statusmap[1]["环境状态 \t:"] = "良好"
	statusmap[2]["充电状态 \t:"] = "20%"
	statusmap[2]["位置   \t:"] = "(124,36)"
	statusmap[2]["消毒进度 \t:"] = "80%"
	statusmap[2]["病人状况 \t:"] = "良好"
	statusmap[3]["充电状态 \t:"] = "90%"
	statusmap[3]["位置   \t:"] = "（233：123）"
	statusmap[3]["药房情况"] = sellout
}
