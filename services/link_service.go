package services

import (
	"levante/model"
	"levante/repositories"
)

type LinkService interface {
	GetLinkGroups() []model.LinkGroupEntry
}

func NewLinkService(repo repositories.LinkRepository) LinkService {
	return &linkService{
		repo: repo,
	}
}

type linkService struct {
	repo repositories.LinkRepository
}


func (this *linkService) GetLinkGroups() []model.LinkGroupEntry {
	var sql = "select l.id,l.url,l.title as title,l.image,l.description,l.owner,l.highlight ,g.title as gtitle from links l left join link_groups g on l.link_group_id = g.id order by g.sort,l.sort,l.id";
	var linkList = this.repo.QueryLinkDatas(sql);
	var groupEntry = model.LinkGroupEntry{Title: "", LinkList: []model.LinkData{}}
	var linkGroupList []model.LinkGroupEntry
	for _, link := range linkList {
		if groupEntry.Title == "" || groupEntry.Title != link.Gtitle {
			if groupEntry.Title != "" {
				linkGroupList = append(linkGroupList, groupEntry)
			}
			groupEntry = model.LinkGroupEntry{Title: link.Gtitle, LinkList: []model.LinkData{}}
		}
		groupEntry.LinkList = append(groupEntry.LinkList, link)
	}
	linkGroupList = append(linkGroupList, groupEntry)

	return linkGroupList

}
