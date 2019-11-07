package controllers

import (
	"github.com/kataras/iris/v12/mvc"
)

type PhotoController struct {
	BaseController
}

func (this *PhotoController) ReadHandle() mvc.Result{
	return ViewPageNotFound

}

func (this *PhotoController) IndexHandle() mvc.Result{
	return ViewPageNotFound

}

func (this *PhotoController) UpdateHandle() mvc.Result{
	return ViewPageNotFound

}
