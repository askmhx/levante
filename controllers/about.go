package controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"iosxc.com/levante/services"
)

type AboutController struct {
	PostService services.PostService
	LinkService services.LinkService
}

func (this *AboutController) Get() mvc.Result {
	return ViewPagePlant("about")
}