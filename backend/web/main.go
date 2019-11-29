package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	//创建iris实例
	app := iris.New()
	//设置错误模式
	app.Logger().SetLevel("debug")
	//注册模板
	tpl := iris.HTML("./backend/web/views", "html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tpl)
	//设置模板目录
	app.HandleDir("/assets","./backend/web/assets")
	//设置公共错误,当状态码>=400的时候
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.View("message", ctx.Values().GetStringDefault("message", "访问的页面出错!"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	//注册控制器

	//启动服务
	app.Run(
		iris.Addr(":8090"),
		//配置启动参数
		iris.WithConfiguration(iris.Configuration{
			IgnoreServerErrors:                nil,
			TimeFormat:                        "2006-01-02 15:04:05",
			Charset:                           "utf-8",
		}),
	)
}
