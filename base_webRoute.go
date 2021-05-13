package main

import (
	"crypto/sha256"
	"fmt"
	_ "fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var Words []rune
var messagenum = "ok"
var Thecity = "北京"

func runweb() {
	// gin.SetMode(gin.ReleaseMode)
	wg.Add(1)
	r := gin.Default()               //定义默认路由
	r.Use(Checklogin())              //定义中间件，每次路由都会使用
	r.Static("/static", "./static")  //css+picture
	r.Static("/layui", "./layui")    //css
	r.NoRoute(func(c *gin.Context) { //404路由
		c.HTML(http.StatusNotFound, "base_404.html", nil)
	})
	r.SetFuncMap(template.FuncMap{ //html  |safe
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//	r.LoadHTMLGlob("./templates/**/*") //load templates
	r.HTMLRender = loadTemplates("./templates") //加载并准备渲染模板
	//-----------------GET------------------
	r.GET("/index", func(c *gin.Context) {
		tag = "base_index.html"
		message = "欢迎来到抗疫机器人在线管理系统！"
		messagetype = "4"
		geth(c)
	})
	r.GET("/addnewok", func(c *gin.Context) {
		tag = "bot_yaofang_addnew_ok.html"
		geth(c)
	})

	r.GET("/addnew", func(c *gin.Context) {
		tag = "bot_yaofang_addnew.html"
		geth(c)
	})
	r.GET("/", func(c *gin.Context) {
		if c.Query("id") == "-1" {
			logout(c)
		} else {
			tag = "base_index.html"
			message = "欢迎来到抗疫机器人在线管理系统！"
			messagetype = "1"
			geth(c)
		}
	})
	r.GET("/register", func(c *gin.Context) {
		tag = "user_register.html"
		message = "欢迎注册！！"
		messagetype = "0"
		geth(c)
	})
	r.GET("/about", func(c *gin.Context) {
		tag = "base_about.html"
		message = "关于"
		messagetype = "0"
		geth(c)
	})
	r.GET("/monitor", func(c *gin.Context) {
		tag = "bot_monitor.html"
		message = ""
		messagetype = ""
		geth(c)
	})
	r.GET("/signin", func(c *gin.Context) {
		tag = "user_signin.html"
		message = "请登录账户！"
		messagetype = "0"
		geth(c)
	})
	r.GET("/shijian", func(c *gin.Context) {
		tag = "base_shijian.html"
		message = "了解一下新冠的时间线吧~"
		messagetype = "0"
		geth(c)
	})

	r.GET("/Automatic_dispensing", func(c *gin.Context) {
		refreshyfdata()
		get_control_panel(c, 3)

	})
	r.GET("/Disindection", func(c *gin.Context) {
		get_control_panel(c, 2)
	})
	r.GET("/Muiti_win_UAV_patrol", func(c *gin.Context) {
		get_control_panel(c, 1)
	})
	r.GET("/Self_positioning_anti_epidemic", func(c *gin.Context) {
		get_control_panel(c, 0)
	})

	r.GET("/home", home)
	r.GET("/connect", connect)
	r.GET("/dashujure", dashujure)
	r.GET("/dashuju", dashuju)
	r.GET("/xdlog", xdlog)
	r.GET("/yfdata", yfdata)
	r.GET("/yfmanager", yfmanager)
	r.GET("/updatayf", updatayf)
	r.GET("/log", datelog)
	r.GET("/SendSms", Rsend)
	//--------------------POST------------------
	r.POST("/log", datelog)
	r.POST("/dashuju", mycity)
	r.POST("/addnew", addnew)
	r.POST("/register", register)
	r.POST("/home", signin)
	//---------Run---------------------/
	r.Run(":13488")
	wg.Done()

}
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(includes)
	// 为layouts/和includes/目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

func geth(c *gin.Context) {
	fmt.Println(tag)
	c.HTML(http.StatusOK, tag, gin.H{
		"message":     message,
		"uname":       uname,
		"messagetype": messagetype,
		"messagenum":  messagenum,
	})
}
func Checklogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		uname, _ = c.Cookie("name")
		acc, _ = c.Cookie("acc")
		a := sha256.Sum256([]byte(uname + "admin"))
		accsha := string(a[:])
		if acc == accsha {
		} else {
			uname = ""
			message = "请登录账户！"
			messagetype = "4"
		}
	}
}

func xdlog(c *gin.Context) {
	c.HTML(http.StatusOK, "xdlog.html", gin.H{
		"message":     "欢迎！",
		"messagetype": "1",
		"uname":       uname,
	})
}

func yfdata(c *gin.Context) {
	acc, _ = c.Cookie("acc")
	refreshyfdata()
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "",
		"count": yfcount,
		"data":  yfdatamap,
	})

}
func yfmanager(c *gin.Context) {
	acc, _ = c.Cookie("acc")
	c.HTML(http.StatusOK, "bot_yaofang.html", gin.H{
		"message":     "欢迎！",
		"messagetype": "1",
		"uname":       uname,
		"sellout":     sellout,
	})
}
func dashujure(c *gin.Context) {
	getjson()
	dashuju(c)

}
func dashuju(c *gin.Context) {
	a := initjson()
	c.HTML(http.StatusOK, "base_dashuju.html", gin.H{
		"message":     "今日速报！",
		"messagetype": "0",
		"uname":       uname,
		"data":        a.Data,
		"gntotal":     a.Gntotal,
		"deathtotal":  a.Deathtotal,
		"curetotal":   a.Curetotal,
		"jwsrNum":     a.Jwsr,
		"wzz":         a.Wzz,
		"econNum":     a.Econnum,
		"mycity":      Thecity,
		"myvalue":     a.Myvalue,
		"mydeathnum":  a.Mydeathnum,
		"mycurenum":   a.Mycurenum,
		"myzerodays":  a.Myzerodays,
	})
}

func mycity(c *gin.Context) {
	Thecity = c.Query("mycity")
	Words = ([]rune)(Thecity)
	fmt.Println(Thecity, "-----")
	for i, c := range Words {
		fmt.Println("--> ", c)
		if c == 30465 || c == 24066 { //江苏“省”南京“市”的“省”和“市”字的编码
			Words = Words[:i]
			break
		}
	}
	fmt.Println(Words, "-----")
	Thecity = string(Words)
}
func get_control_panel(c *gin.Context, i int) {
	initstatus()
	// s, _ := json.Marshal(statusmap[i])

	c.HTML(http.StatusOK, "bot_control_panel.html", gin.H{
		"message":     message,
		"uname":       uname,
		"messagetype": messagetype,
		"messagenum":  messagenum,
		"bot":         i,
		"botname":     botname[i],
		"status":      status(i),
		"body":        todaylog(i),
	})
}
func home(c *gin.Context) {
	refreshyfdata()
	initstatus()
	c.HTML(http.StatusOK, "user_home.html", gin.H{
		"message":     "欢迎来到控制台",
		"messagetype": "1",
		"uname":       uname,
		"tianqi":      tianqi(),
		"sellout":     sellout,
	})

}
func connect(c *gin.Context) {
	bot1 := c.Query("bot")
	ip := c.Query("ip")
	connectip(bot1, ip)

}
