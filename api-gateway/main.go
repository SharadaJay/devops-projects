package main

import (
	"example.com/api-gateway/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/messages", handlers.GetMessagesHandler)
	r.PUT("/state", handlers.PutStateHandler)

	if err := r.Run(":8083"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
