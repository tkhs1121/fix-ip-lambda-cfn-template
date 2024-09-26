package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() (string, error) {
	url := os.Getenv("URL")

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	responseBody := string(body)
	fmt.Println("Reponse from", url, ":", responseBody)

	return responseBody, nil
}

func main() {
	lambda.Start(handler)
}