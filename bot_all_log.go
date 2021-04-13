// Package main provides ...
package main

import (
	"fmt"
	_ "fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"unsafe"

	"github.com/gin-gonic/gin"
)

const (
	logpath = "./data/log/"
)

var f = [...]string{"Self_positioning_anti_epidemic_robot.log", "Multi_wing_UAV_patrol_robot.log", "Disinfection_robot.log", "Automatic_dispensing_robot.log"}
var botname = [...]string{"自定位抗疫机器人", "多翼无人机抗议巡检机器人", "消毒机器人", "药房自动发药机器人"}

func log(c *gin.Context) {
	bot, _ := strconv.Atoi(c.Query("id"))
	logname := f[bot]
	file, _ := os.Open(logpath + logname)
	fmt.Println(logpath + logname)
	body_byte, _ := io.ReadAll(file)
	body := (*string)(unsafe.Pointer(&body_byte)) //body的string直接调用[]byte的内存，防止从[]byte转换到string的日志一整份拷贝。可以节省内存空间
	c.HTML(http.StatusOK, "bot_log.html", gin.H{
		"message":     "",
		"messagetype": "",
		"uname":       uname,
		"botname":     botname[bot],
		"bot":         bot,
		"body":        *body,
	})

}
