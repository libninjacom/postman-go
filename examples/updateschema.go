package main

import (
	"fmt"
	"postman"
	"postman/request/updateschema"
)

func main() {
	client := postman.NewClientFromEnv()
	apiId := "387c2863-6ee3-4a56-8210-225f774edade"
	apiVersionId := "a9879d02-74bf-425a-bbec-6d27aa135507"
	schemaId := "16bb367e-fafb-4ef3-933b-ee3d971866fb"
	res, err := client.UpdateSchema(apiId, apiVersionId, schemaId, updateschema.WithSchema(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
