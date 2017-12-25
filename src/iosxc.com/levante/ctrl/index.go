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
		ColName string
		Count   int
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

	type LinkData struct {
		Id          int
		Title       string
		Url         string
		Image       string
		Description string
		Owner       string
		Highlight   int
		Gtitle      string
	}
	var linkList []LinkData
	var sql = "select l.id,l.url,l.title as title,l.image,l.description,l.owner,l.highlight ,g.title as gtitle from links l left join link_groups g on l.link_group_id = g.id order by g.sort,l.sort";
	this.DB.Raw(sql).Scan(&linkList);

	type LinkGroupEntry struct {
		Title    string
		LinkList []LinkData
	}

	var linkGroupList []LinkGroupEntry

	var groupEntry = LinkGroupEntry{Title: "", LinkList: []LinkData{}}

	for _, link := range linkList {
		if groupEntry.Title == "" || groupEntry.Title != link.Gtitle {
			if groupEntry.Title != "" {
				linkGroupList = append(linkGroupList, groupEntry)
			}
			groupEntry = LinkGroupEntry{Title: link.Gtitle, LinkList: []LinkData{}}
		}
		groupEntry.LinkList = append(groupEntry.LinkList, link)
	}
	linkGroupList = append(linkGroupList, groupEntry)

	context.ViewData("linkGroupList", linkGroupList)
	context.View("front/start.html")
}

func (this *IndexCtrl) AboutHandle(context context.Context) {
	context.View("front/about.html")
}
