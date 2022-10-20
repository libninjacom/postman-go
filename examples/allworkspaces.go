package main

import (
	"fmt"
	"postman"
	"postman/request/allworkspaces"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.AllWorkspaces(allworkspaces.WithType("team"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
