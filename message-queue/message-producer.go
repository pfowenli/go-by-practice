package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type MessageProducer struct {
	Host         string
	Username     string
	Password     string
	QueueName    string
	ExchangeName string
	Connection   *amqp.Connection
	Channel      *amqp.Channel
	Queue        *amqp.Queue
}

func NewMessageProducer(host string, username string, password string, queueName string) *MessageProducer {
	return &MessageProducer{
		Host:         host,
		Username:     username,
		Password:     password,
		QueueName:    queueName,
		ExchangeName: "",
	}
}

func (producer *MessageProducer) createChannel() {
	url := fmt.Sprintf("amqp://%s:%s@%s", producer.Username, producer.Password, producer.Host)

	connection, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	channel, err := connection.Channel()
	failOnError(err, "Failed to open a channel")

	producer.Connection = connection
	producer.Channel = channel
}

func (producer *MessageProducer) declareQueue() *amqp.Queue {
	queue, err := producer.Channel.QueueDeclare(
		producer.QueueName, // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return &queue
}

func (producer *MessageProducer) Initialize() {
	producer.createChannel()
	producer.declareQueue()
}

func (producer *MessageProducer) Terminate() {
	producer.Channel.Close()
	producer.Connection.Close()
}

func (producer *MessageProducer) SendPlainTextMessage(message string) {
	publishing := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(message),
	}

	err := producer.Channel.Publish(
		producer.ExchangeName, // exchange
		producer.QueueName,    // routing key
		false,                 // mandatory
		false,                 // immediate
		publishing,            // publishing
	)
	failOnError(err, "Failed to publish a message")
	log.Printf("Sent message `%s` to queue `%s`", message, producer.QueueName)
}
