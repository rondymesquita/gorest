package app

import (
	"fmt"
	"github.com/rondymesquita/gorest/model"

	"github.com/braintree/manners"
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

func (app *App) Create() {
	//gin.SetMode(gin.ReleaseMode)
	app.Engine = gin.Default()

	app.Engine.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	app.Engine.POST("/create", func(context *gin.Context) {
		var mock model.Mock
		context.BindJSON(&mock)

		routeBuilder := RouteBuilder{app}
		routeBuilder.BuildFrom(mock)

		//TODO validate if route already exists and return a response message for it
		responseData := model.ResponseMessage{Message: constant.RouteCreatedWithSuccess, Status: constant.Success}
		context.JSON(201, responseData)
	})

}

func (app *App) Start() {
	log.Println("Starting Server")
	manners.ListenAndServe(fmt.Sprintf(":%s", Port), app.Engine)
}

func (app *App) Stop() {
	manners.Close()
	log.Println("Server Stopped ")
}
