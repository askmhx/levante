package admin

import (
	"github.com/kataras/iris/context"
	"github.com/jinzhu/gorm"
)

type DashCtrl struct {
	DB *gorm.DB
}

func (this *DashCtrl) IndexHandle(ctx context.Context) {

}
