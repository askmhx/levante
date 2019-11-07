package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iosxc.com/levante/services"
)

type IndexController struct {
	PostService services.PostService
	LinkService services.LinkService
}

func (this *IndexController) Get() mvc.Result{
	var postList = this.PostService.GetList(0,5)
	var mcList = this.PostService.GetMonthList()
	var tagList = this.PostService.GetTagList()
	var dataMap = iris.Map{}
	dataMap["postList"] = postList
	dataMap["monthList"] =  mcList
	dataMap["tagList"] = tagList
	return ViewPageWithDataMap("index",dataMap)

}

