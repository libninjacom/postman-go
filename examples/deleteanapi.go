package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	apiId := "387c2863-6ee3-4a56-8210-225f774edade"
	res, err := client.DeleteAnApi(apiId)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
