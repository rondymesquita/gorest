package model

import (
	"strings"
	"encoding/json"
	"bytes"
)

const(
	Json = "json"
	Get = "get"
	Post = "post"
)

type Mock struct{
	Path string "json:'path' binding:'required'"
	HttpMethod string "json:'httpMethod' binding:'required'"
	Response Response
}

func (mock *Mock) String() string{
	s := []string{"Path", mock.Path, "HttpMethod", mock.HttpMethod, "Response", mock.Response.String()}
	return strings.Join(s, " ")
}

type Response struct{
	Data string "json:'data' binding:'required'"
	Type string "json:'type' binding:'required'"
}

func (response *Response) String() string{
	s := []string{"Data", response.Data, "Type", response.Type}
	return strings.Join(s, " ")
}

type ResponseMessage struct{
	Message string "json:'message' binding:'required'"
	Status string "json:'status' binding:'required'"
}

func (responseMessage *ResponseMessage) ToJSON() *bytes.Buffer{
	responseMessageJson := new(bytes.Buffer)
	json.NewEncoder(responseMessageJson).Encode(responseMessage)
	return responseMessageJson
}