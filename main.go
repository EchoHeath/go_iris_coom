package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-electricity/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))
	//注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	mvc.New(app.Party("/student")).Handle(new(controllers.StudentController))

	app.Run(iris.Addr("localhost:8080"))
}

