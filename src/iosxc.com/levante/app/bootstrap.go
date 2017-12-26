package app

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
	"iosxc.com/levante/util"
	"os"
	"github.com/kataras/iris/sessions"
	"github.com/jinzhu/gorm"
)

//执行顺序依次如下
//config ->home ->logger -> db -> owner -> web -> run

func Run(app *iris.Application,cfgFile string) {
	config := initConfig(app, cfgFile)
	setHome(app, config)
	setLogger(app, config)
	setDatabase(app,config)
	setCtxHolder(app,config)
	setWebView(app, config)
	app.Run(iris.Addr(fmt.Sprintf("%s:%d", config.Server.Addr, config.Server.Port)), iris.WithCharset(config.Server.CharSet))
}

var ctxHolder *ContextHolder

type ContextHolder struct {
	SessionsManager *sessions.Sessions
	Database        *gorm.DB
	Config          *AppConfig
}


func setCtxHolder(app *iris.Application,config *AppConfig) {
	ctxHolder = &ContextHolder{}
	ctxHolder.Config = config
	ctxHolder.Database = db
	ctxHolder.SessionsManager = sessions.New(sessions.Config{Cookie: "mysessioncookie"})
}

func setHome(application *iris.Application, config *AppConfig) {
	if !util.CheckIsExistPath(config.Home) {
		panic("config.Home :" + config.Home + " is not exist!")
	}
}

func setLogger(application *iris.Application, config *AppConfig) {
	logPath := fmt.Sprintf("%s%s", config.Home, config.Log.File)
	file, err := os.Open(logPath)
	if err != nil {
		file, _ = os.Create(logPath)
	}
	defer file.Close()
	logger := NewRequestLogger(config)
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
	app.RegisterView(view.HTML(templatePath, config.View.Template.Ext).Layout(config.View.Template.Layout).Reload(config.View.Template.Reload))
}
