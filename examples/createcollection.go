package main

import (
	"fmt"
	"postman"
	"postman/request/createcollection"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.CreateCollection(createcollection.WithWorkspaceId("your workspace id"), createcollection.WithCollection(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
