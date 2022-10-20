package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	monitorUid := "1e6b6cc1-c760-48e0-968f-4bfaeeae9af1"
	res, err := client.RunAMonitor(monitorUid)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
