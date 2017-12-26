package ctrl

import (
	"iosxc.com/levante/app"
)

type PhotoCtrl struct {
}

func (this *PhotoCtrl) ReadHandle(ctx *app.Context) {
	ctx.Writef("hello world!")

}

func (this *PhotoCtrl) IndexHandle(ctx *app.Context) {
	ctx.Writef("hello world!")

}

func (this *PhotoCtrl) UpdateHandle(ctx *app.Context) {
	ctx.Writef("hello world!")

}
