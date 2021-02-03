package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"iris-test/common"
	"log"
)

func jwtErrorHandler(ctx iris.Context, err error) {
	log.Fatal("jwtErrorHandler", err)
	_, _ = ctx.JSON(common.Response{
		Status: 401,
		Msg:    "未登录",
		Data:   nil,
	})
}

var JWTMiddleware = jwt.New(jwt.Config{
	ValidationKeyGetter: nil,
	ContextKey:          "",
	ErrorHandler:        jwtErrorHandler,
	CredentialsOptional: false,
	Extractor:           nil,
	EnableAuthOnOptions: false,
	SigningMethod:       jwt.SigningMethodHS256,
	Expiration:          false,
})