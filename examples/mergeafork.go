package main

import (
	"fmt"
	"postman"
	"postman/request/mergeafork"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.MergeAFork(mergeafork.WithDestination("your destination"), mergeafork.WithSource("your source"), mergeafork.WithStrategy("your strategy"))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
