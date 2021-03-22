package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"levante/model"
)

type BaseController struct {
	Ctx  iris.Context
}

var ViewPageNotFound = mvc.View{
	Name: "front/error.html",
	Data: model.RspDataPageNotFound,
}

var ViewPagePlant = func(html string) mvc.View {
	return mvc.View{
		Name: "front/"+html+".html",
	}
}

var ViewPageWithDataMap = func(html string,dataMap iris.Map) mvc.View {
	return mvc.View{
		Name: "front/"+html+".html",
		Data: dataMap,
	}
}

var ViewPageWithModel = func(html,key string,model interface{}) mvc.View {
	return mvc.View{
		Name: "front/"+html+".html",
		Data: iris.Map{
			key: model,
		},
	}
}