package main

import (
	"fmt"
	"iris-test/web/routes"
	"log"
	"strconv"

	"github.com/kataras/iris/v12"
)

func logThisMiddleware(ctx iris.Context) {
	//ctx.Path() 请求的url
	ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())

	//ctx.Next 继续往下一个处理方法 中间件 如果没有他 那么就不会执行 usersRoutes
	ctx.Next()
}

func main() {
	app := iris.New()

	// 注册html
	app.RegisterView(iris.HTML("./web/views/", ".html").Reload(true))

	// 中间件，打印

	app.Use(logThisMiddleware)
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

	// cookie 测试
	routes.CookieRoute(app)

	err := app.Run(iris.Addr(":8080"))
	if err != nil {
		log.Fatal("启动失败: ", err)
	}
}
