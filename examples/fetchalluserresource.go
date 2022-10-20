package main

import (
	"fmt"
	"postman"
	"postman/request/fetchalluserresource"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.FetchAllUserResource(fetchalluserresource.WithStartIndex(1), fetchalluserresource.WithCount(50), fetchalluserresource.WithFilter("userName eq \"taylor-lee%40example.com\""))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
