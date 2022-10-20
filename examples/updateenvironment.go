package main

import (
	"fmt"
	"postman"
	"postman/request/updateenvironment"
)

func main() {
	client := postman.NewClientFromEnv()
	environmentUid := "5daabc50-8451-43f6-922d-96b403b4f28e"
	res, err := client.UpdateEnvironment(environmentUid, updateenvironment.WithEnvironment(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
