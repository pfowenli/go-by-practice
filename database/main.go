package main

import (
	"log"
)

func main() {
	db := Database()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
