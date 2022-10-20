package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	workspaceId := "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9"
	res, err := client.SingleWorkspace(workspaceId)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
