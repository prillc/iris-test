package controllers

import (
	"github.com/kataras/iris/v12"
	"iris-test/models"
)

func TestPostJsonHandler(ctx iris.Context) {
	var jsonParam models.JsonParam
	if err := ctx.ReadJSON(&jsonParam); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.Writef("参数错误")
		return
	}

	_, _ = ctx.Writef("ok %v", jsonParam)
}
