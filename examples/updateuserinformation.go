package main

import (
	"fmt"
	"postman"
	"postman/request/updateuserinformation"
)

func main() {
	client := postman.NewClientFromEnv()
	userId := "405775fe15ed41872a8eea4c8aa2b38cda9749812cc55c99"
	res, err := client.UpdateUserInformation(userId, updateuserinformation.WithSchemas([1]string{"your schemas"}), updateuserinformation.WithName(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
