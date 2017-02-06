package model

import (
	"strings"
)

type Mock struct {
	Path       string `json:"Path" binding:"required"`
	HttpMethod string `json:"HttpMethod" binding:"required"`
	Response   Response `json:"Response" binding:"required"`
}

func (mock *Mock) String() string {
	s := []string{"Path", mock.Path, "HttpMethod", mock.HttpMethod, "Response", mock.Response.String()}
	return strings.Join(s, " ")
}

type Response struct {
	Body string `json:"Body" binding:"required"`
	Type string `json:"Type" binding:"required"`
	//StatusCode int `json:"StatusCode" binding:"required"`
	//Headers map[string][]string `json:"Headers" binding:"required"`
}

func (response *Response) String() string {
	s := []string{"Body", response.Body, "Type", response.Type}
	return strings.Join(s, " ")
}

type ResponseMessage struct {
	Message string
	Status  string
}
