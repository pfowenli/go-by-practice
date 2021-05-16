package main

import (
	"fmt"
)

func main() {
	url := ""
	response := CustomResponse{}

	url = "http://localhost:8080/api/v1/hello-world"
	response = GetRequest(url)

	fmt.Printf("url: %s\n", url)
	fmt.Printf("HTTP status: %s\n", response.Status)
	fmt.Printf("text: %s\n", response.Text)

	url = "http://localhost:8080/api/v2/json-body"
	data := map[string]interface{}{"Id": 6, "Type": "Gorilla"}
	response = PostRequest(url, data)

	fmt.Printf("url: %s\n", url)
	fmt.Printf("HTTP status: %s\n", response.Status)
	fmt.Printf("text: %s\n", response.Text)
}
