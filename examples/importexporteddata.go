package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.ImportExportedData()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
