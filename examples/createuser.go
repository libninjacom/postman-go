package main

import (
	"fmt"
	"postman"
	"postman/request/createuser"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.CreateUser(createuser.WithSchemas([1]string{"your schemas"}), createuser.WithUserName("your user name"), createuser.WithActive(true), createuser.WithExternalId("your external id"), createuser.WithGroups([1]string{"your groups"}), createuser.WithLocale("your locale"), createuser.WithName(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
