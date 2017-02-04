package test

import (
	"github.com/rondymesquita/gorest/app"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/rondymesquita/gorest/goresttest"
	"net/http/httptest"
	"github.com/rondymesquita/gorest/model"
	"strconv"
	"github.com/rondymesquita/gorest/constant"
)

var application app.App

func TestCreateRouteShouldExists(t *testing.T) {

	a := application.Create()
	server := httptest.NewServer(a)
	defer server.Close()

	var mockHelper goresttest.MockHelper
	_, mockJson := mockHelper.BuildJsonGet()

	response, _ := goresttest.DoCreateMockRequest(server, mockJson)
	assert.Equal(t, 201, response.StatusCode)
}

func TestCreateAMockRouteWithJsonReturnAndGetMethod(t *testing.T) {

	a := application.Create()
	server := httptest.NewServer(a)
	defer server.Close()

	var mockHelper goresttest.MockHelper
	mock, mockJson := mockHelper.BuildJsonGet()

	resp, _ := goresttest.Get(server, mock.Path)
	assert.Equal(t, 404, resp.StatusCode)

	response, responseMessage := goresttest.DoCreateMockRequest(server, mockJson)
	responseMessageExpected := model.ResponseMessage{Message: constant.RouteCreatedWithSuccess, Status: constant.Success}
	assert.Equal(t, 201, response.StatusCode)
	assert.Equal(t, responseMessageExpected, responseMessage)
	assert.Equal(t, constant.ApplicationJsonUTF8, response.Header.Get(constant.ContentType))

	response, body := goresttest.Get(server, mock.Path)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, strconv.Quote(mock.Response.Body), body)
	assert.Equal(t, constant.ApplicationJsonUTF8, response.Header.Get(constant.ContentType))

}
//
func TestCreateAMockRouteWithCustomStatusCodeAndHeaders(t *testing.T) {
}