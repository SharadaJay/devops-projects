package config

import (
	"fmt"
	"github.com/streadway/amqp"
)

var CurrentState string
var RabbitMQChannel *amqp.Channel
var RunLogTopic string
var TimeStampFormat string

func SetCurrentState(value string) {
	CurrentState = value
}

func SetRabbitMQChannel(value *amqp.Channel) {
	RabbitMQChannel = value
}

func SetRunLogTopic(value string) {
	RunLogTopic = value
}

func SetTimeStampFormat(value string) {
	TimeStampFormat = value
}

func PublishToRabbitMq(ch *amqp.Channel, queueName string, message string) error {

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
