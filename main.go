package main

import (
	"github.com/rondymesquita/gorest/server"
)

func main() {
	var app app.App
	app.Create()
	app.Start()

}
