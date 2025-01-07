package app

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"levante/util"
	"os"
)

// logger -> db -> context -> web -> run
func Launch(app *iris.Application, config *AppConfig) {
	setApplicationLogger(app, config)
	setRequestLogger(app, config)
	setWebView(app, config)
	registerErrorHandler(app)
	err := app.Run(iris.Addr(fmt.Sprintf("%s:%d", config.Server.Addr, config.Server.Port)), iris.WithCharset(config.Server.CharSet))
	if err != nil {
		panic(err)
	}
}

func registerErrorHandler(application *iris.Application) {
	application.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Message", ctx.Values().GetStringDefault("message", "网页丢啦"))
		err := ctx.View("front/error.html")
		if err != nil {
			panic(err)
		}
	})
}

func setApplicationLogger(application *iris.Application, config *AppConfig) {
	logPath := fmt.Sprintf("%s%s", config.Home, config.ApplicationLog.File)
	file, err := os.Open(logPath)
	if err != nil {
		file, _ = os.Create(logPath)
	}
	defer file.Close()
	application.Logger().SetLevel(config.ApplicationLog.Level)
	application.Logger().SetOutput(file)
}

func setRequestLogger(application *iris.Application, config *AppConfig) {
	logPath := fmt.Sprintf("%s%s", config.Home, config.AccessLog.File)
	ac := accesslog.File(logPath)
	ac.AddOutput(os.Stdout)
	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler
	ac.SetFormatter(&accesslog.JSON{
		Indent:    "  ",
		HumanTime: true,
	})
	application.UseRouter(ac.Handler)
	defer ac.Close()
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
	app.HandleDir(config.View.Statics.URI, staticPath)
	app.HandleDir(config.View.Htmls.URI, htmlPath)
	templateView := iris.HTML(templatePath, config.View.Templates.Ext).Layout(config.View.Templates.Layout).Reload(config.View.Templates.Reload)
	app.RegisterView(templateView)
}
