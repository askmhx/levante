package ctrl

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/context"
	"iosxc.com/levante/model"
	"iosxc.com/levante/orm"
	"iosxc.com/levante/util"
)

type PostCtrl struct {
	DB *gorm.DB
}

func (this *PostCtrl) ReadHandle(context context.Context) {
	pid := context.Params().Get("pid")
	post := orm.Post{}
	if err := this.DB.Where("id = ? ", pid).Find(&post).Error; err != nil {
		context.View("404.html")
		return
	}
	context.ViewData("post", post)
	context.View("front/post.html")
}

func (this *PostCtrl) CreateHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}

func (this *PostCtrl) UpdateHandle(ctx context.Context) {
	ctx.Writef("hello world!")
}

func (this *PostCtrl) DeleteHandle(ctx context.Context) {
	pid := ctx.URLParam("pid")
	post := orm.Post{}
	db := ctx.Value(util.CONST_APP_DB).(*gorm.DB)
	if err := db.Where("id = ?", pid).First(&post).Error; err != nil {
		ctx.View("front/404.html")
		return
	}
	post.IsDeleted = true
	db.Update(post)
	ret := model.OperationResult{}
	ret.Desc = "删除成功"
	ret.Code = model.OperationResultCodeSuccess
	ret.RetURL = "https://www.baidu.com"
	ctx.ViewData("result", ret)
	ctx.View("front/result.html")

}
