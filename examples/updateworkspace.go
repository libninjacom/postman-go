package main

import (
	"fmt"
	"postman"
	"postman/request/updateworkspace"
)

func main() {
	client := postman.NewClientFromEnv()
	workspaceId := "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9"
	res, err := client.UpdateWorkspace(workspaceId, updateworkspace.WithWorkspace(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
