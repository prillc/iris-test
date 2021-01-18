package main

import (
	"github.com/kataras/iris/v12/mvc"
	"iris-test/repositories"
	"iris-test/services"
	"iris-test/web/controllers"
	"log"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	mvc.Configure(app.Party("/movies"), movies)

	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		log.Fatal("启动失败: ", err)
	}
}

func movies(app *mvc.Application) {
	resp := repositories.NewMovieRepository()
	movieService := services.NewMovieService(resp)
	app.Register(movieService)
	app.Handle(new(controllers.MovieController))
}
