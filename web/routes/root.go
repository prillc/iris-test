package routes

import "github.com/kataras/iris/v12"

func RouteInit(app *iris.Application) {
	// 路由的注册
	CookieRoute(app)
	ValidationRoute(app)
	RequestRoute(app)
	MovieRoute(app)
}
