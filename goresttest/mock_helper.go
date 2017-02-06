package goresttest

import (
	"github.com/rondymesquita/gorest/model"
	"bytes"
	"encoding/json"
	"fmt"
)

type MockHelper struct {
}

func (mockHelper *MockHelper) BuildJsonGet() (model.Mock, *bytes.Buffer) {
	response := model.Response{
		Type: "json",
		Body: "{'id':'1', 'email':'email@email.com'}",
		Headers: map[string][]string{
			"Accept-Encoding": {"gzip, deflate"},
			"Accept-Language": {"en-us"},
			"Foo": {"Bar", "two"},
		},
		StatusCode: 200,
	}
	mock := model.Mock{
		Path: "/json-get",
		HttpMethod : "get",
		Response: response,
	}

	mockJson := new(bytes.Buffer)
	json.NewEncoder(mockJson).Encode(mock)
	fmt.Println(mockJson)

	return mock, mockJson
}