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
	"net/http/httptest"
)

var host = fmt.Sprintf("http://localhost:%s", app.Port)

func DoCreateMockRequest(server *httptest.Server, payload *bytes.Buffer) (*http.Response, model.ResponseMessage){
	url := fmt.Sprintf("%s%s", server.URL, "/create")
	response, _ := http.Post(url, "application/json; charset=utf-8", payload)
	//client := &http.Client{}
	//request, _ := http.NewRequest("POST", url, payload)
	//request.Header.Set("fulano","sicrano")
	//response, _ := client.Do(request)

	var responseMessage model.ResponseMessage
	json.NewDecoder(response.Body).Decode(&responseMessage)
	return response, responseMessage
}


func Get(server *httptest.Server, uri string) (*http.Response, string){
	url := fmt.Sprintf("%s%s", server.URL, uri)
	response, _ := http.Get(url)
	responseBody, _ := ioutil.ReadAll(response.Body)
	data := strings.TrimSpace(string(responseBody))
	return response, data
}


func url(uri string) string{
	return fmt.Sprintf("%s%s", host, uri)
}