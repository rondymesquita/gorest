package server

import(
	"gorest/model"
	"gopkg.in/gin-gonic/gin.v1"
)

type RouteBuilder struct{
	app *App
}

func (routeBuilder *RouteBuilder) BuildFrom(mock model.Mock){
	routeBuilder.app.Engine.GET(mock.Path, func(context *gin.Context) {
		context.JSON(200, mock.Response.Data)
	})
}