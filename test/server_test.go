package test

import (
	"fmt"
	"gorest/server"
	"testing"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"github.com/stretchr/testify/assert"
	"bytes"
	"encoding/json"
	"gorest/model"
	"gorest/test/helper"
)

var app server.App
var host string

//func TestMain(m *testing.M) {
//	setup()
//	retCode := m.Run()
//	teardown()
//	os.Exit(retCode)
//}

func setup() {
	gin.SetMode(gin.TestMode)
	app.Create()
	host = fmt.Sprintf("http://localhost:%s", app.Port)
	go app.Start()

}

func teardown() {
	app.Stop()
}

func TestAnotherOne(t *testing.T) {
	setup()
	var mockHelper helper.MockHelper
	mock := mockHelper.BuildBasicJsonGet()
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(mock)

	response, _ := http.Post(fmt.Sprintf("%s%s", host, "/create"), "application/json; charset=utf-8", buffer)
	assert.Equal(t, 200, response.StatusCode)
	var output model.Mock
	json.NewDecoder(response.Body).Decode(&output)
	fmt.Println(output)
	teardown()

}

func TestAllFlow(t *testing.T) {
	setup()
	response, _ := http.Get(fmt.Sprintf("%s%s", host, "/json-get"))
	assert.Equal(t, 404, response.StatusCode)

	var mockHelper helper.MockHelper
	mock := mockHelper.BuildBasicJsonGet()
	jsonString := new(bytes.Buffer)
	json.NewEncoder(jsonString).Encode(mock)

	response, _ = http.Post(fmt.Sprintf("%s%s", host, "/create"), "application/json; charset=utf-8", jsonString)
	assert.Equal(t, 200, response.StatusCode)
	var output model.Mock
	json.NewDecoder(response.Body).Decode(&output)

	response, _ = http.Get(fmt.Sprintf("%s%s", host, "/json-get"))
	assert.Equal(t, 200, response.StatusCode)
	teardown()
}
