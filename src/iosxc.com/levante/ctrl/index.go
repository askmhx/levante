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

const nav_index_key = "navIndexKey"
const nav_page_index = "index"
const nav_page_archive = "archive"
const nav_page_start = "start"
const nav_page_about = "about"

func (this *IndexCtrl) IndexHandle(context context.Context) {
	context.ViewData(nav_index_key, nav_page_index)
	context.View("index.html")
}

func (this *IndexCtrl) ArchiveHandle(context context.Context) {
	context.ViewData(nav_index_key, nav_page_archive)
	page := util.GetPage(context)
	var posts []orm.Post
	this.DB.Offset(page.Start()).Limit(page.End()).Find(&posts)
	context.ViewData("posts", posts)
	context.View("archive.html")
}

func (this *IndexCtrl) StartHandle(context context.Context) {
	context.ViewData(nav_index_key, nav_page_start)
	var links []orm.Link
	this.DB.Select(&links)
	context.ViewData("links", links)
	context.View("start.html")
}

func (this *IndexCtrl) AboutHandle(context context.Context) {
	context.ViewData(nav_index_key, nav_page_about)
	context.View("about.html")
}
