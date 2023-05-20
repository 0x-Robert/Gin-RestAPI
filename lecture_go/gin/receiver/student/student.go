package student

import (
	"fmt"
	"ginRest/gin/receiver/persons"
)

// 임베딩
type Students struct {
	persons.Persons
}

func (s Students) Add() {
	fmt.Println("Add Students")
}

func (s Students) Delete() {
	fmt.Println("Delete Student")
}
