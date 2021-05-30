package main

import (
	"time"
)

const (
	host     string = "localhost:5672"
	username string = "guest"
	password string = "guest"
	queue    string = "beverage"
)

func main() {
	messages := []string{"Black Tea", "Americano", "Latte", "Lemonade"}

	consumer := NewMessageConsumer(host, username, password, queue)
	consumer.Initialize()
	defer consumer.Terminate()

	producer := NewMessageProducer(host, username, password, queue)
	producer.Initialize()
	defer producer.Terminate()

	done := make(chan bool, len(messages))
	go consumer.ReceiveMessage(done)

	for index, message := range messages {
		time.Sleep(time.Duration(index) * time.Second)
		producer.SendPlainTextMessage(message)
	}

	<-done
}
