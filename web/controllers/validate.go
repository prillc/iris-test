package controllers

import (
	"github.com/kataras/iris/v12"
	"iris-test/web/validate"
)

func ValidationTest(ctx iris.Context) {
	var foo validate.Foo
	_ = ctx.ReadQuery(&foo)
	if err := validate.Validate.Struct(foo); err != nil {
		_, _ = ctx.Writef("error: %s", err)
		return
	}
	_, _ = ctx.WriteString("ok")
}
