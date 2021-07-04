package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	//iris实例化
	app := iris.New()
	//设置错误模式
	app.Logger().Debug("debug")
	//注册模板
	template := iris.HTML("./backend/web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)
	//设置静态目录
	app.HandleDir("/assets", "./backend/web/assets")
	//出现错误跳转
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问页面出错！"))
		ctx.ViewLayout("")
		ctx.View("shard/error.html")
	})
	//注册控制器
	//启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
