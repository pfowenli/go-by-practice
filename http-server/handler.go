package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomData struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world\n"))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for field, values := range req.Header {
		for _, value := range values {
			log.Printf("%v: %v\n", field, value)
			fmt.Fprintf(w, "%v: %v\n", field, value)
		}
	}
}

func pathParameters(w http.ResponseWriter, req *http.Request) {
	log.Printf("path parameters")

	for field, value := range mux.Vars(req) {
		log.Printf("%v: %v", field, value)
		fmt.Fprintf(w, "%v: %v\n", field, value)
	}
}

func queryStrings(w http.ResponseWriter, req *http.Request) {
	log.Printf("query strings")

	for field, values := range req.URL.Query() {
		for _, value := range values {
			log.Printf("%v: %v", field, value)
			fmt.Fprintf(w, "%v: %v\n", field, value)
		}
	}
}

func jsonBody(w http.ResponseWriter, req *http.Request) {
	log.Printf("json body")

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	data := CustomData{}
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if _, err := json.Marshal(data); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Printf("Id: %v", data.Id)
	log.Printf("Type: %v", data.Type)

	fmt.Fprintf(w, "Id: %v\n", data.Id)
	fmt.Fprintf(w, "Type: %v\n", data.Type)
}
