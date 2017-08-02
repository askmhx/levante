package app

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"iosxc.com/levante/ctrl"
	"iosxc.com/levante/ctrl/admin"
	"github.com/jinzhu/gorm"
)

const url_index = "/"
const url_start = "/start"
const url_archive = "/archive"
const url_about = "/about"
const url_photo = "/photo"
const url_post = "/post/{pid:int}"
const url_admin = "/this-is-a-fantastic-blog-management-system"
const url_admin_login = url_admin + "/login"

func initRoute(application *iris.Application, config *AppConfig,db *gorm.DB) {

	registerErrorHandler(application)

	indexCtrl := ctrl.IndexCtrl{db}

	application.Get(url_index, indexCtrl.IndexHandle)
	application.Get(url_start, indexCtrl.StartHandle)
	application.Get(url_archive, indexCtrl.ArchiveHandle)
	application.Get(url_about, indexCtrl.AboutHandle)

	photoCtrl := ctrl.PhotoCtrl{db}

	application.Get(url_photo, photoCtrl.IndexHandle)

	postCtrl := ctrl.PostCtrl{db}

	application.Get(url_post, postCtrl.ReadHandle)

	dashCtrl := admin.DashCtrl{db}

	application.Get(url_admin, dashCtrl.IndexHandle)

	userCtrl := admin.UserCtrl{db}

	application.Post(url_admin_login, userCtrl.LoginHandle)

}

func registerErrorHandler(application *iris.Application) {
	application.OnErrorCode(iris.StatusNotFound, func(ctx context.Context) {
		ctx.View("404.html")
	})
	//
	//application.OnErrorCode(iris.StatusInternalServerError, func(ctx context.Context) {
	//	ctx.View("500.html")
	//})
}
