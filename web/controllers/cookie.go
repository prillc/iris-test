package controllers

import (
	"github.com/kataras/iris/v12"
)

func SetCookieTest(ctx iris.Context) {
	key := ctx.Params().GetString("key")
	value := ctx.Params().GetString("value")
	ctx.SetCookieKV(key, value)

	_, _ = ctx.Writef("cookie设置成功 %s=%s", key, value)
}

func DeleteCookie(ctx iris.Context) {
	key := ctx.Params().Get("key")
	ctx.RemoveCookie(key)
	_, _ = ctx.Writef("删除成功 %s", key)
}
