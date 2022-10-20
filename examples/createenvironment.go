package main

import (
	"fmt"
	"postman"
	"postman/request/createenvironment"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.CreateEnvironment(createenvironment.WithWorkspaceId("your workspace id"), createenvironment.WithEnvironment(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
