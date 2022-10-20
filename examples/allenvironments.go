package main

import (
	"fmt"
	"postman"
	"postman/request/allenvironments"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.AllEnvironments(allenvironments.WithWorkspaceId("your workspace id"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
