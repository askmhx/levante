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
	wpApiInitRoute(application)

}

//WordPress v2 restful api urls

const url_wp_base_url = "/wp/v2"
const url_wp_api_posts = "/posts"
const url_wp_api_revisions = "/revisions"
const url_wp_api_categories = "/categories"
const url_wp_api_tags = "/tags"
const url_wp_api_pages = "/pages"
const url_wp_api_comments = "/comments"
const url_wp_api_taxonomies = "/taxonomies"
const url_wp_api_media = "/media"
const url_wp_api_user = "/users"
const url_wp_api_types = "/types"
const url_wp_api_statuses = "/statuses"
const url_wp_api_settings = "/settings"

//WordPress v2 restful api

func wpApiInitRoute(app *iris.Application){
	//wpApi := app.Party(url_wp_base_url)
	//postsCtrl := wp.PostsCtrl{}
	//wpApi.Any(url_wp_api_posts,postsCtrl)
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
	app.Run(application, cfgPath)
}
