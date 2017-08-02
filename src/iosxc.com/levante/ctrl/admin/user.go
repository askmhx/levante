package admin

import (
	"github.com/kataras/iris/context"
	"github.com/jinzhu/gorm"
)

type UserCtrl struct {
	DB *gorm.DB
}

func (this *UserCtrl) LoginHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}

func (this *UserCtrl) ResetPwdHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}

func (this *UserCtrl) UpdateHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}

func (this *UserCtrl) DeleteHandle(ctx context.Context) {
	ctx.Writef("hello world!")

}
