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
	response.Data = "{'id':'1', 'email':'email@email.com'}"
	mock.Response = response

	mockJson := new(bytes.Buffer)
	json.NewEncoder(mockJson).Encode(mock)

	return mock, mockJson
}
