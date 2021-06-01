package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

func handleConnection(conn net.Conn) {
	clientAddress := conn.RemoteAddr().String()

	buffer, err := bufio.NewReader(conn).ReadBytes('\n')
	if err != nil {
		log.Printf("Client `%s` disconnected", clientAddress)
		conn.Close()
		return
	}
	message := string(buffer[:len(buffer)-1])

	log.Printf("Received message `%s` from client `%s`", message, clientAddress)

	conn.Write(buffer)
	log.Printf("Relayed message `%s` to client `%s`", message, clientAddress)

	handleConnection(conn)
}

func main() {
	address := fmt.Sprintf("%s:%d", hostname, port)
	listener, err := net.Listen(network, address)
	failOnError(err)
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		failOnError(err)

		clientAddress := connection.RemoteAddr().String()
		log.Printf("Client `%s` connected", clientAddress)

		go handleConnection(connection)
	}
}
