package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	id    int
	name  string
	grade int
}

var (
	db *sql.DB
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func OpenDatabase() *sql.DB {
	if db != nil {
		return db
	}

	connection, err := sql.Open("mysql", "USERNAME:PASSWORD@tcp(127.0.0.1:3306)/test")
	failOnError(err, "Failed to connect to mysql")

	db = connection

	err = db.Ping()
	failOnError(err, "Failed to verify the connection")

	return db
}

func CloseDatabase() {
	defer db.Close()
	db = nil
}

func FindStudentById(id int) Student {
	rows, err := db.Query("SELECT id, name, grade FROM student WHERE id = ?", id)
	failOnError(err, "Failed to execute a query")

	defer rows.Close()

	student := Student{}

	for rows.Next() {
		err := rows.Scan(&student.id, &student.name, &student.grade)
		failOnError(err, "Failed to copy values")
	}

	return student
}
