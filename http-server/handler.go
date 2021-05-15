package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type RequestBody struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world\n"))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for field, values := range req.Header {
		for _, value := range values {
			fmt.Printf("%v: %v\n", field, value)
			fmt.Fprintf(w, "%v: %v\n", field, value)
		}
	}
}

func pathParameters(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("path parameters\n")

	for field, value := range mux.Vars(req) {
		fmt.Printf("%v: %v\n", field, value)
		fmt.Fprintf(w, "%v: %v\n", field, value)
	}
}

func queryStrings(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("query strings\n")

	for field, values := range req.URL.Query() {
		for _, value := range values {
			fmt.Printf("%v: %v\n", field, value)
			fmt.Fprintf(w, "%v: %v\n", field, value)
		}
	}
}

func jsonBody(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello world\n")

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var data RequestBody
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if _, err := json.Marshal(data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Printf("Id: %v\n", data.Id)
	fmt.Fprintf(w, "Id: %v\n", data.Id)
	fmt.Printf("Type: %v\n", data.Type)
	fmt.Fprintf(w, "Type: %v\n", data.Type)
}
