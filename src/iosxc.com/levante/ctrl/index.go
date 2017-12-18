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
	type ColCount struct {
		ColName   string
		Count int
	}

	var mcList []ColCount
	this.DB.Raw("select date_format(updated_at, '%Y-%m') as col_name,count(1) as count from levante.posts group by col_name order by col_name desc").Scan(&mcList);
	context.ViewData("monthList", mcList)
	mcList = nil
	this.DB.Raw("select catalog as col_name,count(1) as count from levante.posts group by col_name order by col_name desc").Scan(&mcList);
	context.ViewData("tagList", mcList)
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
		if (len(links) > 0) {
			linkGroupMap[linkGroup.Title] = links
		}
	}
	context.ViewData("linkGroupMap", linkGroupMap)
	context.View("front/start.html")
}

func (this *IndexCtrl) AboutHandle(context context.Context) {
	context.View("front/about.html")
}
