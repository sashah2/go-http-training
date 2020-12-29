package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-http-training/routes"
)

func main() {
	server := gin.Default()

	for _, e := range routes.Endpoints {
		server.Handle(e.Type, e.Path, e.Handler)
	}
	err := server.Run(":4040")

	if err != nil {
		fmt.Println(err)
	}
}
