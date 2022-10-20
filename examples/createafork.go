package main

import (
	"fmt"
	"postman"
	"postman/request/createafork"
)

func main() {
	client := postman.NewClientFromEnv()
	workspace := "1f0df51a-8658-4ee8-a2a1-d2567dfa09a9"
	collectionUid := "12ece9e1-2abf-4edc-8e34-de66e74114d2"
	res, err := client.CreateAFork(workspace, collectionUid, createafork.WithLabel("your label"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
