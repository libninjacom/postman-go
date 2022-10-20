package main

import (
	"fmt"
	"postman"
	"postman/request/updatecollection"
)

func main() {
	client := postman.NewClientFromEnv()
	collectionUid := "12ece9e1-2abf-4edc-8e34-de66e74114d2"
	res, err := client.UpdateCollection(collectionUid, updatecollection.WithCollection(map[string]string{}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
