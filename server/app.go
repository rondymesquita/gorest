package server

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/braintree/manners"
	"gorest/model"
)

type App struct {
	Engine *gin.Engine
	Port   string
}

func (app *App) Create() {
	// gin.SetMode(gin.TestMode)
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
		responseData := model.ResponseData{Message: "Route created with success.", Status: "SUCCESS"}
		context.JSON(200, responseData)
	})

}

func (app *App) Start() {
	fmt.Println("===> Starting Server")
	manners.ListenAndServe(fmt.Sprintf(":%s", app.Port), app.Engine)
}

func (app *App) Stop() {
	manners.Close()
	fmt.Println("===> Server Stopped ")
}
