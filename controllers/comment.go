package controllers

import "github.com/kataras/iris/v12/mvc"

type CommentController struct {
	BaseController
}

func (this *CommentController) Create() mvc.Result{

	return ViewPageNotFound
}