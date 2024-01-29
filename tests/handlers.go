package tests

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMessagesHandler(c *gin.Context) {
	c.String(http.StatusOK, "SND 1 2022-10-01T06:35:01.373Z 192.168.2.22:8000\nSND 1 2022-10-01T06:35:01.373Z 192.168.2.22:8000 192.168.2.21:78390")
}

func PutStateHandlerSuccess(c *gin.Context) {
	c.String(http.StatusOK, "Successfully Updated State")
}

func PutStateHandlerFailure(c *gin.Context) {
	c.String(http.StatusBadRequest, "Invalid State Value")
}

func GetStateHandler(c *gin.Context) {
	c.String(http.StatusOK, "RUNNING")
}

func GetRunLogHandler(c *gin.Context) {
	c.String(http.StatusOK, "2023-11-01T06.35:01.380Z: INIT->RUNNING")
}

