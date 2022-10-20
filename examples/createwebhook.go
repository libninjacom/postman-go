package main

import (
	"fmt"
	"postman"
	"postman/request/createwebhook"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.CreateWebhook(createwebhook.WithWorkspaceId("your workspace id"), createwebhook.WithWebhook(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
