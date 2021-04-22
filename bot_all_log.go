// Package main provides ...
package main

import (
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

const (
	logpath = "./data/log/"
)

type onelog struct {
	Id   int    `form:"id"`
	Date string `form:"date"`
	Time string `form:"time"`
	Log  string `form:"log"`
}

var f = [...]string{"Self_positioning_anti_epidemic_robot.log", "Multi_wing_UAV_patrol_robot.log", "Disinfection_robot.log", "Automatic_dispensing_robot.log"}
var botname = [...]string{"自定位抗疫机器人", "多翼无人机抗疫巡检机器人", "消毒机器人", "药房自动发药机器人"}
var logname = [...]string{"spaerlog", "mwuprlog", "drlog", "adrlog"}
var mylog = ""
var body = ""

// func getlog(i int) *string {
// 	logname1 := f[i]
// 	file, _ := os.Open(logpath + logname1)
// 	fmt.Println(logpath + logname1)
// 	body_byte, _ := io.ReadAll(file)
// 	body1 := (*string)(unsafe.Pointer(&body_byte)) //body的string直接调用[]byte的内存，防止从[]byte转换到string的日志一整份拷贝。可以节省内存空间
// 	return body1
// }
//
// func log(c *gin.Context) {
// 	bot, _ := strconv.Atoi(c.Query("id"))
// 	body := getlog(bot)
// 	c.HTML(http.StatusOK, "bot_log.html", gin.H{
// 		"message":     "",
// 		"messagetype": "",
// 		"uname":       uname,
// 		"botname":     botname[bot],
// 		"bot":         bot,
// 		"body":        *body,
// 	})
// }
var bot int
var value string

func datelog(c *gin.Context) {
	var logs []onelog
	body = ""
	value1 := c.Query("value")
	bot1 := c.Query("id")
	if bot1 != "" {
		bot, _ = strconv.Atoi(bot1)
	}
	if value1 != "-1" {
		value = value1
	}
	mylog = logname[bot]
	userdb = userdb.Table(mylog)
	if value != "" {
		date1 := value[:10]
		date2 := value[13:]
		fmt.Println(date1[:])
		fmt.Println(date2[:])
		userdb.Where("date >= ? and date <= ?", date1, date2).Find(&logs)
	} else {
		userdb.Find(&logs)
	}
	fmt.Println(logs)
	body = ""
	for _, p := range logs {
		body = body + p.Log + "\n"
	}
	c.HTML(http.StatusOK, "bot_log.html", gin.H{
		"message":     "",
		"value":       value,
		"messagetype": "",
		"uname":       uname,
		"bot":         bot,
		"botname":     botname[bot],
		"body":        body,
	})
}

func (u onelog) TableName() string {
	fmt.Println("table--->" + mylog)
	return mylog
}
