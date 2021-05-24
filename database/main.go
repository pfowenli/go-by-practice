package main

import "log"

func main() {
	OpenDatabase()

	student := FindStudentById(1)
	log.Println(student)

	CloseDatabase()
}
