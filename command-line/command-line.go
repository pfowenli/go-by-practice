package main

import (
	"flag"
	"log"
	"os"
)

type CustomOptionData struct {
	boolPointer    *bool
	intPointer     *int
	stringPointer  *string
	tailingStrings []string
}

func parseArguments() CustomOptionData {
	data := CustomOptionData{}

	data.boolPointer = flag.Bool("bool", false, "a bool")
	data.intPointer = flag.Int("int", 10, "an int")
	data.stringPointer = flag.String("string", "oat", "a string")

	flag.Parse()
	data.tailingStrings = flag.Args()

	return data
}

func displayArguments() {
	programName := os.Args[0]
	log.Printf("Program name: %s\n", programName)

	for index, argv := range os.Args[1:] {
		log.Printf("Argument[%d]: %s\n", index+1, argv)
	}
}

func main() {
	data := parseArguments()

	if *data.boolPointer {
		displayArguments()
	}

	log.Printf("bool: %v", *data.boolPointer)
	log.Printf("int: %v", *data.intPointer)
	log.Printf("string: %v", *data.stringPointer)
	log.Printf("tailing strings: %v", data.tailingStrings)
}
