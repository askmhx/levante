package util

import (
	"github.com/kataras/iris/v12"
	"levante/model"
	"strconv"
)

func GetPage(context iris.Context) model.Page {
	page := model.Page{}
	var err error
	page.PageIndex, err = strconv.ParseUint(context.Params().Get(CONST_PARAM_PAGE_INDEX), 10, 64)
	page.PageSize, err = strconv.ParseUint(context.Params().Get(CONST_PARAM_PAGE_SIZE), 10, 64)
	if err != nil {
		page.PageSize = CONST_DEFAULT_PAGE_SIZE
		page.PageIndex = 0
	}
	return page
}
