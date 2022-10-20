package main

import (
	"fmt"
	"postman"
	"postman/request/createmock"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.CreateMock(createmock.WithWorkspaceId("your workspace id"), createmock.WithMock(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
