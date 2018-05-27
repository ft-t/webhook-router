package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"time"

	"webhook-router/configuration"
	"webhook-router/db"

	"github.com/kataras/iris"
	"github.com/kataras/iris/core/netutil"
)

func InitRouter(configuration *configuration.GlobalConfiguration) {
	app := iris.New()

	app.Any("*", handler)

	app.Run(iris.Addr(fmt.Sprintf(":%v", configuration.Port)))
}

func handler(ctx iris.Context) {
	rules := db.GetRulesByPath(ctx.Path())

	success := false
	client := netutil.Client(time.Duration(20 * time.Second))

	for _, rule := range rules {
		sort.Slice(rule.Endpoints, func(i, j int) bool {
			return rule.Endpoints[i].Priority < rule.Endpoints[j].Priority
		})

		for _, endpoint := range rule.Endpoints {
			request := http.Request{}

			request.URL, _ = url.Parse(endpoint.Url)
			request.Body = ctx.Request().Body
			request.ContentLength = ctx.Request().ContentLength
			request.Header = ctx.Request().Header
			request.Method = ctx.Request().Method
			request.Form = ctx.Request().Form

			response, err := client.Do(&request)

			if err == nil {
				body, err2 := ioutil.ReadAll(response.Body)

				for key, values := range response.Header {
					for _, val := range values {
						ctx.Header(key, val)
					}
				}

				if err2 == nil {
					ctx.StatusCode(response.StatusCode)
					ctx.Write(body)

					success = true
					break
				}
			}
		}

		if success {
			break
		}
	}

	if !success {
		ctx.StatusCode(iris.StatusNotFound)
	}
}
