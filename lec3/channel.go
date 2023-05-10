package main

import (
	"fmt"
	"sync"
)

func calc(wg *sync.WaitGroup, ch chan int) {
	// 내용 생략
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	var wg sync.WaitGroup
	ch2 := make(chan int)
	go calc(&wg, ch2)

}
