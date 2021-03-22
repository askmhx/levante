package repositories

import (
	"github.com/jinzhu/gorm"
	"levante/model"
	"levante/orm"
)

type PostRepository interface {
	SelectMany(start,limit uint64) []orm.Post
	SelectOne(pid uint64) (post orm.Post, found bool)
	QueryList(sql string) []orm.Post
	QueryColCount(sql string) []model.ColCount
}


func NewPostRepository(database *gorm.DB) PostRepository {
	return &postRepository{database: database}
}

type postRepository struct {
	database *gorm.DB
}

func (this *postRepository) QueryColCount(sql string) []model.ColCount{
	var result = []model.ColCount{}
	this.database.Raw(sql).Scan(&result)
	return result
}

func (this *postRepository) QueryList(sql string) []orm.Post{
	var result = []orm.Post{}
	this.database.Raw(sql).Scan(&result)
	return result
}

func (this *postRepository) SelectMany(start, limit uint64) []orm.Post {
	var postList []orm.Post
	 this.database.Order("created_at DESC").Offset(start).Limit(limit).Find(&postList)
	 return postList
}

func (this *postRepository) SelectOne(pid uint64) (result orm.Post, found bool) {
	var post orm.Post
	if err := this.database.Where("id = ? ", pid).Find(&post).Error; err == nil {
		return post,true
	}
	return orm.Post{},false
}


