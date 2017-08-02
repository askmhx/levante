package ctrl

import (
	"github.com/kataras/iris/context"
	"github.com/jinzhu/gorm"
)

type CommentCtrl struct {
	db *gorm.DB
}

func (this *CommentCtrl) ReadHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}

func (this *CommentCtrl) CreateHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}

func (this *CommentCtrl) UpdateHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}

func (this *CommentCtrl) DeleteHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}
