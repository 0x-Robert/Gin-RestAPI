package main

import (
	"ginRest/gin/receiver/persons"
	"ginRest/gin/receiver/student"
)

func main() {
	pers := &persons.Persons{}
	std := &student.Students{}

	pers.Alice()
	pers.Bob()

	std.Add()
	std.Delete()
}
