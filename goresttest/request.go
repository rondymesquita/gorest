package goresttest

import (
	"github.com/rondymesquita/gorest/model"
	"github.com/rondymesquita/gorest/app"
	"encoding/json"
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

var host = fmt.Sprintf("http://localhost:%s", app.Port)

func DoCreateMockRequest(payload *bytes.Buffer) (*http.Response, model.ResponseMessage){
	url := url("/create")
	response, _ := http.Post(url, "application/json; charset=utf-8", payload)

	var responseMessage model.ResponseMessage
	json.NewDecoder(response.Body).Decode(&responseMessage)
	return response, responseMessage
}


func Get(uri string) (*http.Response, string){
	url := url(uri)
	response, _ := http.Get(url)
	responseBody, _ := ioutil.ReadAll(response.Body)
	data := strings.TrimSpace(string(responseBody))
	return response, data
}


func url(uri string) string{
	return fmt.Sprintf("%s%s", host, uri)
}