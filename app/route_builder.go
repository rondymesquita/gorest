package app

import(
	"github.com/rondymesquita/gorest/model"
	"gopkg.in/gin-gonic/gin.v1"
)

type RouteBuilder struct{
	app *App
}

func (routeBuilder *RouteBuilder) BuildFrom(mock model.Mock){
	routeBuilder.app.Engine.GET(mock.Path, func(context *gin.Context) {
		context.JSON(mock.Response.StatusCode, mock.Response.Body)
	})
}