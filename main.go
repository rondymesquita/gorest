package main

import (
	"github.com/rondymesquita/gorest/app"
)

func main() {
	var app app.App
	app.Create()
	app.Start()

}
