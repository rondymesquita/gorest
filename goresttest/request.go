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

func Post(uri string, payload *bytes.Buffer) (*http.Response, model.ResponseData){
	url := url(uri)
	response, _ := http.Post(url, "application/json; charset=utf-8", payload)

	var responseData model.ResponseData
	json.NewDecoder(response.Body).Decode(&responseData)
	return response, responseData
}


func Get(uri string) (*http.Response, string){
	url := url(uri)
	response, _ := http.Get(url)
	responseBody, _ := ioutil.ReadAll(response.Body)

	responseDataString := strings.TrimSpace(string(responseBody))

	return response, responseDataString
}


func url(uri string) string{
	return fmt.Sprintf("%s%s", host, uri)
}