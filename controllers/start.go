package controllers

import (
	"github.com/kataras/iris/mvc"
	"iosxc.com/levante/services"
)

type StartController struct {
	BaseController
	LinkService services.LinkService
}

func (this *StartController) Get() mvc.Result {
	linkGroupList := this.LinkService.GetLinkGroups()
	return ViewPageWithModel("start", "linkGroupList", linkGroupList)
}