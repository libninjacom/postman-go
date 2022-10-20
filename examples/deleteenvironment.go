package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	environmentUid := "5daabc50-8451-43f6-922d-96b403b4f28e"
	res, err := client.DeleteEnvironment(environmentUid)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
