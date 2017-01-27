package test

import (
	"fmt"
	"gorest/server"
	//"os"
	"testing"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"github.com/stretchr/testify/assert"
	"bytes"
	"encoding/json"
	"gorest/model"
)

var app server.App
var host string = "http://localhost"
func setup() {
	gin.SetMode(gin.TestMode)
	app.Create()
	go app.Start()

}

func teardown() {
	app.Stop()
}

//func TestMain(m *testing.M) {
//	setup()
//	retCode := m.Run()
//	teardown()
//	os.Exit(retCode)
//}

//func TestShouldDoAGetOnServer(t *testing.T) {
//
//	response, _ := http.Get("http://localhost:3000/ping")
//	assert.Equal(t, 200, response.StatusCode)
//}
//
//func TestShouldDoAGetOnServerAgain(t *testing.T) {
//
//	response, _ := http.Get("http://localhost:3000/create")
//	assert.Equal(t, 200, response.StatusCode)
//}

func TestAnotherOne(t *testing.T) {
	setup()
	var mockHelper MockHelper
	mock := mockHelper.buildBasicJsonGet()
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(mock)

	response, _ := http.Post("http://localhost:3000/create", "application/json; charset=utf-8", buffer)
	assert.Equal(t, 200, response.StatusCode)
	var output model.Mock
	json.NewDecoder(response.Body).Decode(&output)
	fmt.Println(output)
	teardown()

}

func TestAllFlow(t *testing.T) {
	setup()
	response, _ := http.Get("http://localhost:3000/json-get")
	assert.Equal(t, 404, response.StatusCode)

	var mockHelper MockHelper
	mock := mockHelper.buildBasicJsonGet()
	jsonString := new(bytes.Buffer)
	json.NewEncoder(jsonString).Encode(mock)

	response, _ = http.Post("http://localhost:3000/create", "application/json; charset=utf-8", jsonString)
	assert.Equal(t, 200, response.StatusCode)
	var output model.Mock
	json.NewDecoder(response.Body).Decode(&output)
	//fmt.Println(output)

	response, _ = http.Get("http://localhost:3000/json-get")
	assert.Equal(t, 200, response.StatusCode)
	teardown()
}
