package routes

import (
	"github.com/kataras/iris/v12"
	"iris-test/web/controllers"
)

func CookieRoute(app *iris.Application) {
	party := app.Party("/cookie")
	{
		party.Get("/{key}/{value}", controllers.SetCookieTest)
		party.Delete("/{key}", controllers.DeleteCookie)
	}
}
