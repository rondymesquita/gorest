package model

import (
	"strings"
)

type Mock struct {
	Path       string
	HttpMethod string
	Response   Response
}

func (mock *Mock) String() string {
	s := []string{"Path", mock.Path, "HttpMethod", mock.HttpMethod, "Response", mock.Response.String()}
	return strings.Join(s, " ")
}

type Response struct {
	Body string
	Type string
	StatusCode int
}

func (response *Response) String() string {
	s := []string{"Data", response.Body, "Type", response.Type}
	return strings.Join(s, " ")
}

type ResponseMessage struct {
	Message string
	Status  string
}
