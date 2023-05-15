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
	jData, err := json.Marshal(p) //p 구조체를 바이트배열로 리턴
	if err == nil {
		fmt.Println("err", err)
	}

	var p2 Person
	err = json.Unmarshal(jData, &p2)

	fmt.Println("p2", p2)
	fmt.Println("jData", jData)
}
