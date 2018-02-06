package ctrl

import (
	"iosxc.com/levante/orm"
	"iosxc.com/levante/util"
	"iosxc.com/levante/app"
)

type IndexCtrl struct {
}

func (this *IndexCtrl) IndexHandle(ctx *app.Context) {
	var postList []orm.Post
	ctx.Database().Order("created_at DESC").Limit(5).Find(&postList)
	ctx.ViewData("postList", postList)
	type ColCount struct {
		ColName string
		Count   int
	}

	var mcList []ColCount
	ctx.Database().Raw("select date_format(updated_at, '%Y-%m') as col_name,count(1) as count from levante.posts group by col_name order by col_name desc").Scan(&mcList);
	ctx.ViewData("monthList", mcList)
	mcList = nil
	ctx.Database().Raw("select catalog as col_name,count(1) as count from levante.posts group by col_name order by col_name desc").Scan(&mcList);
	ctx.ViewData("tagList", mcList)
	ctx.View("front/index.html")
}

func (this *IndexCtrl) ArchiveHandle(ctx *app.Context) {
	page := util.GetPage(ctx)
	var postList []orm.Post
	ctx.Database().Offset(page.Start()).Limit(page.End()).Find(&postList)
	ctx.ViewData("postList", postList)
	ctx.View("front/archive.html")
}

func (this *IndexCtrl) StartHandle(ctx *app.Context) {

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
	var sql = "select l.id,l.url,l.title as title,l.image,l.description,l.owner,l.highlight ,g.title as gtitle from links l left join link_groups g on l.link_group_id = g.id order by g.sort,l.sort,,l.id";
	ctx.Database().Raw(sql).Scan(&linkList);

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

	ctx.ViewData("linkGroupList", linkGroupList)
	ctx.View("front/start.html")
}

func (this *IndexCtrl) AboutHandle(ctx *app.Context) {
	ctx.View("front/about.html")
}
