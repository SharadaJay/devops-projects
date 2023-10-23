package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

type Msg struct {
	Message string `json:"message" binding:"required"`
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	servicePort := os.Getenv("SERVICE_2_PORT")
	servicePath := os.Getenv("SERVICE_2_CALL_PATH")
	msgTopic := os.Getenv("MSG_TOPIC")
	logTopic := os.Getenv("LOG_TOPIC")
	timeStampFormat := os.Getenv("TIMESTAMP_FORMAT")
	serviceName := os.Getenv("SERVICE2_SERVICE_NAME")
	rabbitMQServiceName := os.Getenv("RABBITMQ_SERVICE_NAME")

	rabbitmqIPAddress, err := getIPAddress(rabbitMQServiceName)
	connectionStr := "amqp://guest:guest@" + rabbitmqIPAddress + ":5672/"

	isConnected := false
	var conn *amqp.Connection

	for isConnected == false {
		conn, err = amqp.Dial(connectionStr)
		if err != nil {
			isConnected = false
		} else {
			isConnected = true
		}
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Failed to open a channel" + err.Error())
	}
	defer ch.Close()

	ipAddress, err := getIPAddress(serviceName)

	fmt.Println(ipAddress)

	url := "http://" + ipAddress + ":" + servicePort + servicePath

	for i := 1; i <= 20; i++ {

		currentTime := time.Now().UTC()
		formattedTime := currentTime.Format(timeStampFormat)
		message := "SND " + strconv.Itoa(i) + " " + formattedTime + " " + ipAddress + ":" + servicePort

		err := publishToRabbitMq(ch, msgTopic, message)
		if err != nil {
			fmt.Println(err)
		}

		err = callService(ch, logTopic, url, message, timeStampFormat)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(2 * time.Second)

		if i == 20 {
			stopMsg := "SND STOP"
			err := publishToRabbitMq(ch, logTopic, stopMsg)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	fmt.Println("Application is running. Press Ctrl+C to exit.")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

}

func callService(ch *amqp.Channel, logTopic string, url string, message string, timeStampFormat string) error {
	msgStruct := Msg{message}
	jsonMsgStr, _ := json.Marshal(msgStruct)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonMsgStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	code := ""

	if err != nil {
		code = "500"
		_ = publishToRabbitMq(ch, logTopic, err.Error())
	} else {
		code = strconv.Itoa(resp.StatusCode)
	}

	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format(timeStampFormat)
	respMsg := code + " " + formattedTime

	err = publishToRabbitMq(ch, logTopic, respMsg)
	if err != nil {
		return err
	}

	return nil
}

func getIPAddress(serviceName string) (string, error) {
	address, err := net.LookupHost(serviceName)
	fmt.Println(address)
	if err != nil {
		return "", err
	}
	if len(address) == 0 {
		return "", fmt.Errorf("no IP address found for container: %s", serviceName)
	}
	return address[0], nil
}

func publishToRabbitMq(ch *amqp.Channel, queueName string, message string) error {

	_, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
