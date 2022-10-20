package main

import (
	"fmt"
	"postman"
	"postman/request/updateuserstate"
)

func main() {
	client := postman.NewClientFromEnv()
	userId := "405775fe15ed41872a8eea4c8aa2b38cda9749812cc55c99"
	res, err := client.UpdateUserState(userId, updateuserstate.WithSchemas([1]string{"your schemas"}), updateuserstate.WithOperations([1]interface{}{map[string]string{}}))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
