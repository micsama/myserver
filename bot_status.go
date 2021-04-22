// Package main provides ...
package main

import (
	_ "fmt"
)

var statusmap [4]map[string]string

func initstatus() {
	statusmap[0] = make(map[string]string)
	statusmap[1] = make(map[string]string)
	statusmap[2] = make(map[string]string)
	statusmap[3] = make(map[string]string)
	statusmap[0]["\n充电状态 \t"] = "80%"
	statusmap[0]["\n位置   \t"] = "（233：123）"
	statusmap[0]["\n病房环境 \t"] = "良好"
	statusmap[0]["\n病人状况 \t"] = "良好"
	statusmap[1]["\n充电状态 \t"] = "68%"
	statusmap[1]["\n位置   \t"] = "（333：13）"
	statusmap[1]["\n飞机状态 \t"] = "良好"
	statusmap[1]["\n环境状态 \t"] = "良好"
	statusmap[2]["\n充电状态 \t"] = "20%"
	statusmap[2]["\n位置   \t"] = "(124,36)"
	statusmap[2]["\n消毒进度 \t"] = "80%"
	statusmap[2]["\n病人状况 \t"] = "良好"
	statusmap[3]["\n充电状态 \t"] = "90%"
	statusmap[3]["\n位置   \t"] = "（233：123）"
	statusmap[3]["\n药房情况"] = sellout
}
