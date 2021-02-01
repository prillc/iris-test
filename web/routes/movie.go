package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-test/infrastructure"
	"iris-test/services"
	"iris-test/web/controllers"
)

func MovieRoute(app *iris.Application) {
	// iris mvc 注册
	//mvc.New(app.Party("/movie")).Handle(controllers.MovieController{})
	movieService := services.NewMovieService(infrastructure.GlobalRepositories.MovieRepository)
	movieHistoryService := services.NewMovieHistoryService(infrastructure.GlobalRepositories.MovieHistoryRepository)

	mvc.New(app.Party("/movies")).Handle(&controllers.MovieController{Service: movieService})
	mvc.New(app.Party("/movies/histories")).Handle(&controllers.MovieHistoryController{Service: movieHistoryService})
}
