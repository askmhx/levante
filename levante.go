package main

import (
	"flag"
	"iosxc.com/levante/app"
	"github.com/kataras/iris"
	"iosxc.com/levante/ctrl"
	"fmt"
	"runtime"
	"time"
)

const url_index = "/"
const url_start = "/start"
const url_archive = "/archive"
const url_about = "/about"
const url_photo = "/photo"
const url_post = "/post/{pid:int}"

func initRoute(application *iris.Application) {
	registerErrorHandler(application)
	indexCtrl := ctrl.IndexCtrl{}
	photoCtrl := ctrl.PhotoCtrl{}
	postCtrl := ctrl.PostCtrl{}
	application.Get(url_index, app.Handler(indexCtrl.IndexHandle))
	application.Get(url_start, app.Handler(indexCtrl.StartHandle))
	application.Get(url_archive, app.Handler(indexCtrl.ArchiveHandle))
	application.Get(url_about, app.Handler(indexCtrl.AboutHandle))
	application.Get(url_photo, app.Handler(photoCtrl.IndexHandle))
	application.Get(url_post, app.Handler(postCtrl.ReadHandle))
}

func registerErrorHandler(application *iris.Application) {
	application.OnErrorCode(iris.StatusNotFound, app.Handler(func(ctx *app.Context) {
		ctx.View("front/404.html")
	}))
	application.OnErrorCode(iris.StatusInternalServerError, app.Handler(func(ctx *app.Context) {
		ctx.View("front/500.html")
	}))
}

var (
	AppBanner  = "Levante %s Date: %s Build: %s"
	AppVersion = "1.0.0"
	AppDate    = time.Now().Format("2006-01-02 15:04:05")
	GoVersion  = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

func main() {

	fmt.Println(fmt.Sprintf(AppBanner,AppVersion, AppDate, GoVersion))

	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "NOT SET", "app cfg path")
	flag.Parse()
	application := iris.New()
	initRoute(application)
	app.ConfigAndStart(application, cfgPath)
}
