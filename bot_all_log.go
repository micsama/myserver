// Package main provides ...
package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm"
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
type td struct {
	Id   int
	Date string
	Time string
	Name string
	Num  int
	Area string
}
type ry struct {
	Id    int
	Date  string
	Time  string
	Name  string
	Phone int
	Skm   string
}

var timenow time.Time
var botname = [...]string{"自定位抗疫机器人", "抗疫空中四旋翼机器人", "消毒机器人", "药房自动发药机器人"}
var logname = [...]string{"spaerlog", "mwuprlog", "drlog", "adrlog"}
var mylog = ""
var body = ""

var bot int
var value string

func datelog2(c *gin.Context) {
	var tds []td
	var rys []ry
	body = ""
	bot := 1
	value1 := c.Query("value")
	if value1 != "-1" {
		value = value1
	}
	userdb = userdb.Table("rylog")
	if value != "" {
		date1 := value[:10]
		date2 := value[13:]
		fmt.Println(date1[:])
		fmt.Println(date2[:])
		userdb.Where("date >= ? and date <= ?", date1, date2).Find(&rys)
	} else {
		userdb.Find(&rys)
	}
	userdb = userdb.Table("tdlog")
	if value != "" {
		date1 := value[:10]
		date2 := value[13:]
		fmt.Println(date1[:])
		fmt.Println(date2[:])
		userdb.Where("date >= ? and date <= ?", date1, date2).Find(&tds)
	} else {
		userdb.Find(&tds)
	}
	body = "人员日志：\n"
	body = body + "日期 \t\t时间\t姓名\t联系方式\t\t苏康码状态\n"
	for _, p := range rys {
		body = body + p.Date + "\t" + p.Time + "\t" + p.Name + "\t" + strconv.Itoa(p.Phone) + "\t" + p.Skm + "\n"
	}
	body = body + "\n运送信息：\n日期 \t\t时间\t名称\t数量\t运送地区\n"
	for _, p := range tds {
		body = body + p.Date + "\t" + p.Time + "\t" + p.Name + "\t" + strconv.Itoa(p.Num) + "\t" + p.Area + "\n"
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

func datelog(c *gin.Context) {

	var logs []onelog
	body = ""
	value1 := c.Query("value")
	bot1 := c.Query("id")
	if bot1 == "1" {
		datelog2(c)
	} else {

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
}

func todaylog(i int) string {
	body = ""
	timenow = time.Now()
	date1 := timenow.Format("2006-01-02")
	fmt.Println(date1)
	if i == 1 {
		var tds []td
		var rys []ry
		userdb = userdb.Table("rylog")
		userdb.Where("date == ?", date1).Find(&rys)
		userdb = userdb.Table("tdlog")
		userdb.Where("date == ?", date1).Find(&tds)
		body = "人员日志：\n"
		body = body + "日期 \t\t时间\t姓名\t联系方式\t\t苏康码状态\n"
		for _, p := range rys {
			body = body + p.Date + "\t" + p.Time + "\t" + p.Name + "\t" + strconv.Itoa(p.Phone) + "\t" + p.Skm + "\n"
		}
		body = body + "\n运送信息：\n日期 \t\t时间\t名称\t数量\t运送地区\n"
		for _, p := range tds {
			body = body + p.Date + "\t" + p.Time + "\t" + p.Name + "\t" + strconv.Itoa(p.Num) + "\t" + p.Area + "\n"
		}
	} else {
		body = body + "今日日志："
		var logs []onelog
		mylog = logname[bot]
		userdb = userdb.Table(mylog)
		userdb.Where("date = ? ", date1).Find(&logs)
		for _, p := range logs {
			body = body + p.Log + "\n"
		}
	}
	return body
}
func (u onelog) TableName() string {
	fmt.Println("table--->" + mylog)
	return mylog
}
