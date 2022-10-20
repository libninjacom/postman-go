package main

import (
	"fmt"
	"postman"
	"postman/request/createapi"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.CreateApi(createapi.WithWorkspaceId("your workspace id"), createapi.WithApi(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
