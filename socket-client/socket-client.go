package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	network  string = "tcp"
	hostname string = "127.0.0.1"
	port     int    = 80
)

func failOnError(err error) {
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}

func main() {
	address := fmt.Sprintf("%s:%d", hostname, port)
	connection, err := net.Dial(network, address)
	failOnError(err)
	defer connection.Close()
	log.Printf("Connected to server %s", connection.RemoteAddr().String())

	reader := bufio.NewReader(os.Stdin)

	for {
		log.Printf("Enter message here: ")
		buffer, err := reader.ReadBytes('\n')
		failOnError(err)
		connection.Write(buffer)

		message, err := bufio.NewReader(connection).ReadString('\n')
		failOnError(err)
		log.Printf("Server relayed message: %s", message)
	}
}
