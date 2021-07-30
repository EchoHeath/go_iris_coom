package main

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/opentracing/opentracing-go/log"
	"go-electricity/backend/web/controllers"
	"go-electricity/common"
	"go-electricity/repositories"
	"go-electricity/services"
)

func main() {
	//iris实例化
	app := iris.New()
	//设置错误模式
	app.Logger().SetLevel("debug")
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

	//连接数据库
	db, err := common.NewMysqlConn()
	if err != nil {
		log.Error(nil)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//注册控制器
	productRepository := repositories.NewProductManager("product", db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))

	//启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
