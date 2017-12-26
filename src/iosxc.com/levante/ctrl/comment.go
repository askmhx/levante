package ctrl

import (
	"iosxc.com/levante/app"
)

type CommentCtrl struct {
}

func (this *CommentCtrl) ReadHandle(ctx *app.Context) {
	ctx.Writef("hello world!")

}

func (this *CommentCtrl) CreateHandle(ctx *app.Context) {
	ctx.Writef("hello world!")

}