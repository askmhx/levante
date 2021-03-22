package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"levante/services"
	"levante/util"
)

type PostController struct {
	BaseController
	PostService services.PostService
}


func (this *PostController) Get() mvc.Result {
	page := util.GetPage(this.Ctx)
	post := this.PostService.GetList(page.Start(),page.End())
	return ViewPageWithModel("archive", "postList", post)
}

func (this *PostController) GetCatalogBy(catalog string) mvc.Result {
	page := util.GetPage(this.Ctx)
	post := this.PostService.GetList(page.Start(),page.End())
	return ViewPageWithModel("archive", "postList", post)
}

func (this *PostController) GetMonthBy(month string) mvc.Result {
	page := util.GetPage(this.Ctx)
	post := this.PostService.GetList(page.Start(),page.End())
	return ViewPageWithModel("archive", "postList", post)
}


func (this *PostController) GetBy(pid uint64) mvc.Result {
	post, flag := this.PostService.GetByID(pid)
	if !flag {
		return ViewPageNotFound
	}
	return ViewPageWithModel("post", "post", post)
}