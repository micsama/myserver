// Package main provides ...
package main

import (
	_ "fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
func bot_info(c *gin.Context) {
	info := [4]string{"本模块化设计的自定位抗疫机器人可实现以下功能：\n1、机器人能够通过LABVIEW编写的上位机系统实现对病房环境数据的监测以及响应病人的呼叫等功能。\n2、该机器人能够通过电机以及避障模块自主运动并准确无误的到达指定位置，并且能够通过加装的升降台以及舵机代替医生护士进行查房送药。\n3、结合摄像头、OPENMV红外测温等功能的辅助抗疫功能，极大减轻医护人员的负担，减少因与病患接触而导致的病毒感染。", "抗议空中四旋翼机器人设计由四旋翼无人机运输系统和物资管理系统两部分组成，可以广泛用于方舱医院物资运输。四旋翼无人机运输系统可以快速按照上位机指令把物资送达目的地，减少物资运输时间。四旋翼无人机运输系统上电后，MPU6050模块获取四旋翼无人机姿态数据，并通过I2C协议传输给主控制器。气压计模块把测的数据传给主控制器，经换算得到四旋翼无人机实时高度，用于实现四旋翼无人机定高飞行。光流模块用于反馈四旋翼无人机位置信息，用于四旋翼无人机把物资送到上位机指定目的地。本设计通过WIFI模块把物资名称、运输时间、数量等信息无线传输给上位机，上位机可对物资信息进行分析、记录。物资管理人员也可通过上位机物资管理系统选择四旋翼无人机物资送达的目的地。", "基于室内导航的消毒防疫机器人由导航定位系统，雾化消毒系统，视日轨迹跟踪供电系统组成。导航定位系统能够使机器人快速适应不同的室内环境，通过激光雷达与惯性导航单元等获取的数据进行较为精准的建图，定位与导航等功能。雾化消毒系统能够将消毒水雾化喷洒，高效消毒。视日轨迹跟踪供电系统根据视日轨迹跟踪算法，通过北斗模块与时钟芯片获取设备所在位置的经纬度与时间，计算出设备所在位置的太阳高度角和方位角，驱动电机实现全天候太阳能板自动朝向太阳，提高太阳能电池板的电能转化效率，对设备进行供电储能。", "该系统主要由发药药柜和补药机器人两部分构成。其中发药药柜由STM32F1主控、OPENMV机器视觉模块、舵机模块、通讯模块构成；补药机器人主要由机械臂模块、无线通讯模块、运动模块、电源模块等构成。\n自动化药柜可以通过OPENMV机器视觉模块自动识别处方二维码，然后将相关信息上传至Labview上位机中并且自动化药柜自动完成出药过程，当药柜柜药物缺少时，可通过机器人进行补药。"}
	i, _ := c.GetQuery("id")
	bot, _ := strconv.Atoi(i)
	c.HTML(http.StatusOK, "bot_information.html", gin.H{
		"botimg":      botimg[bot],
		"botname":     botname[bot],
		"body":        info[bot],
		"message":     "今日速报！",
		"messagetype": "0",
		"uname":       uname,
	})

}
