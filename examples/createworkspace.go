package main

import (
	"fmt"
	"postman"
	"postman/request/createworkspace"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.CreateWorkspace(createworkspace.WithWorkspace(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
