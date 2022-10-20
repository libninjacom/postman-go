package main

import (
	"fmt"
	"postman"
	"postman/request/createapiversion"
)

func main() {
	client := postman.NewClientFromEnv()
	apiId := "387c2863-6ee3-4a56-8210-225f774edade"
	res, err := client.CreateApiVersion(apiId, createapiversion.WithVersion(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
