package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	apiId := "387c2863-6ee3-4a56-8210-225f774edade"
	apiVersionId := "a9879d02-74bf-425a-bbec-6d27aa135507"
	relationType := "your relation type"
	entityId := "e3d951bf-873f-49ac-a658-b2dcb91d3289"
	res, err := client.SyncRelationsWithSchema(apiId, apiVersionId, relationType, entityId)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
