package router

import (
	"github.com/kataras/iris"
	"fmt"
	"webhook-router/configuration"
	)

func InitRouter(configuration *configuration.GlobalConfiguration) {
	app := iris.New()

	app.Any("*", handler)

	app.Run(iris.Addr(fmt.Sprintf(":%v", configuration.Port)))
}

func handler(ctx iris.Context) {
	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
}
