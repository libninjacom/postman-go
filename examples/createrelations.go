package main

import (
	"fmt"
	"postman"
	"postman/request/createrelations"
)

func main() {
	client := postman.NewClientFromEnv()
	apiId := "387c2863-6ee3-4a56-8210-225f774edade"
	apiVersionId := "a9879d02-74bf-425a-bbec-6d27aa135507"
	res, err := client.CreateRelations(apiId, apiVersionId, createrelations.WithDocumentation([1]string{"your documentation"}), createrelations.WithEnvironment([1]string{"your environment"}), createrelations.WithMock([1]string{"your mock"}), createrelations.WithMonitor([1]string{"your monitor"}), createrelations.WithTest([1]string{"your test"}), createrelations.WithContracttest([1]string{"your contracttest"}), createrelations.WithTestsuite([1]string{"your testsuite"}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
