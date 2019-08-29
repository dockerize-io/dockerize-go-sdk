package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
import "github.com/kostyashevchuk/dockerify/Swagger"

func main() {
	token := ""
	client := Swagger.Client{Server: "https://api.dockerize.io", Client: http.Client{},
		RequestEditor: func(req *http.Request, ctx context.Context) error {
			req.Header.Add("Authorization", "Bearer "+token)
			return nil
		}}
	response, err := client.GetDockerCredentials(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Status)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
