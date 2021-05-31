package main

import (
	"log"
	"os"
)

func main() {
	programName := os.Args[0]
	log.Printf("Program name: %s\n", programName)

	for index, argv := range os.Args[1:] {
		log.Printf("Argument[%d]: %s\n", index+1, argv)
	}
}
