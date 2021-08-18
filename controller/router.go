package controller

import (
	"bbs/controller/api"
	"bbs/pkg/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
)

func Router() {
	app := iris.New()
	// 配置中间件 Logger，recover，cors，AllowMethods
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	/*app.OnAnyErrorCode(func(ctx iris.Context) {
		println("OnAnyErrorCode", ctx.GetErr().Error())
	})*/
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/test").Handle(new(api.TestController))
		m.Party("/article").Handle(new(api.ArticleController))
		m.Party("/topic").Handle(new(api.TopicController))
	})
	server := &http.Server{Addr: ":" + config.Instance.Port}
	app.Run(iris.Server(server))
}
