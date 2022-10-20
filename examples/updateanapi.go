package main

import (
	"fmt"
	"postman"
	"postman/request/updateanapi"
)

func main() {
	client := postman.NewClientFromEnv()
	apiId := "387c2863-6ee3-4a56-8210-225f774edade"
	res, err := client.UpdateAnApi(apiId, updateanapi.WithApi(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
