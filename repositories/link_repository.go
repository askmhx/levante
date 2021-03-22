package repositories

import (
	"github.com/jinzhu/gorm"
	"levante/model"
)

type LinkRepository interface {
	QueryLinkDatas(sql string) []model.LinkData
}


func NewLinkRepository(database *gorm.DB) LinkRepository {
	return &linkRepository{database: database}
}

type linkRepository struct {
	database *gorm.DB
}


func (this *linkRepository) QueryLinkDatas(sql string) []model.LinkData{
	var result []model.LinkData
	this.database.Raw(sql).Scan(&result)
	return result
}


