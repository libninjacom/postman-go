package main

import (
	"fmt"
	"postman"
	"postman/request/createmonitor"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.CreateMonitor(createmonitor.WithWorkspaceId("your workspace id"), createmonitor.WithMonitor(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
