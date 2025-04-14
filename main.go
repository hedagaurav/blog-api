package main

import (
	"github.com/hedagaurav/blog-api/app"
)

// todo: separate the blog and user code to create separate microservices.

func main() {

	// Start the server
	if err := app.Run(); err != nil {
		panic(err)
	}
}
