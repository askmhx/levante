package ctrl

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/context"
	"iosxc.com/levante/orm"
	"iosxc.com/levante/util"
)

type IndexCtrl struct {
	DB *gorm.DB
}


func (this *IndexCtrl) IndexHandle(context context.Context) {
	var postList []orm.Post
	this.DB.Order("created_at DESC").Limit(5).Find(&postList)
	context.ViewData("postList", postList)
	//TODO
	//context.ViewData("monthList", postList)
	//context.ViewData("lastList", postList)
	//context.ViewData("tagList", nil)
	context.View("front/index.html")
}

func (this *IndexCtrl) ArchiveHandle(context context.Context) {
	page := util.GetPage(context)
	var postList []orm.Post
	this.DB.Offset(page.Start()).Limit(page.End()).Find(&postList)
	context.ViewData("postList", postList)
	context.View("front/archive.html")
}

func (this *IndexCtrl) StartHandle(context context.Context) {
	linkGroupsDO := []orm.LinkGroup{}
	this.DB.Order("sort DESC").Find(&linkGroupsDO)
	linkGroupMap := map[string][]orm.Link{}
	for _, linkGroup := range linkGroupsDO {
		links := []orm.Link{}
		this.DB.Model(&linkGroup).Related(&links).Order("sort ASC")
		if(len(links)>0){
			linkGroupMap[linkGroup.Title]= links
		}
	}
	context.ViewData("linkGroupMap", linkGroupMap)
	context.View("front/start.html")
}

func (this *IndexCtrl) AboutHandle(context context.Context) {
	context.View("front/about.html")
}
