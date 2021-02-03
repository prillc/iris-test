package routes

import (
	"github.com/iris-contrib/middleware/jwt"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"iris-test/web/controllers"
)

func RequestRoute(app *iris.Application) {
	app.Post("/req", controllers.TestPostJsonHandler)

	jwtRoute := app.Party("/jwt")
	{
		jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("My Secret"), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})
		jwtRoute.Use(jwtHandler.Serve)
		jwtRoute.Get("", controllers.TestJWTHandler)
	}
}
