package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func Database() *sql.DB {
	if db != nil {
		return db
	}

	connection, err := sql.Open("mysql", "USERNAME:PASSWORD@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	db = connection

	return db
}
