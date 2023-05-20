package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Pnum string
}

func main() {

	p := Person{Name: "codz", Age: 23, Pnum: "01098765432"}
	//jData, err := json.Marshal(p1)
	jData, err := json.Marshal(p)
	if err == nil {
		fmt.Println(err)
	}
	fmt.Println("p", p)
	fmt.Println("jData", jData)
}
