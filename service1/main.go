package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
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
	logPath := os.Getenv("LOG_PATH")
	timeStampFormat := os.Getenv("TIMESTAMP_FORMAT")
	serviceName := os.Getenv("SERVICE2_SERVICE_NAME")

	ipAddress, err := getIPAddress(serviceName)

	fmt.Println(ipAddress)

	url := "http://" + ipAddress + ":" + servicePort + servicePath

	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, os.FileMode(0666))
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	for i := 1; i <= 20; i++ {

		currentTime := time.Now().UTC()
		formattedTime := currentTime.Format(timeStampFormat)
		logMsg := strconv.Itoa(i) + " " + formattedTime + " " + ipAddress + ":" + servicePort

		err := writeLog(file, logMsg)
		if err != nil {
			log.Panic(err)
		}

		err = callService(file, i, url, logMsg)
		if err != nil {
			log.Panic(err)
		}
		time.Sleep(2 * time.Second)

		if i == 20 {
			stopMsg := "STOP"
			err := writeLog(file, stopMsg)
			if err != nil {
				log.Panic(err)
			}

			err = callService(file, i, url, stopMsg)
			if err != nil {
				log.Panic(err)
			}
			err = file.Close()
			if err != nil {
				os.Exit(0)
			}
			os.Exit(0)
		}

	}

}

func callService(file *os.File, i int, url string, logMsg string) error {
	msgStruct := Msg{logMsg}
	jsonMsgStr, _ := json.Marshal(msgStruct)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonMsgStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil && i != 20 {
		_, err2 := file.WriteString(err.Error() + "\n")
		if err2 != nil {
			return err
		}
	}
	return nil
}

func writeLog(file *os.File, logMsg string) error {
	_, err := file.WriteString(logMsg + "\n")
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
