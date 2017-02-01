package goresttest

import (
	"github.com/rondymesquita/gorest/app"
	"testing"
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/stretchr/testify/assert"
	"github.com/rondymesquita/gorest/model"
	"github.com/rondymesquita/gorest/goresttest"
	"strconv"
	"github.com/rondymesquita/gorest/util"
)

var application app.App

//func TestMain(m *testing.M) {
//	setup()
//	retCode := m.Run()
//	teardown()
//	os.Exit(retCode)
//}

func setup() {
	gin.SetMode(gin.TestMode)
	application.Create()
	go application.Start()

}

func teardown() {
	application.Stop()
}

func TestCreateRouteShouldExists(t *testing.T) {
	setup()
	var mockHelper goresttest.MockHelper
	_, mockJson := mockHelper.BuildJsonGet()

	response, _ := goresttest.DoCreateMockRequest(mockJson)
	assert.Equal(t, 200, response.StatusCode)
	teardown()
}

func TestCreateAMockRouteWithJsonReturnAndGetMethod(t *testing.T) {
	setup()

	var mockHelper goresttest.MockHelper
	mock, mockJson := mockHelper.BuildJsonGet()

	resp, _ := goresttest.Get(mock.Path)
	assert.Equal(t, 404, resp.StatusCode)

	response, responseMessage := goresttest.DoCreateMockRequest(mockJson)
	responseMessageExpected := model.ResponseMessage{Message: util.RouteCreatedWithSuccess, Status: util.Success}
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, responseMessageExpected, responseMessage)
	assert.Equal(t, "application/json; charset=utf-8", response.Header.Get("Content-Type"))

	response, data := goresttest.Get("/json-get")
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, strconv.Quote(mock.Response.Data), data)
	assert.Equal(t, "application/json; charset=utf-8", response.Header.Get("Content-Type"))

	teardown()
}
