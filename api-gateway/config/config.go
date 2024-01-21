package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
	"path/filepath"
)

var Service1URL string
var Service2URL string
var MonitorURL string

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file iyrieyti")
	}

	service1Port := os.Getenv("SERVICE_1_PORT")
	service1Name := os.Getenv("SERVICE1_SERVICE_NAME")
	service2Port := os.Getenv("SERVICE_2_PORT")
	service2Name := os.Getenv("SERVICE2_SERVICE_NAME")
	monitorPort := os.Getenv("MONITOR_PORT")
	monitorName := os.Getenv("MONITOR_SERVICE_NAME")

	service1IpAddress, err := getIPAddress(service1Name)
	service2IpAddress, err := getIPAddress(service2Name)
	monitorIpAddress, err := getIPAddress(monitorName)

	Service1URL = "http://" + service1IpAddress + ":" + service1Port
	Service2URL = "http://" + service2IpAddress + ":" + service2Port
	MonitorURL = "http://" + monitorIpAddress + ":" + monitorPort

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

func dir(envFile string) string {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			break
		}

		parent := filepath.Dir(currentDir)
		if parent == currentDir {
			panic(fmt.Errorf("go.mod not found"))
		}
		currentDir = parent
	}

	return filepath.Join(currentDir, envFile)
}
