package handlers

import (
	"bytes"
	"com.example.docker.compose/service1/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PutStateHandler(c *gin.Context) {
	var msgBody string
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading request body"})
		return
	}
	msgBody = buf.String()
	fmt.Println("RECEIVED STATE", msgBody)
	config.SetCurrentState(msgBody)
	fmt.Println("UPDATED STATE VARIABLE", config.CurrentState)
	c.Data(http.StatusOK, "text/plain", []byte("State updated successfully"))
}

func GetStateHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte(config.CurrentState))
}
