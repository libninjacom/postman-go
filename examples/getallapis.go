package main

import (
	"fmt"
	"postman"
	"postman/request/getallapis"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.GetAllApis(getallapis.WithWorkspace("your workspace"), getallapis.WithSince("your since"), getallapis.WithUntil("your until"), getallapis.WithCreatedBy("your created by"), getallapis.WithUpdatedBy("your updated by"), getallapis.WithIsPublic(true), getallapis.WithName("your name"), getallapis.WithSummary("your summary"), getallapis.WithDescription("your description"), getallapis.WithSort("your sort"), getallapis.WithDirection("your direction"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
