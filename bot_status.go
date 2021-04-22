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
	statusmap[0]["\n充电状态"] = "80%"
	statusmap[0]["\n位置"] = "（233：123）"
	statusmap[0]["\n病房环境"] = "良好"
	statusmap[0]["\n病人身体状况"] = "良好"
	statusmap[1]["\n充电状态"] = "68%"
	statusmap[1]["\n位置"] = "（333：13）"
	statusmap[1]["\n飞机状态"] = "良好"
	statusmap[1]["\n未佩戴口罩人数"] = "良好"
	statusmap[2]["\n充电状态"] = "20%"
	statusmap[2]["\n位置"] = "(124,36)"
	statusmap[2]["\n今日消毒进度"] = "80%"
	statusmap[2]["\n病人身体状况"] = "良好"
	statusmap[3]["\n充电状态"] = "90%"
	statusmap[3]["\n位置"] = "（233：123）"
	statusmap[3]["\n药房情况"] = sellout
}
