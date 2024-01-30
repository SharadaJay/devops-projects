package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

var Service1URL string
var Service2URL string
var MonitorURL string
var RabbitMQURL string
var RabbitMQUser string
var RabbitMQPwd string

const HTTP_CONST = "http://"

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service1Port := os.Getenv("SERVICE_1_PORT")
	service1Name := os.Getenv("SERVICE1_SERVICE_NAME")
	service2Port := os.Getenv("SERVICE_2_PORT")
	service2Name := os.Getenv("SERVICE2_SERVICE_NAME")
	monitorPort := os.Getenv("MONITOR_PORT")
	monitorName := os.Getenv("MONITOR_SERVICE_NAME")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")
	rabbitMQName := os.Getenv("RABBITMQ_SERVICE_NAME")
	RabbitMQUser = os.Getenv("RABBITMQ_USER")
	RabbitMQPwd = os.Getenv("RABBITMQ_PWD")

	service1IpAddress, err := getIPAddress(service1Name)
	service2IpAddress, err := getIPAddress(service2Name)
	monitorIpAddress, err := getIPAddress(monitorName)
	rabbitMQIpAddress, err := getIPAddress(rabbitMQName)

	Service1URL = HTTP_CONST + service1IpAddress + ":" + service1Port
	Service2URL = HTTP_CONST + service2IpAddress + ":" + service2Port
	MonitorURL = HTTP_CONST + monitorIpAddress + ":" + monitorPort
	RabbitMQURL = HTTP_CONST + rabbitMQIpAddress + ":" + rabbitMQPort

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
