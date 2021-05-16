package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	res, err := http.Get("http://localhost:8080/api/v1/hello-world")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Printf("HTTP response status: %v\n", res.Status)
	fmt.Printf("HTTP response body: %v\n", string(body))
}
