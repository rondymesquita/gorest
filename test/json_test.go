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
	"io/ioutil"
	"strings"
	"strconv"
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

func TestCreateRouteShouldExists(t *testing.T) {
	setup()
	var mockHelper helper.MockHelper
	mock := mockHelper.BuildBasicJsonGet()
	jsonString := new(bytes.Buffer)
	json.NewEncoder(jsonString).Encode(mock)

	response, _ := http.Post(fmt.Sprintf("%s%s", host, "/create"), "application/json; charset=utf-8", jsonString)
	assert.Equal(t, 200, response.StatusCode)
	teardown()
}

func TestCreateAMockRouteWithJsonReturnAndGetMethod(t *testing.T) {
	setup()
	response, _ := http.Get(fmt.Sprintf("%s%s", host, "/json-get"))
	assert.Equal(t, 404, response.StatusCode)

	var mockHelper helper.MockHelper
	mock := mockHelper.BuildBasicJsonGet()
	jsonString := new(bytes.Buffer)
	json.NewEncoder(jsonString).Encode(mock)
	response, _ = http.Post(fmt.Sprintf("%s%s", host, "/create"), "application/json; charset=utf-8", jsonString)

	var responseData model.ResponseData
	json.NewDecoder(response.Body).Decode(&responseData)
	responseDataExpected := model.ResponseData{Message: "Route created with success.", Status: "SUCCESS"}

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, responseDataExpected, responseData)
	assert.Equal(t, "application/json; charset=utf-8", response.Header.Get("Content-Type"))

	response, _ = http.Get(fmt.Sprintf("%s%s", host, "/json-get"))
	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, strconv.Quote(mock.Response.Data), strings.TrimSpace(string(responseBody)))

	teardown()
}
