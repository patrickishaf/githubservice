package common

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var baseUrl string = "https://api.github.com"

func HttpGet(urlPath string) ([]byte, error) {
	url := baseUrl + urlPath
	response, err := http.Get(url)

	if err != nil {
		log.Println("failed to make HttpGet request:", err.Error())
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		log.Println("failed tp parse response json")
		return nil, err
	}

	fmt.Printf("GET request successful")

	return responseBody, nil
}
