package orm

import "github.com/jinzhu/gorm"

type Link struct {
	gorm.Model
	OprBaseModel
	Url         string
	Title       string
	Image       string
	Description string
	Owner       string
	Rating      uint
	Group       string
	Visible     string
	Highlight   string
	Sort        int
}
