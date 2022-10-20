package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	userId := "405775fe15ed41872a8eea4c8aa2b38cda9749812cc55c99"
	res, err := client.FetchUserResource(userId)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
