package app

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/view"
	"iosxc.com/levante/util"
	"os"
)

var application *iris.Application

func Run(cfgFile string) {
	application = iris.New()
	config := initConfig(application, cfgFile)
	setHome(application, config)
	setLogger(application, config)
	setRoute(application, config)
	setWebView(application, config)
	application.Run(iris.Addr(fmt.Sprintf("%s:%d", config.Server.Addr, config.Server.Port)), iris.WithCharset(config.Server.CharSet))
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

func setRoute(application *iris.Application, config *AppConfig) {
	initRoute(application, config, initDatabase(config))
}

func setWebView(application *iris.Application, config *AppConfig) {
	staticPath := fmt.Sprintf("%s%s", config.Home, config.View.Static.Path)
	templatePath := fmt.Sprintf("%s%s", config.Home, config.View.Template.Path)

	if !util.CheckIsExistPath(staticPath) {
		panic("staticPath :" + staticPath + " is not exist!")
	}

	if !util.CheckIsExistPath(templatePath) {
		panic("templatePath :" + templatePath + " is not exist!")
	}
	application.StaticWeb(config.View.Static.URI, staticPath)
	application.RegisterView(view.HTML(templatePath, config.View.Template.Ext).Layout(config.View.Template.Layout).Reload(config.View.Template.Reload))
}
