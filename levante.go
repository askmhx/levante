package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris"
	"iosxc.com/levante/app"
	"runtime"
	"time"
)

var (
	AppBanner  = "Levante %s Date: %s Build: %s"
	AppVersion = "1.0.1"
	AppDate    = time.Now().Format("2006-01-02 15:04:05")
	GoVersion  = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

func main() {
	fmt.Println(fmt.Sprintf(AppBanner, AppVersion, AppDate, GoVersion))
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "NOT SET", "app cfg path")
	flag.Parse()
	application := iris.New()
	var config = app.InitConfig(application, cfgPath)
	var database = app.InitDatabase(config)
	app.InitRoute(application, config, database)
	app.Launch(application, config)
}
