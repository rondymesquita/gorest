package app

import(
	"github.com/rondymesquita/gorest/model"
	"gopkg.in/gin-gonic/gin.v1"
	"fmt"
)

type RouteBuilder struct{
	app *App
}

func (routeBuilder *RouteBuilder) BuildFrom(mock model.Mock){
	routeBuilder.app.Engine.GET(mock.Path, func(context *gin.Context) {
		//context.Header("name", "Fulano")
		headers := map[string][]string{
			"Accept-Encoding": {"gzip, deflate"},
			"Accept-Language": {"en-us"},
			"Foo": {"Bar", "two"},
		}
		for key, values := range headers{
			fmt.Println(key)
			fmt.Println(values)
			for subValue := range values{
				context.Header(key, values[subValue])
			}
		}
		context.JSON(200, mock.Response.Body)
	})
}