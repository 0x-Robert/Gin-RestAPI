package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
	Pnum string
}

func main() {
	p := Person{Name: "codz", Age: 23, Pnum: "01098765432"}
	jData, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("jData", jData)

	var p2 Person
	err2 := json.Unmarshal(jData, &p2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(p2)
	fmt.Println(p2.Name)

	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(p); err != nil {
		panic(err)
	}

	fmt.Println(p)
	jdata, _ := json.Marshal(p)
	fmt.Println(jdata)
}
