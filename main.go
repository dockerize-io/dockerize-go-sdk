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
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImtyYXNuaWtvdi0tMzU0YWZmMSIsImlhdCI6MTU2NzEwMTc0MSwiZXhwIjoxNTY3MTQ0OTQxLCJhdWQiOiJodHRwczovL2FwaS5kb2NrZXJpemUuaW8iLCJpc3MiOiJEb2NrZXJpemUgTExDIiwic3ViIjoiIn0.1HdVcssKZBwvruIhdE9mEE5jsreKHSIPXBKT080UrAg"
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
