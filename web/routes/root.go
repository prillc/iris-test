package routes

import "github.com/kataras/iris/v12"

func RouteInit(app *iris.Application) {
	CookieRoute(app)
	ValidationRoute(app)
}
