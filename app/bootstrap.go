package app

import (
	"fmt"
	"html/template"
	"os"

	"github.com/kataras/iris/v12"
	"gopkg.in/russross/blackfriday.v2"
	"iosxc.com/levante/util"
)

//logger -> db -> context -> web -> run
func Launch(app *iris.Application, config *AppConfig) {
	setLogger(app, config)
	setWebView(app, config)
	registerErrorHandler(app)
	app.Run(iris.Addr(fmt.Sprintf("%s:%d", config.Server.Addr, config.Server.Port)), iris.WithCharset(config.Server.CharSet))
}

func registerErrorHandler(application *iris.Application) {
	application.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().GetStringDefault("message", "网页丢啦"))
		ctx.View("front/error.html")
	})
}

func setLogger(application *iris.Application, config *AppConfig) {
	logPath := fmt.Sprintf("%s%s", config.Home, config.Log.File)
	file, err := os.Open(logPath)
	if err != nil {
		file, _ = os.Create(logPath)
	}
	defer file.Close()
	logger := NewRequestLogger(config)
	application.Logger().SetLevel(config.Log.Level)
	application.Use(logger)
}

func setWebView(app *iris.Application, config *AppConfig) {
	staticPath := fmt.Sprintf("%s%s", config.Home, config.View.Static.Path)
	htmlPath := fmt.Sprintf("%s%s", config.Home, config.View.HTML.Path)
	templatePath := fmt.Sprintf("%s%s", config.Home, config.View.Template.Path)

	if !util.CheckIsExistPath(staticPath) {
		panic("staticPath :" + staticPath + " is not exist!")
	}

	if !util.CheckIsExistPath(htmlPath) {
		panic("htmlPath :" + htmlPath + " is not exist!")
	}

	if !util.CheckIsExistPath(templatePath) {
		panic("templatePath :" + templatePath + " is not exist!")
	}
	app.HandleDir(config.View.Static.URI, staticPath)
	app.HandleDir(config.View.HTML.URI, htmlPath)
	templateView := iris.HTML(templatePath, config.View.Template.Ext).Layout(config.View.Template.Layout).Reload(config.View.Template.Reload)
	templateView.AddFunc("markdown", func(arg string) template.HTML {
		buf := blackfriday.Run([]byte(arg))
		return template.HTML(buf)
	})
	app.RegisterView(templateView)
}
