package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	apiId := "387c2863-6ee3-4a56-8210-225f774edade"
	apiVersionId := "a9879d02-74bf-425a-bbec-6d27aa135507"
	res, err := client.GetContractTestRelations(apiId, apiVersionId)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
