package ctrl

import (
	"github.com/kataras/iris/context"
	"github.com/jinzhu/gorm"
)

type PhotoCtrl struct {
	DB *gorm.DB
}

func (this *PhotoCtrl) ReadHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}

func (this *PhotoCtrl) IndexHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}

func (this *PhotoCtrl) UpdateHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}
