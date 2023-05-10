package main

import (
	"fmt"
	"time"
)

func counter(s string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Println(s, " : ", i)
	}
}

func main() {

	go counter("a", 10)

	go func() {
		counter("b", 10)
	}()
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Println("c : ", i)
		}
	}()

	counter("d", 10)
}
