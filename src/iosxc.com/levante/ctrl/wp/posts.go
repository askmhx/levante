package wp

import (
	"github.com/jinzhu/gorm"
)

type PostsCtrl struct {
	DB *gorm.DB
}