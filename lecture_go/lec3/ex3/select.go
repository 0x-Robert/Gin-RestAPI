package main

import (
	"fmt"
	"time"
)

func main() {
	ping := time.Tick(100 * time.Millisecond)
	pong := time.After(500 * time.Millisecond)

	for {
		select {
		case <-ping:
			fmt.Println("ping")
		case <-pong:
			fmt.Println("pong")
			return
		default:
			fmt.Println("-")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
