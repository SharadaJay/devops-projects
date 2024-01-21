package handlers

import (
	"example.com/api-gateway/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GetMessagesHandler(c *gin.Context) {

	monitorURL := config.MonitorURL
	resp, err := http.Get(monitorURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(resp.StatusCode, "text/plain; charset=utf-8", body)
}
