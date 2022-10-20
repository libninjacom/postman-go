package main

import (
	"fmt"
	"postman"
)

func main() {
	client := postman.NewClientFromEnv()
	mockUid := "e3d951bf-873f-49ac-a658-b2dcb91d3289"
	res, err := client.UnpublishMock(mockUid)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
