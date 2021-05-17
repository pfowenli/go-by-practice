package main

import (
	"fmt"
)

func main() {
	httpClient := *NewCustomHttpClient()
	url := ""
	data := map[string]interface{}{}
	response := CustomResponse{}

	url = "http://localhost:8080/api/v1/hello-world"
	response = httpClient.GetRequest(url, data)

	fmt.Printf("url: %s\n", url)
	fmt.Printf("HTTP status: %s\n", response.Status)
	fmt.Printf("text: %s\n", response.Text)

	url = "http://localhost:8080/api/v2/query"
	data = map[string]interface{}{"id": 6, "type": "Gorilla"}
	response = httpClient.GetRequest(url, data)

	fmt.Printf("url: %s\n", url)
	fmt.Printf("HTTP status: %s\n", response.Status)
	fmt.Printf("text: %s\n", response.Text)

	url = "http://localhost:8080/api/v2/json-body"
	data = map[string]interface{}{"id": 6, "type": "Gorilla"}
	response = httpClient.PostRequest(url, data)

	fmt.Printf("url: %s\n", url)
	fmt.Printf("HTTP status: %s\n", response.Status)
	fmt.Printf("text: %s\n", response.Text)
}
