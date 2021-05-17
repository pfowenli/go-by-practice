package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type CustomResponse struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}
type CustomHttpClient struct {
	client *http.Client
}

func (c *CustomHttpClient) GetRequest(url string, data map[string]interface{}) CustomResponse {
	pairs := []string{}
	for field, value := range data {
		pair := field + "=" + fmt.Sprintf("%v", value)
		pairs = append(pairs, pair)
	}
	queryString := strings.Join(pairs, "&")

	res, err := c.client.Get(url + "?" + queryString)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return CustomResponse{
		Status: res.Status,
		Text:   string(body),
	}
}

func (c *CustomHttpClient) PostRequest(url string, data map[string]interface{}) CustomResponse {
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(jsonData)

	res, err := c.client.Post(url, "application/json", buffer)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return CustomResponse{
		Status: res.Status,
		Text:   string(body),
	}
}

func NewCustomHttpClient() *CustomHttpClient {
	return &CustomHttpClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
