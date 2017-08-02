package util

import (
	"github.com/kataras/iris/context"
	"iosxc.com/levante/model"
	"strconv"
)

func GetPage(context context.Context) model.Page {
	page := model.Page{}
	page.PageIndex, _ = strconv.ParseUint(context.Params().Get(CONST_PARAM_PAGE_INDEX), 10, 64)
	page.PageSize, _ = strconv.ParseUint(context.Params().Get(CONST_PARAM_PAGE_SIZE), 10, 64)
	return page
}
