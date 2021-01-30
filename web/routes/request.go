package routes

import (
	"github.com/kataras/iris/v12"
	"iris-test/web/controllers"
)

func RequestRoute(app *iris.Application) {
	app.Post("/req", controllers.TestPostJsonHandler)
}
