package main

import "log"

func main() {
	message := "Black Tea"
	log.Printf("message: %s", message)

	producer := NewMessageProducer(
		"localhost:5672",
		"guest",
		"guest",
		"beverage",
	)
	producer.Initialize()
	defer producer.Terminate()

	producer.SendPlainTextMessage(message)
}
