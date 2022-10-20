package main

import (
	"fmt"
	"postman"
	"postman/request/createcollectionfromschema"
)

func main() {
	client := postman.NewClientFromEnv()
	apiId := "387c2863-6ee3-4a56-8210-225f774edade"
	apiVersionId := "a9879d02-74bf-425a-bbec-6d27aa135507"
	schemaId := "16bb367e-fafb-4ef3-933b-ee3d971866fb"
	name := "your name"
	relations := [1]interface{}{map[string]string{}}
	res, err := client.CreateCollectionFromSchema(apiId, apiVersionId, schemaId, createcollectionfromschema.WithWorkspaceId("your workspace id"), name, relations)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
