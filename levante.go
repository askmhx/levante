package main

import (
	"flag"
	"iosxc.com/levante/app"
	"github.com/kataras/iris"
	"iosxc.com/levante/ctrl"
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
	application.Get(url_archive,  app.Handler(indexCtrl.ArchiveHandle))
	application.Get(url_about,  app.Handler(indexCtrl.AboutHandle))
	application.Get(url_photo,  app.Handler(photoCtrl.IndexHandle))
	application.Get(url_post,  app.Handler(postCtrl.ReadHandle))
}

func registerErrorHandler(application *iris.Application) {
	application.OnErrorCode(iris.StatusNotFound,  app.Handler(func(ctx *app.Context) {
		ctx.View("front/404.html")
	}))
	application.OnErrorCode(iris.StatusInternalServerError,  app.Handler(func(ctx *app.Context) {
		ctx.View("front/500.html")
	}))
}


func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "cfg", "NOT SET", "app cfg path")
	flag.Parse()
	application := iris.New()
	initRoute(application)
	app.ConfigAndStart(application, cfgPath)
}
