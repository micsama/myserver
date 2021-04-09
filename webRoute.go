package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
)

func runweb() {
	wg.Add(1)
	r := gin.Default()
	r.Use(Checklogin())
	r.Static("/static", "./static")  //css
	r.Static("/layui", "./layui")    //css
	r.NoRoute(func(c *gin.Context) { //404
		c.HTML(http.StatusNotFound, "404.html", nil)
	})
	r.SetFuncMap(template.FuncMap{ //html  |safe
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	//	r.LoadHTMLGlob("./templates/**/*") //load templates
	r.HTMLRender = loadTemplates("./templates")

	//-----------------GET------------------
	r.GET("/index", func(c *gin.Context) {
		tag = "index.html"
		message = "欢迎来到抗疫机器人在线管理系统！"
		messagetype = "4"
		geth(c)
	})
	r.GET("/yfdata", yfdata)
	r.GET("/home", func(c *gin.Context) {
		tag = "home.html"
		message = "欢迎来到控制台！"
		messagetype = "1"
		geth(c)
	})
	r.GET("/", func(c *gin.Context) {
		tag = "index.html"
		message = "欢迎来到抗疫机器人在线管理系统！"
		messagetype = "1"
		geth(c)
	})
	r.GET("/register", func(c *gin.Context) {
		tag = "register.html"
		message = "欢迎注册！！"
		messagetype = "0"
		geth(c)
	})
	r.GET("/about", func(c *gin.Context) {
		tag = "about.html"
		message = "关于"
		messagetype = "0"
		geth(c)
	})
	r.GET("/shijian", func(c *gin.Context) {
		tag = "shijian.html"
		message = "了解一下新冠的时间线吧~"
		messagetype = "0"
		geth(c)
	})
	r.GET("/chengshi", chengshi)
	r.GET("/dashuju", dashuju)
	r.GET("/xdlog", xdlog)
	r.GET("/yfmanager", yfmanager)
	r.GET("/logout", logout)
	r.GET("/signin", func(c *gin.Context) {
		tag = "signin.html"
		message = "请登录账户！"
		messagetype = "0"
		geth(c)
	})

	//--------------------POST------------------
	r.POST("/register", register)
	r.POST("/signin", signin)

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
	c.HTML(http.StatusOK, tag, gin.H{
		"message":     message,
		"uname":       uname,
		"messagetype": messagetype,
	})
}
func Checklogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		uname, _ = c.Cookie("name")
		if uname != "" {
			uname = " " + uname + " "
		} else {
			message = "请登录账户！"
			messagetype = "4"
		}
		acc = "admin"
		c.Next()
	}
}

func xdlog(c *gin.Context) {
	acc, _ = c.Cookie("acc")
	if acc == "admin" {
		c.HTML(http.StatusOK, "xdlog.html", gin.H{
			"message":     "欢迎！",
			"messagetype": "1",
			"uname":       uname,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":     "您的权限不足！正在返回首页！",
			"messagetype": "4",
			"uname":       uname,
		})
	}
}

func yfdata(c *gin.Context) {
	acc, _ = c.Cookie("acc")
	if acc == "admin" {
		refreshyfdata()
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"msg":   "",
			"count": yfcount,
			"data":  yfdatamap,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":     "您的权限不足！正在返回首页！",
			"messagetype": "4",
			"uname":       uname,
		})
	}

}
func yfmanager(c *gin.Context) {
	acc, _ = c.Cookie("acc")
	if acc == "admin" {
		c.HTML(http.StatusOK, "yfmanager.html", gin.H{
			"message":     "欢迎！",
			"messagetype": "1",
			"uname":       uname,
		})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message":     "您的权限不足！正在返回首页！",
			"messagetype": "4",
			"uname":       uname,
		})
	}
}
func dashuju(c *gin.Context) {
	a := initjson()
	c.HTML(http.StatusOK, "dashuju.html", gin.H{
		"message":     "今日速报！",
		"messagetype": "0",
		"uname":       uname,
		"data":        a.Data,
		"gntotal":     a.Gntotal,
		"deathtotal":  a.Deathtotal,
		"curetotal":   a.Curetotal,
	})
}
func chengshi(c *gin.Context) {
	c.HTML(http.StatusOK, "chengshi.html", gin.H{
		"message":     "欢迎！",
		"messagetype": "1",
		"uname":       uname,
	})
}
