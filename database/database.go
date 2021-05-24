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

func OpenDatabase() *sql.DB {
	if db != nil {
		return db
	}

	connection, err := sql.Open("mysql", "USERNAME:PASSWORD@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	db = connection

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseDatabase() {
	defer db.Close()
	db = nil
}

func FindStudentById(id int) Student {
	rows, err := db.Query("SELECT id, name, grade FROM student WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	student := Student{}

	for rows.Next() {
		err := rows.Scan(&student.id, &student.name, &student.grade)
		if err != nil {
			log.Fatal(err)
		}
	}

	return student
}
