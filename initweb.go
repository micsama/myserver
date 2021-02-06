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
		geth(c)
	})
	r.GET("/home", func(c *gin.Context) {
		tag = "home.html"
		geth(c)
	})
	r.GET("/", func(c *gin.Context) {
		tag = "index.html"
		geth(c)
	})
	r.GET("/register", func(c *gin.Context) {
		tag = "register.html"
		geth(c)
	})
	r.GET("/bot2", bot1)
	r.GET("/bot1", bot2)
	r.GET("/logout", logout)
	r.GET("/signin", func(c *gin.Context) {
		tag = "signin.html"
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
		"message": "hi",
		"uname":   uname,
	})
}
func Checklogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		uname, _ = c.Cookie("name")
		if uname != "" {
			uname = "欢迎：" + uname + " 点击进入控制台"
		}
		acc = "admin"
		c.Next()
	}
}
