package handlers

import (
	"bytes"
	"com.example.docker.compose/service1/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	if msgBody == "INIT" || msgBody == "RUNNING" || msgBody == "PAUSED" || msgBody == "SHUTDOWN" {
		previousState := config.CurrentState
		config.SetCurrentState(msgBody)
		fmt.Println("UPDATED STATE VARIABLE", config.CurrentState)
		c.Data(http.StatusOK, "text/plain", []byte("State updated successfully"))
		if previousState != config.CurrentState {
			currentTime := time.Now().UTC()
			formattedTime := currentTime.Format(config.TimeStampFormat)
			message := formattedTime + ": " + previousState + "->" + config.CurrentState
			err := config.PublishToRabbitMq(config.RabbitMQChannel, config.RunLogTopic, message)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	} else {
		c.Data(http.StatusBadRequest, "text/plain", []byte("Invalid State Value"))
		return
	}
}

func GetStateHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte(config.CurrentState))
}
