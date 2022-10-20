package main

import (
	"fmt"
	"postman"
	"postman/request/schemasecurityvalidation"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.SchemaSecurityValidation(schemasecurityvalidation.WithSchema(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
