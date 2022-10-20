package main

import (
	"fmt"
	"postman"
	"postman/request/allcollections"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.AllCollections(allcollections.WithWorkspaceId("your workspace id"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
