package server

import (
	"fmt"
	"gorest/model"

	"github.com/braintree/manners"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"gorest/util"
)

type App struct {
	Engine *gin.Engine
	Port   string
}

func (app *App) Create() {
	//gin.SetMode(gin.ReleaseMode)
	app.Engine = gin.Default()
	app.Port = "3000"

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
		responseData := model.ResponseData{Message: constant.RouteCreatedWithSuccess, Status: "SUCCESS"}
		context.JSON(200, responseData)
	})

}

func (app *App) Start() {
	log.Println("Starting Server")
	manners.ListenAndServe(fmt.Sprintf(":%s", app.Port), app.Engine)
}

func (app *App) Stop() {
	manners.Close()
	log.Println("Server Stopped ")
}
