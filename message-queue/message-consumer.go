package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type MessageConsumer struct {
	Host         string
	Username     string
	Password     string
	QueueName    string
	ExchangeName string
	Connection   *amqp.Connection
	Channel      *amqp.Channel
	Queue        *amqp.Queue
}

func NewMessageConsumer(host string, username string, password string, queueName string) *MessageConsumer {
	return &MessageConsumer{
		Host:         host,
		Username:     username,
		Password:     password,
		QueueName:    queueName,
		ExchangeName: "",
	}
}

func (consumer *MessageConsumer) createChannel() {
	url := fmt.Sprintf("amqp://%s:%s@%s", consumer.Username, consumer.Password, consumer.Host)

	connection, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")

	consumer.Connection = connection
	consumer.Channel = channel
}

func (consumer *MessageConsumer) declareQueue() *amqp.Queue {
	queue, err := consumer.Channel.QueueDeclare(
		consumer.QueueName, // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return &queue
}

func (consumer *MessageConsumer) Initialize() {
	consumer.createChannel()
	consumer.declareQueue()
}

func (consumer *MessageConsumer) Terminate() {
	consumer.Channel.Close()
	consumer.Connection.Close()
}

func (consumer *MessageConsumer) ReceiveMessage(done chan bool) {
	delivery, err := consumer.Channel.Consume(
		consumer.QueueName, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	failOnError(err, "Failed to receive a message")

	go func() {
		for d := range delivery {
			log.Printf("Receive message `%s` from queue `%s`", d.Body, consumer.QueueName)

			done <- true
		}
	}()
}
