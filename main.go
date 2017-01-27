package main

import (
	"gorest/server"
)

func main() {
	var app server.App
	app.Create()
	app.Start()

}
