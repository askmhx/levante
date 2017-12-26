package app

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/jinzhu/gorm"
	"sync"
	"github.com/kataras/iris/context"
)

type Context struct {
	iris.Context
	session *sessions.Session
	database *gorm.DB
	config *AppConfig
}

func (ctx *Context) Session() *sessions.Session {
	if ctx.session == nil {
		ctx.session = ctxHolder.SessionsManager.Start(ctx.Context)
	}
	return ctx.session
}

func (ctx *Context) Database() *gorm.DB {
	if ctx.database == nil {
		ctx.database = ctxHolder.Database
	}
	return ctx.database
}

func (ctx *Context) Config() *AppConfig {
	if ctx.config == nil {
		ctx.config = ctxHolder.Config
	}
	return ctx.config
}


var contextPool = sync.Pool{New: func() interface{} {
	return &Context{}
}}

func acquire(original iris.Context) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Context = original
	ctx.session = nil
	return ctx
}

func release(ctx *Context) {
	contextPool.Put(ctx)
}

func Handler(h func(*Context)) context.Handler {
	return func(original iris.Context) {
		ctx := acquire(original)
		h(ctx)
		release(ctx)
	}
}
