package main

import (
	"fmt"
	"iris-test/infrastructure"
	"iris-test/web/initial"
	"iris-test/web/middleware"
	"iris-test/web/routes"
	"iris-test/web/validate"
	"log"
	"strconv"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	// 注册中间件
	middleware.Register(app)

	// 注册html
	app.RegisterView(iris.HTML("./web/views/", ".html").Reload(true))
	// 最基本的Get请求
	app.Get("/ping", func(ctx iris.Context) {
		_, _ = ctx.WriteString("ping")
	})
	// 返回.html
	app.Get("/index", func(ctx iris.Context) {
		_ = ctx.View("index.html")
	})
	// 路由的分组
	userRoute := app.Party("/users")
	{
		// 带参数的路由
		userRoute.Get("/{id:int min(1)}", func(ctx iris.Context) {
			userId, err := ctx.Params().GetInt("id")
			fmt.Println(userId, err)
			_, _ = ctx.WriteString(strconv.Itoa(userId))
		})
	}

	// 路由初始化
	routes.RouteInit(app)
	// 初始化校验
	validate.ValidatorInit()
	// 迁移数据库
	initial.ModelMigrate(infrastructure.GlobalRepositories.DB)

	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		log.Fatal("启动失败: ", err)
	}
}
