package util

import (
	"levante/model"
	"strconv"

	"github.com/kataras/iris/v12"
)

func GetPage(context iris.Context) model.Page {

	var err error
	pIndex, err := strconv.ParseInt(context.Params().Get(CONST_PARAM_PAGE_INDEX), 10, 64)
	pSize, err := strconv.ParseInt(context.Params().Get(CONST_PARAM_PAGE_SIZE), 10, 64)
	page := model.Page{PageIndex: int(pIndex), PageSize: int(pSize)}
	if err != nil {
		page.PageSize = CONST_DEFAULT_PAGE_SIZE
		page.PageIndex = 0
	}
	return page
}
