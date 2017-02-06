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
	var responseMessage model.ResponseMessage
	json.NewDecoder(response.Body).Decode(&responseMessage)
	return response, responseMessage
}


func Get(server *httptest.Server, uri string) (*http.Response, string){
	url := fmt.Sprintf("%s%s", server.URL, uri)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	//req.Header.Add("name", "Fulano")
	//req.Header = map[string][]string{
	//			"Accept-Encoding": {"gzip, deflate"},
	//			"Accept-Language": {"en-us"},
	//			"Foo": {"Bar", "two"},
	//		}

	response, _ := client.Do(req)

	responseBody, _ := ioutil.ReadAll(response.Body)
	data := strings.TrimSpace(string(responseBody))
	return response, data
}


func url(uri string) string{
	return fmt.Sprintf("%s%s", host, uri)
}