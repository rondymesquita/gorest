package helper

import (
	"gorest/model"
)

type MockHelper struct {
}

func (mockHelper *MockHelper) BuildBasicJsonGet() model.Mock {
	var mock model.Mock
	mock.Path = "/json-get"
	mock.HttpMethod = "get"
	var response model.Response
	response.Type = "json"
	response.Data = "{'id':'1', 'email':'email@email.com'}"
	mock.Response = response
	return mock
}
