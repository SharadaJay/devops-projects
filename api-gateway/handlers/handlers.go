package handlers

import (
	"encoding/base64"
	"encoding/json"
	"example.com/api-gateway/config"
	"example.com/api-gateway/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

const TEXT_PLAIN = "text/plain"

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

	c.Data(resp.StatusCode, TEXT_PLAIN, body)
}

func PutStateHandler(c *gin.Context) {

	service1PutStateUrl := config.Service1URL + "/state"
	req, err := http.NewRequest("PUT", service1PutStateUrl, c.Request.Body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", TEXT_PLAIN)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Data(resp.StatusCode, TEXT_PLAIN, body)
}

func GetStateHandler(c *gin.Context) {

	service1getStateUrl := config.Service1URL + "/state"
	resp, err := http.Get(service1getStateUrl)
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

	c.Data(resp.StatusCode, TEXT_PLAIN, body)
}

func GetRunLogHandler(c *gin.Context) {

	monitorGetRunLogURL := config.MonitorURL + "/run-log"
	resp, err := http.Get(monitorGetRunLogURL)
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

	c.Data(resp.StatusCode, TEXT_PLAIN, body)
}

func GetMQStatisticHandler(c *gin.Context) {

	auth := base64.StdEncoding.EncodeToString([]byte(config.RabbitMQUser + ":" + config.RabbitMQPwd))

	rabbitMQGetOverallURL := config.RabbitMQURL + "/api/overview"
	rabbitMQGetQueuesURL := config.RabbitMQURL + "/api/queues"

	req, err := http.NewRequest("GET", rabbitMQGetOverallURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error creating GetOverall request": err.Error()})
		return
	}

	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error making GetOverall API call": err.Error()})
		return
	}
	defer resp.Body.Close()

	var overallStat models.OverallStat
	err = json.NewDecoder(resp.Body).Decode(&overallStat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error reading response body for GetOverall": err.Error()})
		return
	}

	// get queue statistics call
	req2, err := http.NewRequest("GET", rabbitMQGetQueuesURL, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error creating GetQueues request": err.Error()})
		return
	}

	req2.Header.Set("Authorization", "Basic "+auth)

	client2 := &http.Client{}
	resp2, err := client2.Do(req2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error making GetQueues API call": err.Error()})
		return
	}
	defer resp2.Body.Close()

	var queueStat []models.QueueStat
	err = json.NewDecoder(resp2.Body).Decode(&queueStat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error reading response body for GetQueues": err.Error()})
		return
	}

	var statResponse models.StatResponse

	statResponse.OverallStats.ClusterName = overallStat.ClusterName
	statResponse.OverallStats.TotalNumberOfQueues = overallStat.ObjectTotals.Queues
	statResponse.OverallStats.MessagesDeliveredRecently = overallStat.MessageStats.DeliverGet
	statResponse.OverallStats.MessageDeliveryRate = overallStat.MessageStats.DeliverGetDetails.Rate
	statResponse.OverallStats.MessagesPublishedRecently = overallStat.MessageStats.Publish
	statResponse.OverallStats.MessagePublishingRate = overallStat.MessageStats.PublishDetails.Rate

	var queueStatsResponseArray []models.QueueStatResponse
	for _, queue := range queueStat {
		var queueStatsResponse models.QueueStatResponse
		queueStatsResponse.Name = queue.Name
		queueStatsResponse.MessagesDeliveredRecently = queue.MessageStats.DeliverGet
		queueStatsResponse.MessageDeliveryRate = queue.MessageStats.DeliverGetDetails.Rate
		queueStatsResponse.MessagesPublishedRecently = queue.MessageStats.Publish
		queueStatsResponse.MessagePublishingRate = queue.MessageStats.PublishDetails.Rate
		queueStatsResponseArray = append(queueStatsResponseArray, queueStatsResponse)
	}

	statResponse.QueueStats = queueStatsResponseArray

	c.JSON(http.StatusOK, statResponse)
}
