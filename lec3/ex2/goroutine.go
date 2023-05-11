package main

import (
	"fmt"
	"time"
)

func counter2(s string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Println(s, " : ", i)
	}
}

func main() {

	go counter2("a", 10)

	go func() {
		counter2("b", 10)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Println("c : ", i)
		}
	}()

	counter2("d", 10)
}
