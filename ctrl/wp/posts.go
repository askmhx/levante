package wp

import (
	"iosxc.com/levante/orm"
	"iosxc.com/levante/app"
	"github.com/kataras/iris"
)

type WpPostsCtrl struct {
	iris.Controller
}

func (this *WpPostsCtrl) Get(ctx *app.Context) {
	pid := ctx.Params().Get("pid")
	post := orm.Post{}
	if err := ctx.Database().Where("id = ? ", pid).Find(&post).Error; err != nil {
		ctx.View("404.html")
		return
	}
	ctx.ViewData("post", post)
	ctx.View("front/post.html")
}

func (this *WpPostsCtrl) Post(ctx *app.Context) {
	pid := ctx.Params().Get("pid")
	post := orm.Post{}
	if err := ctx.Database().Where("id = ? ", pid).Find(&post).Error; err != nil {
		ctx.View("404.html")
		return
	}
	ctx.ViewData("post", post)
	ctx.View("front/post.html")
}

func (this *WpPostsCtrl) Delete(ctx *app.Context) {
	pid := ctx.Params().Get("pid")
	post := orm.Post{}
	if err := ctx.Database().Where("id = ? ", pid).Find(&post).Error; err != nil {
		ctx.View("404.html")
		return
	}
	ctx.ViewData("post", post)
	ctx.View("front/post.html")
}