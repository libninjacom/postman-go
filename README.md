<div id="top"></div>

<p align="center">
    <a href="https://github.com/libninjacom/postman-go/stargazers">
        <img src="https://img.shields.io/github/stars/libninjacom/postman-go.svg?style=flat-square" alt="Stars" />
    </a>
    <a href="https://github.com/libninjacom/postman-go/actions">
        <img src="https://img.shields.io/github/workflow/status/libninjacom/postman-go/ci?style=flat-square" alt="Build Status" />
    </a>
</p>

Postman client, generated from the OpenAPI spec.

# Usage

```go
package main

import (
	"fmt"
	"postman"
	"postman/request/getallapis"
)

func main() {
	client := postman.NewClientFromEnv()
	res, err := client.GetAllApis(
		getallapis.WithWorkspace("your workspace"),
		getallapis.WithSince("your since"),
		getallapis.WithUntil("your until"),
		getallapis.WithCreatedBy("your created by"),
		getallapis.WithUpdatedBy("your updated by"),
		getallapis.WithIsPublic(true),
		getallapis.WithName("your name"),
		getallapis.WithSummary("your summary"),
		getallapis.WithDescription("your description"),
		getallapis.WithSort("your sort"),
		getallapis.WithDirection("your direction")
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```

This example loads configuration from environment variables, specifically:

* `POSTMAN_API_KEY`

# Documentation

* [API Documentation](https://www.postman.com/postman/workspace/postman-public-workspace/documentation/12959542-c8142d51-e97c-46b6-bd77-52bb66712c9a)

You can see working examples of every API call in the `examples/` directory.

# Contributing

Contributions are welcome!

*Library created with [Libninja](https://www.libninja.com).*