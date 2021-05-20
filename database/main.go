package main

import (
	"database/sql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "USERNAME:PASSWORD@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
