package server

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/braintree/manners"
	"gorest/model"
)

type App struct {
	Engine *gin.Engine
	Port string
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
		app.buildFrom(mock)

		context.JSON(200, mock)
	})

}

func (app *App) buildFrom(mock model.Mock) {
	fmt.Println(mock.Response.Type)
	app.Engine.GET(mock.Path, func(context *gin.Context) {
		context.JSON(200, mock.Response.Data)
	})
}

func (app *App) Start() {
	//app.Engine.Run(":3000")
	fmt.Println("===> Server Started")
	manners.ListenAndServe(fmt.Sprintf(":%s", app.Port), app.Engine)
}

func (app *App) Stop() {
	manners.Close()
	fmt.Println("===> Server Stopped ")
}
