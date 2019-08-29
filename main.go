package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)
import "github.com/kostyashevchuk/dockerify"

func main() {
	fmt.Print("Hello")
	client := Swagger.Client{Server: "http://localhost:8080", Client: http.Client{}}
	response, err := client.CreateApp(context.Background(), Swagger.CreateAppJSONRequestBody{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Body)
}
