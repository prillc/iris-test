package routes

import (
	"github.com/kataras/iris/v12"
	"iris-test/web/controllers"
)

func ValidationRoute(app *iris.Application) {
	app.Get("/validate", controllers.ValidationTest)
}
