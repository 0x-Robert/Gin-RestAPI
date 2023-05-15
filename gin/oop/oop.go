package main

import "fmt"

type Persons struct {
}

func (b Persons) Add() {
	fmt.Println("Add Person")
}

func (b Persons) Delete() {
	fmt.Println("Delete Person")
}

func main() {
	pers := &Persons{}
	pers.Add()
	pers.Delete()
}
