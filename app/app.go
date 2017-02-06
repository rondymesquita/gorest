package app

import (
	"fmt"
	"github.com/rondymesquita/gorest/model"

	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"github.com/rondymesquita/gorest/constant"
)

const(
	Port = "9091"
)

type App struct {
	Engine *gin.Engine
}

func (app *App) Create() *gin.Engine{
	//gin.SetMode(gin.ReleaseMode)
	app.Engine = gin.Default()

	app.Engine.GET("/ping", func(context *gin.Context) {
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
			fmt.Println("========================")
		}
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	app.Engine.POST("/create", func(context *gin.Context) {
		var mock model.Mock
		context.BindJSON(&mock)
		fmt.Println(mock.String())

		routeBuilder := RouteBuilder{app}
		routeBuilder.BuildFrom(mock)

		//TODO validate if route already exists and return a response message for it
		responseData := model.ResponseMessage{Message: constant.RouteCreatedWithSuccess, Status: constant.Success}
		context.JSON(201, responseData)
	})
	return app.Engine

}

func (app *App) Start() {
	log.Println("Starting Server")
	//manners.ListenAndServe(fmt.Sprintf(":%s", Port), app.Engine)
	app.Engine.Run(fmt.Sprintf(":%s", Port))
}
//
//func (app *App) Stop() {
//	manners.Close()
//	log.Println("Server Stopped ")
//}
