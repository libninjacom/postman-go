package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.AllMocks()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
