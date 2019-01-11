package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/view"
	"gopkg.in/russross/blackfriday.v2"
	"html/template"
	"iosxc.com/levante/util"
	"os"
)

//logger -> db -> context -> web -> run

func Launch(app *iris.Application, config *AppConfig) {
	setLogger(app, config)
	setDatabase(app, config)
	setCtxHolder(app, config)
	setWebView(app, config)
	app.Run(iris.Addr(fmt.Sprintf("%s:%d", config.Server.Addr, config.Server.Port)), iris.WithCharset(config.Server.CharSet))
}

var ctxHolder *ContextHolder

type ContextHolder struct {
	SessionsManager *sessions.Sessions
	Database        *gorm.DB
	Config          *AppConfig
}

func setCtxHolder(app *iris.Application, config *AppConfig) {
	ctxHolder = &ContextHolder{}
	ctxHolder.Config = config
	ctxHolder.Database = db
	ctxHolder.SessionsManager = sessions.New(sessions.Config{Cookie: "mysessioncookie"})
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

func setDatabase(app *iris.Application, config *AppConfig) {
	initDatabase(config)
}

func setWebView(app *iris.Application, config *AppConfig) {
	staticPath := fmt.Sprintf("%s%s", config.Home, config.View.Static.Path)
	templatePath := fmt.Sprintf("%s%s", config.Home, config.View.Template.Path)

	if !util.CheckIsExistPath(staticPath) {
		panic("staticPath :" + staticPath + " is not exist!")
	}

	if !util.CheckIsExistPath(templatePath) {
		panic("templatePath :" + templatePath + " is not exist!")
	}
	app.StaticWeb(config.View.Static.URI, staticPath)
	templateView := view.HTML(templatePath, config.View.Template.Ext).Layout(config.View.Template.Layout).Reload(config.View.Template.Reload)
	templateView.AddFunc("markdown",func(arg string) template.HTML {
			buf := blackfriday.Run([]byte(arg))
			return template.HTML(buf)
	})
	app.RegisterView(templateView)
}

