package main

import (
	"fmt"
	"postman"
	"postman/request/updatemonitor"
)

func main() {
	client := postman.NewClientFromEnv()
	monitorUid := "1e6b6cc1-c760-48e0-968f-4bfaeeae9af1"
	res, err := client.UpdateMonitor(monitorUid, updatemonitor.WithMonitor(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
