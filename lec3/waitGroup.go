package main

import (
	"fmt"
	"sync"
	"time"
)

func counter(n int) int {
	for i := 0; i < n; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}

}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			counter(i)
		}()
	}
	wg.Wait()
}
