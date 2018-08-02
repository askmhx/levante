package ctrl

import (
	"iosxc.com/levante/orm"
	"iosxc.com/levante/app"
	)

type PostCtrl struct {
}

func (this *PostCtrl) ReadHandle(ctx *app.Context) {
	pid := ctx.Params().Get("pid")
	post := orm.Post{}
	if err := ctx.Database().Where("id = ? ", pid).Find(&post).Error; err != nil {
		ctx.View("404.html")
		return
	}
	ctx.ViewData("post", post)
	ctx.View("front/post.html")
}