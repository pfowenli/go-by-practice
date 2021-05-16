package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type CustomResponse struct {
	Status string `json:"status"`
	Text   string `json:"text"`
}

func GetRequest(url string) CustomResponse {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Get(url)
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

func PostRequest(url string, data map[string]interface{}) CustomResponse {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(jsonData)

	res, err := client.Post(url, "application/json", buffer)
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
