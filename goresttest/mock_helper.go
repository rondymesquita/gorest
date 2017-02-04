package goresttest

import (
	"github.com/rondymesquita/gorest/model"
	"bytes"
	"encoding/json"
)

type MockHelper struct {
}

func (mockHelper *MockHelper) BuildJsonGet() (model.Mock, *bytes.Buffer) {
	var mock model.Mock
	mock.Path = "/json-get"
	mock.HttpMethod = "get"
	var response model.Response
	response.Type = "json"
	response.Body = "{'id':'1', 'email':'email@email.com'}"
	mock.Response = response

	mockJson := new(bytes.Buffer)
	json.NewEncoder(mockJson).Encode(mock)

	return mock, mockJson
}

func (mockHelper *MockHelper) BuildJsonGetWith(statusCode int) (model.Mock, *bytes.Buffer) {
	var mock model.Mock
	mock.Path = "/json-get"
	mock.HttpMethod = "post"
	var response model.Response
	response.Type = "json"
	response.Body = "{'id':'1', 'email':'email@email.com'}"
	response.StatusCode = statusCode
	mock.Response = response

	mockJson := new(bytes.Buffer)
	json.NewEncoder(mockJson).Encode(mock)

	return mock, mockJson
}
