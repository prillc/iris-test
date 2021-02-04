package routes

import (
	"github.com/iris-contrib/middleware/jwt"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"iris-test/web/controllers"
	"iris-test/web/middleware"
)

func RequestRoute(app *iris.Application) {
	app.Post("/req", controllers.TestPostJsonHandler)

	jwtRoute := app.Party("/jwt")
	{
		jwtRoute.Use(middleware.JWTMiddleware.Serve)
		jwtRoute.Get("", controllers.TestJWTHandler)
	}
}
