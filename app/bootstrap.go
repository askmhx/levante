package app

import (
	"fmt"
	"github.com/kataras/iris"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"iosxc.com/levante/util"
	"os"
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
	staticPath := fmt.Sprintf("%s%s", config.Home, config.View.Statics.Path)
	htmlPath := fmt.Sprintf("%s%s", config.Home, config.View.Htmls.Path)
	templatePath := fmt.Sprintf("%s%s", config.Home, config.View.Templates.Path)

	if !util.CheckIsExistPath(staticPath) {
		panic("staticPath :" + staticPath + " is not exist!")
	}

	if !util.CheckIsExistPath(htmlPath) {
		panic("htmlPath :" + htmlPath + " is not exist!")
	}

	if !util.CheckIsExistPath(templatePath) {
		panic("templatePath :" + templatePath + " is not exist!")
	}
	app.StaticWeb(config.View.Statics.URI, staticPath)
	app.StaticWeb(config.View.Htmls.URI, htmlPath)
	templateView := iris.HTML(templatePath, config.View.Templates.Ext).Layout(config.View.Templates.Layout).Reload(config.View.Templates.Reload)
	templateView.AddFunc("markdown",func(arg string) template.HTML {
			buf := blackfriday.Run([]byte(arg))
			return template.HTML(buf)
	})
	app.RegisterView(templateView)
}

