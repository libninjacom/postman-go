package main

import (
	"fmt"
	"postman"
	"postman/request/importexternalapispecification"
)

func main() {
	client := postman.NewClientFromEnv()
	body := map[string]string{}
	res, err := client.ImportExternalApiSpecification(importexternalapispecification.WithWorkspaceId("your workspace id"), body)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
