package orm

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	OprBaseModel
	Title      string `gorm:"primary_key"`
	Content    string `gorm:"primary_key"`
	Author     string
	Support    uint64
	Oppose     uint64
	SourceLink string
}
